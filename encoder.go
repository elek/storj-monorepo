// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf

import (
	"encoding/binary"

	"google.golang.org/protobuf/encoding/protowire"
)

// Encoder implements encoding of protobuf format.
type Encoder struct {
	codec   *Codec
	buffer  []byte
	backing [64]byte
}

// NewEncoder creates a new Encoder.
func NewEncoder() *Encoder {
	codec := &Codec{}
	codec.encode.codec = codec
	codec.encode.buffer = codec.encode.backing[:0]
	return &codec.encode
}

// Codec returns the associated codec.
func (enc *Encoder) Codec() *Codec { return enc.codec }

// Buffer returns the encoded internal buffer.
func (enc *Encoder) Buffer() []byte { return enc.buffer }

// Message decodes a message.
func (enc *Encoder) Message(field FieldNumber, fn func(*Codec) bool) {
	enc.encodeAnyBytes(field, func() bool {
		return fn(enc.codec)
	})
}

// PresentMessage encodes an always present message.
func (enc *Encoder) PresentMessage(field FieldNumber, fn func(*Codec) bool) {
	enc.encodeAnyBytes(field, func() bool {
		lengthStart := len(enc.buffer)
		fn(enc.codec)
		return len(enc.buffer) > lengthStart
	})
}

// RepeatedMessage encodes a repeated message.
func (enc *Encoder) RepeatedMessage(field FieldNumber, fn func(c *Codec, index int) bool) {
	for index := 0; ; index++ {
		ok := enc.encodeAnyBytes(field, func() bool {
			return fn(enc.codec, index)
		})
		if !ok {
			return
		}
	}
}

// encodeAnyBytes encodes field as Bytes and handles encoding the length.
func (enc *Encoder) encodeAnyBytes(field FieldNumber, fn func() bool) bool {
	tagStart := len(enc.buffer)
	enc.buffer = appendTag(enc.buffer, field, protowire.BytesType)
	lengthStart := len(enc.buffer)
	// We'll guess that we need 2 bytes for length.
	// If we need less, then the copy is fast, and needing more is unlikely.
	var lengthBufferPrediction [2]byte
	enc.buffer = append(enc.buffer, lengthBufferPrediction[:]...)
	messageStart := len(enc.buffer)
	// encode the submessage
	ok := fn()
	if !ok {
		// The message was nil, we can remove the tag.
		enc.buffer = enc.buffer[:tagStart]
		return false
	}
	messageLength := len(enc.buffer) - messageStart
	bytesForSize := protowire.SizeVarint(uint64(messageLength))
	if bytesForSize == len(lengthBufferPrediction) {
		binary.PutUvarint(enc.buffer[lengthStart:messageStart], uint64(messageLength))
		return true
	}
	if bytesForSize > len(lengthBufferPrediction) {
		enc.buffer = append(enc.buffer, make([]byte, bytesForSize-len(lengthBufferPrediction))...)
	}

	copy(enc.buffer[lengthStart+bytesForSize:], enc.buffer[messageStart:])
	binary.PutUvarint(enc.buffer[lengthStart:lengthStart+bytesForSize], uint64(messageLength))
	enc.buffer = enc.buffer[:lengthStart+bytesForSize+messageLength]
	return true
}
