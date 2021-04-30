// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/cors/v3/cors.proto

package envoy_extensions_filters_http_cors_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type Cors struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cors) Reset()         { *m = Cors{} }
func (m *Cors) String() string { return proto.CompactTextString(m) }
func (*Cors) ProtoMessage()    {}
func (*Cors) Descriptor() ([]byte, []int) {
	return fileDescriptor_b14462ebc0c27ba0, []int{0}
}

func (m *Cors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cors.Unmarshal(m, b)
}
func (m *Cors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cors.Marshal(b, m, deterministic)
}
func (m *Cors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cors.Merge(m, src)
}
func (m *Cors) XXX_Size() int {
	return xxx_messageInfo_Cors.Size(m)
}
func (m *Cors) XXX_DiscardUnknown() {
	xxx_messageInfo_Cors.DiscardUnknown(m)
}

var xxx_messageInfo_Cors proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Cors)(nil), "envoy.extensions.filters.http.cors.v3.Cors")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/cors/v3/cors.proto", fileDescriptor_b14462ebc0c27ba0)
}

var fileDescriptor_b14462ebc0c27ba0 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0xce, 0x2f, 0x2a, 0xd6, 0x2f,
	0x33, 0x06, 0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xaa, 0x60, 0x1d, 0x7a, 0x08, 0x1d,
	0x7a, 0x50, 0x1d, 0x7a, 0x20, 0x1d, 0x7a, 0x60, 0x95, 0x65, 0xc6, 0x52, 0xb2, 0xa5, 0x29, 0x05,
	0x89, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x60, 0x83, 0x8b, 0x4b, 0x12, 0x4b, 0x4a,
	0xa1, 0xa6, 0x48, 0x29, 0x62, 0x48, 0x97, 0xa5, 0x16, 0x81, 0x8c, 0xcb, 0xcc, 0x4b, 0x87, 0x28,
	0x51, 0x32, 0xe1, 0x62, 0x71, 0xce, 0x2f, 0x2a, 0xb6, 0xd2, 0x99, 0x75, 0xb4, 0x43, 0x4e, 0x9d,
	0x0b, 0x6a, 0x6f, 0x72, 0x7e, 0x5e, 0x5a, 0x66, 0x3a, 0xd4, 0x4e, 0x64, 0x2b, 0x8d, 0xf4, 0x40,
	0xaa, 0x9d, 0xbc, 0x76, 0x35, 0x9c, 0xb8, 0xc8, 0xc6, 0x24, 0xc0, 0xc4, 0x65, 0x9c, 0x99, 0xaf,
	0x07, 0xd6, 0x53, 0x50, 0x94, 0x5f, 0x51, 0xa9, 0x47, 0x94, 0xb3, 0x9d, 0x38, 0x41, 0x86, 0x04,
	0x80, 0xec, 0x0f, 0x60, 0x4c, 0x62, 0x03, 0x3b, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0x08, 0x01, 0x7d, 0x25, 0x01, 0x00, 0x00,
}