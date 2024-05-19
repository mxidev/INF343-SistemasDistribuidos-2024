// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/message.proto

package proto

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

type Message struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f3a5e1293a7bcd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "proto.Message")
}

func init() {
	proto.RegisterFile("proto/message.proto", fileDescriptor_33f3a5e1293a7bcd)
}

var fileDescriptor_33f3a5e1293a7bcd = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x03, 0xf3, 0x84, 0x58, 0xc1, 0x94,
	0x92, 0x2c, 0x17, 0xbb, 0x2f, 0x44, 0x5c, 0x48, 0x88, 0x8b, 0x25, 0x29, 0x3f, 0xa5, 0x52, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x36, 0x2a, 0xe3, 0xe2, 0x83, 0x4a, 0x07, 0xa7, 0x16,
	0x95, 0x65, 0x26, 0xa7, 0x0a, 0xe9, 0x71, 0xf1, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0xf8,
	0xe6, 0x97, 0xe6, 0x95, 0x08, 0xf1, 0x41, 0xcc, 0xd3, 0x83, 0x2a, 0x93, 0x42, 0xe3, 0x0b, 0x99,
	0x70, 0x09, 0x41, 0xd5, 0x7b, 0xe6, 0xa5, 0xe5, 0x17, 0xe5, 0x26, 0x96, 0x64, 0xe6, 0xe7, 0x11,
	0xd2, 0xe5, 0x24, 0x1a, 0x25, 0xec, 0x93, 0x98, 0x94, 0x5f, 0x94, 0x58, 0x92, 0x5f, 0x94, 0x99,
	0xaf, 0x60, 0xa2, 0x0f, 0x96, 0x4e, 0x62, 0x03, 0x53, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5e, 0xba, 0x2a, 0x45, 0xd2, 0x00, 0x00, 0x00,
}
