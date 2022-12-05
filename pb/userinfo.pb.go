// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: userinfo.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
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

type GetUserInfoRequest struct {
	Header               *RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetUserInfoRequest) Reset()         { *m = GetUserInfoRequest{} }
func (m *GetUserInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoRequest) ProtoMessage()    {}
func (*GetUserInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_785a78c34699a93d, []int{0}
}
func (m *GetUserInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoRequest.Unmarshal(m, b)
}
func (m *GetUserInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetUserInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoRequest.Merge(m, src)
}
func (m *GetUserInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoRequest.Size(m)
}
func (m *GetUserInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoRequest proto.InternalMessageInfo

func (m *GetUserInfoRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type GetUserInfoResponse struct {
	PaidTier             bool     `protobuf:"varint,1,opt,name=paid_tier,json=paidTier,proto3" json:"paid_tier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInfoResponse) Reset()         { *m = GetUserInfoResponse{} }
func (m *GetUserInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserInfoResponse) ProtoMessage()    {}
func (*GetUserInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_785a78c34699a93d, []int{1}
}
func (m *GetUserInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInfoResponse.Unmarshal(m, b)
}
func (m *GetUserInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetUserInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInfoResponse.Merge(m, src)
}
func (m *GetUserInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserInfoResponse.Size(m)
}
func (m *GetUserInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInfoResponse proto.InternalMessageInfo

func (m *GetUserInfoResponse) GetPaidTier() bool {
	if m != nil {
		return m.PaidTier
	}
	return false
}

func init() {
	proto.RegisterType((*GetUserInfoRequest)(nil), "userinfo.GetUserInfoRequest")
	proto.RegisterType((*GetUserInfoResponse)(nil), "userinfo.GetUserInfoResponse")
}

func init() { proto.RegisterFile("userinfo.proto", fileDescriptor_785a78c34699a93d) }

var fileDescriptor_785a78c34699a93d = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0x4e, 0x2d,
	0xca, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0xf8,
	0x72, 0x53, 0x4b, 0x12, 0x11, 0x32, 0x4a, 0xae, 0x5c, 0x42, 0xee, 0xa9, 0x25, 0xa1, 0xc5, 0xa9,
	0x45, 0x9e, 0x79, 0x69, 0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xfa, 0x5c, 0x6c,
	0x19, 0xa9, 0x89, 0x29, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xe2, 0x7a, 0x70,
	0x6d, 0x50, 0x25, 0x1e, 0x60, 0xe9, 0x20, 0xa8, 0x32, 0x25, 0x23, 0x2e, 0x61, 0x14, 0x63, 0x8a,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xa4, 0xb9, 0x38, 0x0b, 0x12, 0x33, 0x53, 0xe2, 0x4b, 0x32,
	0xa1, 0x46, 0x71, 0x04, 0x71, 0x80, 0x04, 0x42, 0x32, 0x53, 0x8b, 0x8c, 0xfc, 0xb8, 0x38, 0x60,
	0x1a, 0x84, 0x9c, 0xb8, 0x98, 0xdd, 0x53, 0x4b, 0x84, 0x64, 0xf4, 0xe0, 0x0e, 0xc7, 0x74, 0x95,
	0x94, 0x2c, 0x0e, 0x59, 0x88, 0x65, 0x4e, 0x22, 0x51, 0x42, 0xc5, 0x25, 0xf9, 0x45, 0x59, 0x7a,
	0x99, 0xf9, 0xfa, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0xfa, 0x05, 0x49, 0x49, 0x6c, 0x60, 0x7f,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x44, 0x76, 0xb2, 0x02, 0x13, 0x01, 0x00, 0x00,
}
