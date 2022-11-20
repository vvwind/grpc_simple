// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/databus-serv.proto

package grpc_simple

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SendRequest struct {
	//!< Keeps a value1.
	Prm1 float32 `protobuf:"fixed32,1,opt,name=prm1,proto3" json:"prm1,omitempty"`
	//!< Keeps a value2.
	Prm2                 float32  `protobuf:"fixed32,2,opt,name=prm2,proto3" json:"prm2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f332c66c57a401, []int{0}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetPrm1() float32 {
	if m != nil {
		return m.Prm1
	}
	return 0
}

func (m *SendRequest) GetPrm2() float32 {
	if m != nil {
		return m.Prm2
	}
	return 0
}

type SendResponse struct {
	//!< Keeps a result of request.
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11f332c66c57a401, []int{1}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "main.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "main.SendResponse")
}

func init() {
	proto.RegisterFile("protos/databus-serv.proto", fileDescriptor_11f332c66c57a401)
}

var fileDescriptor_11f332c66c57a401 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x49, 0x2c, 0x49, 0x4c, 0x2a, 0x2d, 0xd6, 0x2d, 0x4e, 0x2d, 0x2a, 0xd3,
	0x03, 0x8b, 0x09, 0xb1, 0xe4, 0x26, 0x66, 0xe6, 0x29, 0x99, 0x72, 0x71, 0x07, 0xa7, 0xe6, 0xa5,
	0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0x14, 0x14, 0xe5, 0x1a, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x30, 0x05, 0x81, 0xd9, 0x50, 0x31, 0x23, 0x09, 0x26, 0xb8, 0x98, 0x91,
	0x92, 0x1a, 0x17, 0x0f, 0x44, 0x5b, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17, 0x5b,
	0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0x54, 0x27, 0x94, 0x67, 0x64, 0xcf, 0xc5, 0xe7, 0x02, 0xb1,
	0x3a, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x97, 0x8b, 0xa5, 0x38, 0x35, 0x2f, 0x45,
	0x48, 0x50, 0x0f, 0x64, 0xbf, 0x1e, 0x92, 0xe5, 0x52, 0x42, 0xc8, 0x42, 0x10, 0x83, 0x9d, 0xb4,
	0xa3, 0x34, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xcb, 0xca, 0xca,
	0x33, 0xf3, 0x52, 0xf4, 0xd3, 0x8b, 0x0a, 0x92, 0xe3, 0x8b, 0x33, 0x73, 0x0b, 0x72, 0x52, 0xad,
	0x91, 0xd8, 0x49, 0x6c, 0x60, 0x9f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x00, 0x9f,
	0xcd, 0xf6, 0x00, 0x00, 0x00,
}
