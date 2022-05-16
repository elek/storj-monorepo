// Code generated by protoc-gen-pico. DO NOT EDIT.
// source: test.proto
//
// versions:
//     protoc-gen-pico: (devel)
//     protoc:          v3.20.0

package picotest

import (
	picobuf "storj.io/picobuf"
	pic "storj.io/picobuf/internal/picotest/pic"
	picoconv "storj.io/picobuf/picoconv"
	picowire "storj.io/picobuf/picowire"
	time "time"
)

type Basic struct {
	Byte0        int32  `json:"byte0,omitempty"`
	Byte1        int32  `json:"byte1,omitempty"`
	Byte100      int32  `json:"byte100,omitempty"`
	Byte255      int32  `json:"byte255,omitempty"`
	StringEmpty  string `json:"string_empty,omitempty"`
	StringHello  string `json:"string_hello,omitempty"`
	BytesZero    []byte `json:"bytes_zero,omitempty"`
	BytesNumbers []byte `json:"bytes_numbers,omitempty"`
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
	Name    string   `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
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
	Street string `json:"street,omitempty"`
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
	Int32     int32      `json:"int32,omitempty"`
	Int64     int64      `json:"int64,omitempty"`
	Uint32    uint32     `json:"uint32,omitempty"`
	Uint64    uint64     `json:"uint64,omitempty"`
	Sint32    int32      `json:"sint32,omitempty"`
	Sint64    int64      `json:"sint64,omitempty"`
	Fixed32   uint32     `json:"fixed32,omitempty"`
	Fixed64   uint64     `json:"fixed64,omitempty"`
	Sfixed32  int32      `json:"sfixed32,omitempty"`
	Sfixed64  int64      `json:"sfixed64,omitempty"`
	Float     float32    `json:"float,omitempty"`
	Double    float64    `json:"double,omitempty"`
	Bool      bool       `json:"bool,omitempty"`
	String_   string     `json:"string,omitempty"`
	Bytes     []byte     `json:"bytes,omitempty"`
	Message   *Message   `json:"message,omitempty"`
	Int32S    []int32    `json:"int32s,omitempty"`
	Int64S    []int64    `json:"int64s,omitempty"`
	Uint32S   []uint32   `json:"uint32s,omitempty"`
	Uint64S   []uint64   `json:"uint64s,omitempty"`
	Sint32S   []int32    `json:"sint32s,omitempty"`
	Sint64S   []int64    `json:"sint64s,omitempty"`
	Fixed32S  []uint32   `json:"fixed32s,omitempty"`
	Fixed64S  []uint64   `json:"fixed64s,omitempty"`
	Sfixed32S []int32    `json:"sfixed32s,omitempty"`
	Sfixed64S []int64    `json:"sfixed64s,omitempty"`
	Floats    []float32  `json:"floats,omitempty"`
	Doubles   []float64  `json:"doubles,omitempty"`
	Bools     []bool     `json:"bools,omitempty"`
	Strings   []string   `json:"strings,omitempty"`
	Bytess    [][]byte   `json:"bytess,omitempty"`
	Messages  []*Message `json:"messages,omitempty"`
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
		x := new(Message)
		c.Loop(x.Decode)
		m.Messages = append(m.Messages, x)
	})
}

type Message struct {
	Int32 int32 `json:"int32,omitempty"`
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
	Id  pic.ID `json:"id,omitempty"`
	Alt string `json:"alt,omitempty"`
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
	Values map[int32]int32 `json:"values,omitempty"`
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
	Value        string   `json:"value,omitempty"`
	Opt          *string  `json:"opt,omitempty"`
	Rep          []string `json:"rep,omitempty"`
	PresentBasic string   `json:"present_basic,omitempty"`
	PresentOpt   string   `json:"present_opt,omitempty"`
	PresentRep   []string `json:"present_rep,omitempty"`
	OptBytes     *[]byte  `json:"opt_bytes,omitempty"`
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
	Value        *Message   `json:"value,omitempty"`
	Opt          *Message   `json:"opt,omitempty"`
	Rep          []*Message `json:"rep,omitempty"`
	PresentBasic Message    `json:"present_basic,omitempty"`
	PresentOpt   Message    `json:"present_opt,omitempty"`
	PresentRep   []Message  `json:"present_rep,omitempty"`
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
		x := new(Message)
		c.Loop(x.Decode)
		m.Rep = append(m.Rep, x)
	})
	c.PresentMessage(4, m.PresentBasic.Decode)
	c.PresentMessage(5, m.PresentOpt.Decode)
	c.RepeatedMessage(6, func(c *picobuf.Decoder) {
		m.PresentRep = append(m.PresentRep, Message{})
		c.Loop(m.PresentRep[len(m.PresentRep)-1].Decode)
	})
}

type VariationsCustom struct {
	Value        pic.ID  `json:"value,omitempty"`
	Opt          *pic.ID `json:"opt,omitempty"`
	PresentBasic pic.ID  `json:"present_basic,omitempty"`
	PresentOpt   pic.ID  `json:"present_opt,omitempty"`
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
	if c.PendingField() == 2 {
		if m.Opt == nil {
			m.Opt = new(pic.ID)
		}
		m.Opt.PicoDecode(c, 2)
	}
	m.PresentBasic.PicoDecode(c, 4)
	m.PresentOpt.PicoDecode(c, 5)
}

type Timestamp struct {
	Seconds int64 `json:"seconds,omitempty"`
	Nanos   int32 `json:"nanos,omitempty"`
}

func (m *Timestamp) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Int64(1, &m.Seconds)
	c.Int32(2, &m.Nanos)
	return true
}

func (m *Timestamp) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Int64(1, &m.Seconds)
	c.Int32(2, &m.Nanos)
}

type CustomMessageTypes struct {
	Normal                        *Timestamp       `json:"normal,omitempty"`
	CustomType                    *pic.Timestamp   `json:"custom_type,omitempty"`
	PresentCustomType             pic.Timestamp    `json:"present_custom_type,omitempty"`
	CustomTypeCast                *time.Time       `json:"custom_type_cast,omitempty"`
	PresentCustomTypeCast         time.Time        `json:"present_custom_type_cast,omitempty"`
	RepeatedCustomType            []*pic.Timestamp `json:"repeated_custom_type,omitempty"`
	RepeatedPresentCustomType     []pic.Timestamp  `json:"repeated_present_custom_type,omitempty"`
	RepeatedCustomTypeCast        []*time.Time     `json:"repeated_custom_type_cast,omitempty"`
	RepeatedPresentCustomTypeCast []time.Time      `json:"repeated_present_custom_type_cast,omitempty"`
}

func (m *CustomMessageTypes) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Message(1, m.Normal.Encode)
	m.CustomType.PicoEncode(c, 2)
	m.PresentCustomType.PicoEncode(c, 3)
	(*picoconv.Timestamp)(m.CustomTypeCast).PicoEncode(c, 4)
	(*picoconv.Timestamp)(&m.PresentCustomTypeCast).PicoEncode(c, 5)
	for _, x := range m.RepeatedCustomType {
		x.PicoEncode(c, 6)
	}
	for i := range m.RepeatedPresentCustomType {
		x := &m.RepeatedPresentCustomType[i]
		x.PicoEncode(c, 7)
	}
	for _, x := range m.RepeatedCustomTypeCast {
		(*picoconv.Timestamp)(x).PicoEncode(c, 8)
	}
	for i := range m.RepeatedPresentCustomTypeCast {
		x := &m.RepeatedPresentCustomTypeCast[i]
		(*picoconv.Timestamp)(x).PicoEncode(c, 9)
	}
	return true
}

func (m *CustomMessageTypes) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Message(1, func(c *picobuf.Decoder) {
		if m.Normal == nil {
			m.Normal = new(Timestamp)
		}
		m.Normal.Decode(c)
	})
	if c.PendingField() == 2 {
		if m.CustomType == nil {
			m.CustomType = new(pic.Timestamp)
		}
		m.CustomType.PicoDecode(c, 2)
	}
	m.PresentCustomType.PicoDecode(c, 3)
	if c.PendingField() == 4 {
		if m.CustomTypeCast == nil {
			m.CustomTypeCast = new(time.Time)
		}
		(*picoconv.Timestamp)(m.CustomTypeCast).PicoDecode(c, 4)
	}
	(*picoconv.Timestamp)(&m.PresentCustomTypeCast).PicoDecode(c, 5)
	for c.PendingField() == 6 {
		x := new(pic.Timestamp)
		x.PicoDecode(c, 6)
		m.RepeatedCustomType = append(m.RepeatedCustomType, x)
	}
	for c.PendingField() == 7 {
		m.RepeatedPresentCustomType = append(m.RepeatedPresentCustomType, pic.Timestamp{})
		m.RepeatedPresentCustomType[len(m.RepeatedPresentCustomType)-1].PicoDecode(c, 7)
	}
	for c.PendingField() == 8 {
		x := new(time.Time)
		(*picoconv.Timestamp)(x).PicoDecode(c, 8)
		m.RepeatedCustomTypeCast = append(m.RepeatedCustomTypeCast, x)
	}
	for c.PendingField() == 9 {
		m.RepeatedPresentCustomTypeCast = append(m.RepeatedPresentCustomTypeCast, time.Time{})
		x := &m.RepeatedPresentCustomTypeCast[len(m.RepeatedPresentCustomTypeCast)-1]
		(*picoconv.Timestamp)(x).PicoDecode(c, 9)
	}
}
