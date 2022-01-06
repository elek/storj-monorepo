// Code generated by protoc-gen-pico. DO NOT EDIT.
// source: test.proto
//
// versions:
//     protoc-gen-pico: (devel)
//     protoc:          v3.19.1

package picotest

import (
	picobuf "storj.io/picobuf"
	pic "storj.io/picobuf/internal/picotest/pic"
	picowire "storj.io/picobuf/picowire"
)

type Basic struct {
	Byte0        int32
	Byte1        int32
	Byte100      int32
	Byte255      int32
	StringEmpty  string
	StringHello  string
	BytesZero    []byte
	BytesNumbers []byte
}

func (m *Basic) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Int32(1, &m.Byte0)
	c.Int32(2, &m.Byte1)
	c.Int32(3, &m.Byte100)
	c.Int32(4, &m.Byte255)
	c.String(5, &m.StringEmpty)
	c.String(6, &m.StringHello)
	c.Bytes(7, &m.BytesZero)
	c.Bytes(8, &m.BytesNumbers)
	return true
}

func (m *Basic) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Int32(1, &m.Byte0)
	c.Int32(2, &m.Byte1)
	c.Int32(3, &m.Byte100)
	c.Int32(4, &m.Byte255)
	c.String(5, &m.StringEmpty)
	c.String(6, &m.StringHello)
	c.Bytes(7, &m.BytesZero)
	c.Bytes(8, &m.BytesNumbers)
}

type Person struct {
	Name    string
	Address *Address
}

func (m *Person) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.String(1, &m.Name)
	c.Message(2, m.Address.Encode)
	return true
}

func (m *Person) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.String(1, &m.Name)
	c.Message(2, func(c *picobuf.Decoder) {
		if m.Address == nil {
			m.Address = new(Address)
		}
		m.Address.Decode(c)
	})
}

type Address struct {
	Street string
}

func (m *Address) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.String(1, &m.Street)
	return true
}

func (m *Address) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.String(1, &m.Street)
}

type AllTypes struct {
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
}

func (m *AllTypes) Encode(c *picobuf.Encoder) bool {
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
	for _, x := range m.Messages {
		c.AlwaysMessage(32, x.Encode)
	}
	return true
}

func (m *AllTypes) Decode(c *picobuf.Decoder) {
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

type Piece struct {
	Id  pic.ID
	Alt string
}

func (m *Piece) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	m.Id.PicoEncode(c, 1)
	(*pic.RawString)(&m.Alt).PicoEncode(c, 2)
	return true
}

func (m *Piece) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	m.Id.PicoDecode(c, 1)
	(*pic.RawString)(&m.Alt).PicoDecode(c, 2)
}

type Map struct {
	Values map[int32]int32
}

func (m *Map) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	(*picowire.MapInt32Int32)(&m.Values).PicoEncode(c, 1)
	return true
}

func (m *Map) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	(*picowire.MapInt32Int32)(&m.Values).PicoDecode(c, 1)
}

type VariationsScalar struct {
	Value        string
	Opt          *string
	Rep          []string
	PresentBasic string
	PresentOpt   string
	PresentRep   []string
	OptBytes     *[]byte
}

func (m *VariationsScalar) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.String(1, &m.Value)
	if m.Opt != nil {
		c.String(2, m.Opt)
	}
	c.RepeatedString(3, &m.Rep)
	c.String(4, &m.PresentBasic)
	c.String(5, &m.PresentOpt)
	c.RepeatedString(6, &m.PresentRep)
	if m.OptBytes != nil {
		c.Bytes(7, m.OptBytes)
	}
	return true
}

func (m *VariationsScalar) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.String(1, &m.Value)
	if c.PendingField() == 2 {
		m.Opt = new(string)
		c.String(2, m.Opt)
	}
	c.RepeatedString(3, &m.Rep)
	c.String(4, &m.PresentBasic)
	c.String(5, &m.PresentOpt)
	c.RepeatedString(6, &m.PresentRep)
	if c.PendingField() == 7 {
		m.OptBytes = new([]byte)
		c.Bytes(7, m.OptBytes)
	}
}

type VariationsMessage struct {
	Value        *Message
	Opt          *Message
	Rep          []*Message
	PresentBasic Message
	PresentOpt   Message
	PresentRep   []Message
}

func (m *VariationsMessage) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Message(1, m.Value.Encode)
	c.Message(2, m.Opt.Encode)
	for _, x := range m.Rep {
		c.AlwaysMessage(3, x.Encode)
	}
	c.PresentMessage(4, m.PresentBasic.Encode)
	c.PresentMessage(5, m.PresentOpt.Encode)
	for i := range m.PresentRep {
		x := &m.PresentRep[i]
		c.AlwaysMessage(6, x.Encode)
	}
	return true
}

func (m *VariationsMessage) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Message(1, func(c *picobuf.Decoder) {
		if m.Value == nil {
			m.Value = new(Message)
		}
		m.Value.Decode(c)
	})
	c.Message(2, func(c *picobuf.Decoder) {
		if m.Opt == nil {
			m.Opt = new(Message)
		}
		m.Opt.Decode(c)
	})
	c.RepeatedMessage(3, func(c *picobuf.Decoder) {
		mm := new(Message)
		c.Loop(mm.Decode)
		m.Rep = append(m.Rep, mm)
	})
	c.PresentMessage(4, m.PresentBasic.Decode)
	c.PresentMessage(5, m.PresentOpt.Decode)
	c.RepeatedMessage(6, func(c *picobuf.Decoder) {
		m.PresentRep = append(m.PresentRep, Message{})
		c.Loop(m.PresentRep[len(m.PresentRep)-1].Decode)
	})
}

type VariationsCustom struct {
	Value        pic.ID
	Opt          *pic.ID
	PresentBasic pic.ID
	PresentOpt   pic.ID
}

func (m *VariationsCustom) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	m.Value.PicoEncode(c, 1)
	m.Opt.PicoEncode(c, 2)
	m.PresentBasic.PicoEncode(c, 4)
	m.PresentOpt.PicoEncode(c, 5)
	return true
}

func (m *VariationsCustom) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	m.Value.PicoDecode(c, 1)
	m.Opt.PicoDecode(c, 2)
	m.PresentBasic.PicoDecode(c, 4)
	m.PresentOpt.PicoDecode(c, 5)
}
