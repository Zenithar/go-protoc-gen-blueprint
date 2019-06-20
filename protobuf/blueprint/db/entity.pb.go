// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity.proto

/*
Package db is a generated protocol buffer package.

It is generated from these files:
	entity.proto

It has these top-level messages:
	Entity
	Attribute
*/
package db

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Entity struct {
	Enabled     bool   `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *Entity) Reset()                    { *m = Entity{} }
func (m *Entity) String() string            { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()               {}
func (*Entity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Entity) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Attribute struct {
	Primary     bool   `protobuf:"varint,1,opt,name=primary" json:"primary,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Attribute) Reset()                    { *m = Attribute{} }
func (m *Attribute) String() string            { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()               {}
func (*Attribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Attribute) GetPrimary() bool {
	if m != nil {
		return m.Primary
	}
	return false
}

func (m *Attribute) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

var E_Entity = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*Entity)(nil),
	Field:         98765400,
	Name:          "db.entity",
	Tag:           "bytes,98765400,opt,name=entity",
	Filename:      "entity.proto",
}

var E_Attribute = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*Attribute)(nil),
	Field:         98765401,
	Name:          "db.attribute",
	Tag:           "bytes,98765401,opt,name=attribute",
	Filename:      "entity.proto",
}

func init() {
	proto.RegisterType((*Entity)(nil), "db.Entity")
	proto.RegisterType((*Attribute)(nil), "db.Attribute")
	proto.RegisterExtension(E_Entity)
	proto.RegisterExtension(E_Attribute)
}

func init() { proto.RegisterFile("entity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xcf, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x25, 0xa0, 0x40, 0xae, 0xb0, 0x78, 0xb2, 0x90, 0x10, 0x56, 0xa7, 0x4c, 0x8e,
	0x04, 0x5b, 0x37, 0x06, 0xca, 0x54, 0x21, 0x79, 0x62, 0xb5, 0xf1, 0x11, 0x59, 0x4a, 0xed, 0xc8,
	0x76, 0x87, 0xbe, 0x03, 0x23, 0x0f, 0x07, 0x6f, 0x83, 0x6a, 0xd7, 0x05, 0x91, 0x2d, 0x77, 0xb9,
	0xff, 0x93, 0x7f, 0xb8, 0x42, 0x1b, 0x4d, 0xdc, 0xf3, 0xc9, 0xbb, 0xe8, 0x48, 0xad, 0xd5, 0x0d,
	0x1b, 0x9c, 0x1b, 0x46, 0xec, 0xd3, 0x46, 0xed, 0xde, 0x7b, 0x8d, 0xe1, 0xcd, 0x9b, 0x29, 0x3a,
	0x9f, 0xaf, 0x96, 0xaf, 0xd0, 0x3c, 0xa5, 0x14, 0xa1, 0x70, 0x81, 0x56, 0xaa, 0x11, 0x35, 0xad,
	0x58, 0xd5, 0x5d, 0x8a, 0x32, 0x12, 0x02, 0xe7, 0x56, 0x6e, 0x91, 0xd6, 0xac, 0xea, 0x5a, 0x91,
	0xbe, 0x09, 0x83, 0x45, 0xb1, 0x8c, 0xb3, 0xf4, 0x2c, 0xfd, 0xfa, 0xbb, 0x5a, 0x3e, 0x43, 0xfb,
	0x18, 0xa3, 0x37, 0x6a, 0x17, 0xf1, 0x80, 0x4f, 0xde, 0x6c, 0xa5, 0xdf, 0x17, 0xfc, 0x38, 0xfe,
	0x87, 0xea, 0x19, 0xb4, 0x5a, 0x43, 0x93, 0x8b, 0x91, 0x3b, 0x9e, 0xfb, 0xf0, 0xd2, 0x87, 0x6f,
	0x30, 0x04, 0x39, 0xe0, 0x4b, 0xba, 0x0c, 0xf4, 0xeb, 0xf3, 0xa3, 0x67, 0x55, 0xb7, 0xb8, 0x07,
	0xae, 0x15, 0xcf, 0xc5, 0xc4, 0x31, 0xbd, 0xda, 0x40, 0x2b, 0x4f, 0x0f, 0xba, 0x9d, 0x51, 0x6b,
	0x83, 0xa3, 0x2e, 0xd0, 0x77, 0x81, 0xae, 0x0f, 0xd0, 0xa9, 0x87, 0xf8, 0x15, 0x54, 0x93, 0x92,
	0x0f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x51, 0x95, 0xa2, 0x76, 0x01, 0x00, 0x00,
}