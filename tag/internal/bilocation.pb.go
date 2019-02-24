// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bilocation.proto

package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Set struct {
	Tags                 []*Tag   `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Set) Reset()         { *m = Set{} }
func (m *Set) String() string { return proto.CompactTextString(m) }
func (*Set) ProtoMessage()    {}
func (*Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_bilocation_867ebc3ed535a4e5, []int{0}
}
func (m *Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Set.Unmarshal(m, b)
}
func (m *Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Set.Marshal(b, m, deterministic)
}
func (dst *Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Set.Merge(dst, src)
}
func (m *Set) XXX_Size() int {
	return xxx_messageInfo_Set.Size(m)
}
func (m *Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Set proto.InternalMessageInfo

func (m *Set) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Tag struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Classifiers          []string `protobuf:"bytes,2,rep,name=classifiers,proto3" json:"classifiers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_bilocation_867ebc3ed535a4e5, []int{1}
}
func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (dst *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(dst, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetClassifiers() []string {
	if m != nil {
		return m.Classifiers
	}
	return nil
}

func init() {
	proto.RegisterType((*Set)(nil), "internal.Set")
	proto.RegisterType((*Tag)(nil), "internal.Tag")
}

func init() { proto.RegisterFile("bilocation.proto", fileDescriptor_bilocation_867ebc3ed535a4e5) }

var fileDescriptor_bilocation_867ebc3ed535a4e5 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xca, 0xcc, 0xc9,
	0x4f, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcc,
	0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0x51, 0xd2, 0xe0, 0x62, 0x0e, 0x4e, 0x2d, 0x11, 0x52, 0xe4,
	0x62, 0x29, 0x49, 0x4c, 0x2f, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd5, 0x83, 0xc9,
	0xeb, 0x85, 0x24, 0xa6, 0x07, 0x81, 0xa5, 0x94, 0xac, 0xb9, 0x98, 0x43, 0x12, 0xd3, 0x85, 0x84,
	0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21,
	0x05, 0x2e, 0xee, 0xe4, 0x9c, 0xc4, 0xe2, 0xe2, 0xcc, 0xb4, 0xcc, 0xd4, 0xa2, 0x62, 0x09, 0x26,
	0x05, 0x66, 0x0d, 0xce, 0x20, 0x64, 0xa1, 0x24, 0x36, 0xb0, 0xbd, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe4, 0x89, 0x0f, 0x0e, 0x8b, 0x00, 0x00, 0x00,
}