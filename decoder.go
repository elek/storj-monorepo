// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protowire"

	"storj.io/picobuf/internal/bitset"
)

const (
	fieldDecodingErrored = FieldNumber(-1)
	fieldDecodingDone    = FieldNumber(-2)
)

// Decoder implements decoding of protobuf messages.
type Decoder struct {
	messageDecodeState
	stack []messageDecodeState

	codec *Codec
	err   error
}

type messageDecodeState struct {
	pendingField FieldNumber
	pendingWire  protowire.Type

	buffer []byte
	filled bitset.Small
	// TODO: we might need to differentiate between filled and zeroed fields.
}

// NewDecoder returns a new Decoder.
func NewDecoder(data []byte) *Decoder {
	codec := &Codec{
		encoding: false,
	}
	codec.decode.codec = codec
	codec.decode.buffer = data
	codec.decode.nextField(0)
	return &codec.decode
}

// Codec returns the associated codec.
func (dec *Decoder) Codec() *Codec { return dec.codec }

// Err returns error that occurred during decoding.
func (dec *Decoder) Err() error {
	return dec.err
}

func (dec *Decoder) pushState(message []byte) {
	dec.stack = append(dec.stack, dec.messageDecodeState)
	dec.messageDecodeState = messageDecodeState{
		buffer: message,
	}
	dec.nextField(0)
}

func (dec *Decoder) popState() {
	if len(dec.stack) == 0 {
		dec.fail(0, "stack mangled")
		return
	}
	dec.messageDecodeState = dec.stack[len(dec.stack)-1]
	dec.stack = dec.stack[:len(dec.stack)-1]
}

// MessageReferenceUpdate should set the message to nil
// when allocate == false and allocate otherwise.
type MessageReferenceUpdate func(allocate bool)

// Message decodes a message.
func (dec *Decoder) Message(field FieldNumber, isNil func() bool, update MessageReferenceUpdate, fn func(*Codec)) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			update(false)
		}
		return
	}
	if dec.pendingWire != protowire.BytesType {
		dec.fail(field, "expected wire type Bytes")
		return
	}

	message, n := protowire.ConsumeBytes(dec.buffer)
	dec.pushState(message)
	update(true)
	fn(dec.codec)
	dec.popState()

	dec.nextField(n)
	dec.filled.Set(int32(field))
}

// PresentMessage decodes an always present message.
func (dec *Decoder) PresentMessage(field FieldNumber, fn func(*Codec)) {
	if field != dec.pendingField {
		if !dec.filled.Set(int32(field)) {
			dec.pushState(nil)
			fn(dec.codec)
			dec.popState()
		}
		return
	}
	if dec.pendingWire != protowire.BytesType {
		dec.fail(field, "expected wire type Bytes")
		return
	}

	message, n := protowire.ConsumeBytes(dec.buffer)
	dec.pushState(message)
	fn(dec.codec)
	dec.popState()

	dec.nextField(n)
	dec.filled.Set(int32(field))
}

//go:noinline
func (dec *Decoder) fail(field FieldNumber, msg string) {
	// TODO: use static error types
	dec.pendingField = fieldDecodingErrored
	dec.err = fmt.Errorf("failed while parsing %v: %s", field, msg)
}

func (dec *Decoder) nextField(advance int) {
	dec.buffer = dec.buffer[advance:]
	if len(dec.buffer) == 0 {
		dec.pendingField = fieldDecodingDone
		return
	}

	field, wire, n := protowire.ConsumeTag(dec.buffer)
	if n < 0 {
		dec.fail(0, "failed to parse") // TODO: better error message
		return
	}
	dec.buffer = dec.buffer[n:]
	dec.pendingField, dec.pendingWire = FieldNumber(field), wire
}

func init() {
	// silence structcheck linting bug about unused field
	var z messageDecodeState
	z.pendingField = 0
	z.pendingWire = 0
	z.filled.Set(0)
}
