// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_status_request_multi.proto

package atlas

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

type StatusRequestMode int32

const (
	StatusRequestMode_RESERVED        StatusRequestMode = 0
	StatusRequestMode_UNSPECIFIED     StatusRequestMode = 100
	StatusRequestMode_ALL             StatusRequestMode = 200
	StatusRequestMode_FIRST_FOUND     StatusRequestMode = 300
	StatusRequestMode_FIRST_NOT_FOUND StatusRequestMode = 400
)

var StatusRequestMode_name = map[int32]string{
	0:   "RESERVED",
	100: "UNSPECIFIED",
	200: "ALL",
	300: "FIRST_FOUND",
	400: "FIRST_NOT_FOUND",
}

var StatusRequestMode_value = map[string]int32{
	"RESERVED":        0,
	"UNSPECIFIED":     100,
	"ALL":             200,
	"FIRST_FOUND":     300,
	"FIRST_NOT_FOUND": 400,
}

func (x StatusRequestMode) String() string {
	return proto.EnumName(StatusRequestMode_name, int32(x))
}

func (StatusRequestMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_039e16d7649a7741, []int{0}
}

type StatusRequestMulti struct {
	Domain               *Domain           `protobuf:"bytes,100,opt,name=domain,proto3" json:"domain,omitempty"`
	Mode                 StatusRequestMode `protobuf:"varint,200,opt,name=mode,proto3,enum=atlas.StatusRequestMode" json:"mode,omitempty"`
	Entities             []*StatusRequest  `protobuf:"bytes,300,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StatusRequestMulti) Reset()         { *m = StatusRequestMulti{} }
func (m *StatusRequestMulti) String() string { return proto.CompactTextString(m) }
func (*StatusRequestMulti) ProtoMessage()    {}
func (*StatusRequestMulti) Descriptor() ([]byte, []int) {
	return fileDescriptor_039e16d7649a7741, []int{0}
}

func (m *StatusRequestMulti) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequestMulti.Unmarshal(m, b)
}
func (m *StatusRequestMulti) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequestMulti.Marshal(b, m, deterministic)
}
func (m *StatusRequestMulti) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequestMulti.Merge(m, src)
}
func (m *StatusRequestMulti) XXX_Size() int {
	return xxx_messageInfo_StatusRequestMulti.Size(m)
}
func (m *StatusRequestMulti) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequestMulti.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequestMulti proto.InternalMessageInfo

func (m *StatusRequestMulti) GetDomain() *Domain {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *StatusRequestMulti) GetMode() StatusRequestMode {
	if m != nil {
		return m.Mode
	}
	return StatusRequestMode_RESERVED
}

func (m *StatusRequestMulti) GetEntities() []*StatusRequest {
	if m != nil {
		return m.Entities
	}
	return nil
}

func init() {
	proto.RegisterEnum("atlas.StatusRequestMode", StatusRequestMode_name, StatusRequestMode_value)
	proto.RegisterType((*StatusRequestMulti)(nil), "atlas.StatusRequestMulti")
}

func init() { proto.RegisterFile("type_status_request_multi.proto", fileDescriptor_039e16d7649a7741) }

var fileDescriptor_039e16d7649a7741 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xa9, 0x2c, 0x48,
	0x8d, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x8e, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x89,
	0xcf, 0x2d, 0xcd, 0x29, 0xc9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9,
	0x49, 0x2c, 0x96, 0x12, 0x04, 0xab, 0x4b, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x83, 0xc8, 0x48, 0x49,
	0x62, 0xd1, 0x0a, 0x91, 0x52, 0x9a, 0xcb, 0xc8, 0x25, 0x14, 0x0c, 0x96, 0x08, 0x82, 0x88, 0xfb,
	0x82, 0x4c, 0x14, 0x52, 0xe5, 0x62, 0x83, 0x98, 0x20, 0x91, 0xa2, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4,
	0xab, 0x07, 0x36, 0x5c, 0xcf, 0x05, 0x2c, 0x18, 0x04, 0x95, 0x14, 0xd2, 0xe5, 0x62, 0xc9, 0xcd,
	0x4f, 0x49, 0x95, 0x38, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x67, 0x24, 0x01, 0x55, 0x85, 0x6a, 0x60,
	0x7e, 0x4a, 0x6a, 0x10, 0x58, 0x99, 0x90, 0x21, 0x17, 0x47, 0x6a, 0x5e, 0x49, 0x66, 0x49, 0x66,
	0x6a, 0xb1, 0xc4, 0x1a, 0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x11, 0x6c, 0x5a, 0x82, 0xe0, 0xca,
	0xb4, 0x52, 0xb8, 0x04, 0x31, 0x4c, 0x13, 0xe2, 0xe1, 0xe2, 0x08, 0x72, 0x0d, 0x76, 0x0d, 0x0a,
	0x73, 0x75, 0x11, 0x60, 0x10, 0xe2, 0xe7, 0xe2, 0x0e, 0xf5, 0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74,
	0xf3, 0x74, 0x75, 0x11, 0x48, 0x11, 0xe2, 0xe0, 0x62, 0x76, 0xf4, 0xf1, 0x11, 0x38, 0xc1, 0x28,
	0x24, 0xc0, 0xc5, 0xed, 0xe6, 0x19, 0x14, 0x1c, 0x12, 0xef, 0xe6, 0x1f, 0xea, 0xe7, 0x22, 0xb0,
	0x86, 0x49, 0x48, 0x84, 0x8b, 0x1f, 0x22, 0xe2, 0xe7, 0x0f, 0x13, 0x9d, 0xc0, 0x9c, 0xc4, 0x06,
	0x0e, 0x0c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x46, 0xfa, 0x1a, 0x64, 0x01, 0x00,
	0x00,
}
