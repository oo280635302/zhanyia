// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package msg

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

type InternalArray struct {
	InternalArray        []int64  `protobuf:"varint,1,rep,packed,name=internal_array,json=internalArray,proto3" json:"internal_array,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalArray) Reset()         { *m = InternalArray{} }
func (m *InternalArray) String() string { return proto.CompactTextString(m) }
func (*InternalArray) ProtoMessage()    {}
func (*InternalArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *InternalArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalArray.Unmarshal(m, b)
}
func (m *InternalArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalArray.Marshal(b, m, deterministic)
}
func (m *InternalArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalArray.Merge(m, src)
}
func (m *InternalArray) XXX_Size() int {
	return xxx_messageInfo_InternalArray.Size(m)
}
func (m *InternalArray) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalArray.DiscardUnknown(m)
}

var xxx_messageInfo_InternalArray proto.InternalMessageInfo

func (m *InternalArray) GetInternalArray() []int64 {
	if m != nil {
		return m.InternalArray
	}
	return nil
}

// 消消乐-图
type ClearJoyImage struct {
	// 单排-宽度
	Width int32 `protobuf:"varint,1,opt,name=Width,proto3" json:"Width,omitempty"`
	// 高度
	Height int32 `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	// 本体
	Body                 []int64  `protobuf:"varint,3,rep,packed,name=Body,proto3" json:"Body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearJoyImage) Reset()         { *m = ClearJoyImage{} }
func (m *ClearJoyImage) String() string { return proto.CompactTextString(m) }
func (*ClearJoyImage) ProtoMessage()    {}
func (*ClearJoyImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *ClearJoyImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearJoyImage.Unmarshal(m, b)
}
func (m *ClearJoyImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearJoyImage.Marshal(b, m, deterministic)
}
func (m *ClearJoyImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearJoyImage.Merge(m, src)
}
func (m *ClearJoyImage) XXX_Size() int {
	return xxx_messageInfo_ClearJoyImage.Size(m)
}
func (m *ClearJoyImage) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearJoyImage.DiscardUnknown(m)
}

var xxx_messageInfo_ClearJoyImage proto.InternalMessageInfo

func (m *ClearJoyImage) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ClearJoyImage) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ClearJoyImage) GetBody() []int64 {
	if m != nil {
		return m.Body
	}
	return nil
}

type Reply struct {
	Image                *ClearJoyImage `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetImage() *ClearJoyImage {
	if m != nil {
		return m.Image
	}
	return nil
}

func init() {
	proto.RegisterType((*InternalArray)(nil), "InternalArray")
	proto.RegisterType((*ClearJoyImage)(nil), "ClearJoyImage")
	proto.RegisterType((*Reply)(nil), "reply")
}

func init() {
	proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899)
}

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe3, 0xe2, 0xf5, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b,
	0xcc, 0x71, 0x2c, 0x2a, 0x4a, 0xac, 0x14, 0x52, 0xe5, 0xe2, 0xcb, 0x84, 0x0a, 0xc4, 0x27, 0x82,
	0x44, 0x24, 0x18, 0x15, 0x98, 0x35, 0x98, 0x83, 0x78, 0x33, 0x91, 0x95, 0x29, 0x05, 0x72, 0xf1,
	0x3a, 0xe7, 0xa4, 0x26, 0x16, 0x79, 0xe5, 0x57, 0x7a, 0xe6, 0x26, 0xa6, 0xa7, 0x0a, 0x89, 0x70,
	0xb1, 0x86, 0x67, 0xa6, 0x94, 0x64, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x38, 0x42,
	0x62, 0x5c, 0x6c, 0x1e, 0xa9, 0x99, 0xe9, 0x19, 0x25, 0x12, 0x4c, 0x60, 0x61, 0x28, 0x4f, 0x48,
	0x88, 0x8b, 0xc5, 0x29, 0x3f, 0xa5, 0x52, 0x82, 0x19, 0x6c, 0x36, 0x98, 0xad, 0xa4, 0xcb, 0xc5,
	0x5a, 0x94, 0x5a, 0x90, 0x53, 0x29, 0xa4, 0xc2, 0xc5, 0x9a, 0x09, 0x32, 0x13, 0x6c, 0x14, 0xb7,
	0x11, 0x9f, 0x1e, 0x8a, 0x4d, 0x41, 0x10, 0xc9, 0x24, 0x36, 0xb0, 0x07, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x1c, 0x32, 0x82, 0xd9, 0xcd, 0x00, 0x00, 0x00,
}