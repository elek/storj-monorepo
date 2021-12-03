// Code generated by protoc-gen-pico. DO NOT EDIT.
// source: msg-bas.proto
//
// versions:
//     protoc-gen-pico: (devel)
//     protoc:          v3.19.1

package bas

import (
	picobuf "storj.io/picobuf"
)

type Types struct {
	Int32     int32
	Int64     int64
	Uint32    uint32
	Uint64    uint64
	Sint32    int32
	Sint64    int64
	Fixed32   uint32
	Fixed64   uint64
	Sfixed32  int32
	Sfixed64  int64
	Float     float32
	Double    float64
	Bool      bool
	String_   string
	Bytes     []byte
	Message   *Message
	Int32S    []int32
	Int64S    []int64
	Uint32S   []uint32
	Uint64S   []uint64
	Sint32S   []int32
	Sint64S   []int64
	Fixed32S  []uint32
	Fixed64S  []uint64
	Sfixed32S []int32
	Sfixed64S []int64
	Floats    []float32
	Doubles   []float64
	Bools     []bool
	Strings   []string
	Bytess    [][]byte
	Messages  []*Message
	Values    map[string]string
}

func (m *Types) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Int32(1, &m.Int32)
	c.Int64(2, &m.Int64)
	c.Uint32(3, &m.Uint32)
	c.Uint64(4, &m.Uint64)
	c.Sint32(5, &m.Sint32)
	c.Sint64(6, &m.Sint64)
	c.Fixed32(7, &m.Fixed32)
	c.Fixed64(8, &m.Fixed64)
	c.Sfixed32(9, &m.Sfixed32)
	c.Sfixed64(10, &m.Sfixed64)
	c.Float(11, &m.Float)
	c.Double(12, &m.Double)
	c.Bool(13, &m.Bool)
	c.String(14, &m.String_)
	c.Bytes(15, &m.Bytes)
	c.Message(16, m.Message.Encode)
	c.RepeatedInt32(17, &m.Int32S)
	c.RepeatedInt64(18, &m.Int64S)
	c.RepeatedUint32(19, &m.Uint32S)
	c.RepeatedUint64(20, &m.Uint64S)
	c.RepeatedSint32(21, &m.Sint32S)
	c.RepeatedSint64(22, &m.Sint64S)
	c.RepeatedFixed32(23, &m.Fixed32S)
	c.RepeatedFixed64(24, &m.Fixed64S)
	c.RepeatedSfixed32(25, &m.Sfixed32S)
	c.RepeatedSfixed64(26, &m.Sfixed64S)
	c.RepeatedFloat(27, &m.Floats)
	c.RepeatedDouble(28, &m.Doubles)
	c.RepeatedBool(29, &m.Bools)
	c.RepeatedString(30, &m.Strings)
	c.RepeatedBytes(31, &m.Bytess)
	c.RepeatedMessage(32, func(c *picobuf.Encoder, index int) bool {
		if index >= len(m.Messages) {
			return false
		}
		m.Messages[index].Encode(c)
		return true
	})
	c.MapStringString(33, &m.Values)
	return true
}

func (m *Types) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Int32(1, &m.Int32)
	c.Int64(2, &m.Int64)
	c.Uint32(3, &m.Uint32)
	c.Uint64(4, &m.Uint64)
	c.Sint32(5, &m.Sint32)
	c.Sint64(6, &m.Sint64)
	c.Fixed32(7, &m.Fixed32)
	c.Fixed64(8, &m.Fixed64)
	c.Sfixed32(9, &m.Sfixed32)
	c.Sfixed64(10, &m.Sfixed64)
	c.Float(11, &m.Float)
	c.Double(12, &m.Double)
	c.Bool(13, &m.Bool)
	c.String(14, &m.String_)
	c.Bytes(15, &m.Bytes)
	c.Message(16, func(c *picobuf.Decoder) {
		if m.Message == nil {
			m.Message = new(Message)
		}
		m.Message.Decode(c)
	})
	c.RepeatedInt32(17, &m.Int32S)
	c.RepeatedInt64(18, &m.Int64S)
	c.RepeatedUint32(19, &m.Uint32S)
	c.RepeatedUint64(20, &m.Uint64S)
	c.RepeatedSint32(21, &m.Sint32S)
	c.RepeatedSint64(22, &m.Sint64S)
	c.RepeatedFixed32(23, &m.Fixed32S)
	c.RepeatedFixed64(24, &m.Fixed64S)
	c.RepeatedSfixed32(25, &m.Sfixed32S)
	c.RepeatedSfixed64(26, &m.Sfixed64S)
	c.RepeatedFloat(27, &m.Floats)
	c.RepeatedDouble(28, &m.Doubles)
	c.RepeatedBool(29, &m.Bools)
	c.RepeatedString(30, &m.Strings)
	c.RepeatedBytes(31, &m.Bytess)
	c.RepeatedMessage(32, func(c *picobuf.Decoder) {
		mm := new(Message)
		c.Loop(mm.Decode)
		m.Messages = append(m.Messages, mm)
	})
	c.MapStringString(33, &m.Values)
}

type Message struct {
	Int32 int32
}

func (m *Message) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Int32(1, &m.Int32)
	return true
}

func (m *Message) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Int32(1, &m.Int32)
}
