// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package picobuf_test

import (
	"testing"

	"github.com/zeebo/assert"

	"storj.io/picobuf"
	"storj.io/picobuf/internal/picotest"
)

func puint32(v uint32) *uint32 { return &v }
func pstring(v string) *string { return &v }
func pbytes(v []byte) *[]byte  { return &v }

func TestEncoder_Types(t *testing.T) {
	enc := picobuf.NewEncoder()

	enc.Uint32(1, puint32(0))
	enc.Uint32(2, puint32(1))
	enc.Uint32(3, puint32(100))
	enc.Uint32(4, puint32(255))

	enc.String(5, pstring(""))
	enc.String(6, pstring("hello"))

	enc.Bytes(7, pbytes([]byte{0}))
	enc.Bytes(8, pbytes([]byte{1, 2, 3, 4}))

	result := enc.Buffer()
	assert.Equal(t, result, []byte{0x10, 0x1, 0x18, 0x64, 0x20, 0xff, 0x1, 0x32, 0x5, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x1, 0x0, 0x42, 0x4, 0x1, 0x2, 0x3, 0x4})
}

func TestEncoder_Repeated(t *testing.T) {
	enc := picobuf.NewEncoder()

	xs := []int32{3, 270, 86942}
	enc.RepeatedInt32(4, &xs)
	result := enc.Buffer()
	assert.Equal(t, result, []byte{0x22, 0x06, 0x03, 0x8E, 0x02, 0x9E, 0xA7, 0x05})
}

func TestEncoder_SubMessage(t *testing.T) {
	person := &picotest.Person{
		Name: "Hello",
		Address: &picotest.Address{
			Street: "Home",
		},
	}

	data, err := picobuf.Marshal(person)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x6, 0xa, 0x4, 0x48, 0x6f, 0x6d, 0x65})

	person = &picotest.Person{
		Name:    "Hello",
		Address: nil,
	}
	data, err = picobuf.Marshal(person)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f})

	person = &picotest.Person{
		Name:    "Hello",
		Address: &picotest.Address{},
	}
	data, err = picobuf.Marshal(person)
	assert.NoError(t, err)
	assert.Equal(t, data, []byte{0xa, 0x5, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0})
}
