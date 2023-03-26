// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service.proto

package service

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type In struct {
	In                   int64    `protobuf:"varint,1,opt,name=in,proto3" json:"in,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *In) Reset()         { *m = In{} }
func (m *In) String() string { return proto.CompactTextString(m) }
func (*In) ProtoMessage()    {}
func (*In) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}
func (m *In) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_In.Unmarshal(m, b)
}
func (m *In) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_In.Marshal(b, m, deterministic)
}
func (m *In) XXX_Merge(src proto.Message) {
	xxx_messageInfo_In.Merge(m, src)
}
func (m *In) XXX_Size() int {
	return xxx_messageInfo_In.Size(m)
}
func (m *In) XXX_DiscardUnknown() {
	xxx_messageInfo_In.DiscardUnknown(m)
}

var xxx_messageInfo_In proto.InternalMessageInfo

func (m *In) GetIn() int64 {
	if m != nil {
		return m.In
	}
	return 0
}

func (m *In) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Out struct {
	Out                  int64    `protobuf:"varint,1,opt,name=out,proto3" json:"out,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Out) Reset()         { *m = Out{} }
func (m *Out) String() string { return proto.CompactTextString(m) }
func (*Out) ProtoMessage()    {}
func (*Out) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}
func (m *Out) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Out.Unmarshal(m, b)
}
func (m *Out) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Out.Marshal(b, m, deterministic)
}
func (m *Out) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Out.Merge(m, src)
}
func (m *Out) XXX_Size() int {
	return xxx_messageInfo_Out.Size(m)
}
func (m *Out) XXX_DiscardUnknown() {
	xxx_messageInfo_Out.DiscardUnknown(m)
}

var xxx_messageInfo_Out proto.InternalMessageInfo

func (m *Out) GetOut() int64 {
	if m != nil {
		return m.Out
	}
	return 0
}

func (m *Out) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*In)(nil), "service.In")
	proto.RegisterType((*Out)(nil), "service.Out")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x34, 0xb8,
	0x98, 0x3c, 0xf3, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83,
	0x98, 0x32, 0xf3, 0x84, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x98, 0x14, 0x18, 0x35,
	0x78, 0x82, 0xc0, 0x6c, 0x25, 0x6d, 0x2e, 0x66, 0xff, 0xd2, 0x12, 0x21, 0x01, 0x2e, 0xe6, 0xfc,
	0xd2, 0x12, 0xa8, 0x5a, 0x10, 0x13, 0x9b, 0x62, 0xa3, 0x95, 0x8c, 0x5c, 0xec, 0xc1, 0x10, 0x2b,
	0x84, 0x54, 0xb8, 0xd8, 0x7d, 0x53, 0x4b, 0x32, 0xf2, 0x53, 0x0c, 0x85, 0xb8, 0xf5, 0x60, 0xce,
	0xf0, 0xcc, 0x93, 0xe2, 0x81, 0x73, 0x40, 0xe6, 0xaa, 0xc1, 0x54, 0x19, 0xe1, 0x51, 0xa5, 0xc1,
	0x88, 0x50, 0x67, 0x8c, 0x47, 0x9d, 0x01, 0xa3, 0x90, 0x06, 0x4c, 0x9d, 0x09, 0x5e, 0xf3, 0x0c,
	0x18, 0x9d, 0x74, 0xa2, 0xb4, 0x8a, 0x4b, 0xf2, 0x8b, 0xb2, 0xf4, 0x32, 0xf3, 0xf5, 0x53, 0x8a,
	0x0a, 0x92, 0xf5, 0x33, 0xf3, 0x4a, 0x52, 0x8b, 0xf2, 0x12, 0x73, 0xc0, 0x8c, 0xf4, 0xa2, 0xc4,
	0x92, 0xcc, 0xfc, 0x3c, 0x7d, 0xa8, 0xae, 0x24, 0x36, 0x70, 0x00, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x7a, 0xab, 0xb7, 0x51, 0x51, 0x01, 0x00, 0x00,
}