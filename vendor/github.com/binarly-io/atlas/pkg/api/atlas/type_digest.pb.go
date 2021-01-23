// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_digest.proto

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

// DigestType represents all types of digests in the system.
type DigestType int32

const (
	// Due to first enum value has to be zero in proto3
	DigestType_DIGEST_RESERVED DigestType = 0
	// Unspecified means we do not know its type
	DigestType_DIGEST_UNSPECIFIED DigestType = 100
	// MD5 digest
	DigestType_DIGEST_MD5 DigestType = 200
	// SHA256 digest
	DigestType_DIGEST_SHA256 DigestType = 300
)

var DigestType_name = map[int32]string{
	0:   "DIGEST_RESERVED",
	100: "DIGEST_UNSPECIFIED",
	200: "DIGEST_MD5",
	300: "DIGEST_SHA256",
}

var DigestType_value = map[string]int32{
	"DIGEST_RESERVED":    0,
	"DIGEST_UNSPECIFIED": 100,
	"DIGEST_MD5":         200,
	"DIGEST_SHA256":      300,
}

func (x DigestType) String() string {
	return proto.EnumName(DigestType_name, int32(x))
}

func (DigestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1d7af31d4a49ad9, []int{0}
}

// Digest represents abstract digest.
type Digest struct {
	// Type of the digest. MD5 or SHA256 or something else
	Type DigestType `protobuf:"varint,100,opt,name=type,proto3,enum=atlas.DigestType" json:"type,omitempty"`
	// Data. Any arbitrary sequence of bytes no longer than 2^32
	Data                 []byte   `protobuf:"bytes,200,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Digest) Reset()         { *m = Digest{} }
func (m *Digest) String() string { return proto.CompactTextString(m) }
func (*Digest) ProtoMessage()    {}
func (*Digest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1d7af31d4a49ad9, []int{0}
}

func (m *Digest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Digest.Unmarshal(m, b)
}
func (m *Digest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Digest.Marshal(b, m, deterministic)
}
func (m *Digest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Digest.Merge(m, src)
}
func (m *Digest) XXX_Size() int {
	return xxx_messageInfo_Digest.Size(m)
}
func (m *Digest) XXX_DiscardUnknown() {
	xxx_messageInfo_Digest.DiscardUnknown(m)
}

var xxx_messageInfo_Digest proto.InternalMessageInfo

func (m *Digest) GetType() DigestType {
	if m != nil {
		return m.Type
	}
	return DigestType_DIGEST_RESERVED
}

func (m *Digest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("atlas.DigestType", DigestType_name, DigestType_value)
	proto.RegisterType((*Digest)(nil), "atlas.Digest")
}

func init() { proto.RegisterFile("type_digest.proto", fileDescriptor_e1d7af31d4a49ad9) }

var fileDescriptor_e1d7af31d4a49ad9 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xa9, 0x2c, 0x48,
	0x8d, 0x4f, 0xc9, 0x4c, 0x4f, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d,
	0x2c, 0xc9, 0x49, 0x2c, 0x56, 0x72, 0xe1, 0x62, 0x73, 0x01, 0x0b, 0x0b, 0xa9, 0x72, 0xb1, 0x80,
	0x54, 0x49, 0xa4, 0x28, 0x30, 0x6a, 0xf0, 0x19, 0x09, 0xea, 0x81, 0xe5, 0xf5, 0x20, 0x92, 0x21,
	0x95, 0x05, 0xa9, 0x41, 0x60, 0x69, 0x21, 0x61, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x89, 0x13,
	0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0x8e, 0x56, 0x1c, 0x17, 0x17, 0x42, 0xa1, 0x90, 0x30,
	0x17, 0xbf, 0x8b, 0xa7, 0xbb, 0x6b, 0x70, 0x48, 0x7c, 0x90, 0x6b, 0xb0, 0x6b, 0x50, 0x98, 0xab,
	0x8b, 0x00, 0x83, 0x90, 0x18, 0x97, 0x10, 0x54, 0x30, 0xd4, 0x2f, 0x38, 0xc0, 0xd5, 0xd9, 0xd3,
	0xcd, 0xd3, 0xd5, 0x45, 0x20, 0x45, 0x88, 0x9f, 0x8b, 0x0b, 0x2a, 0xee, 0xeb, 0x62, 0x2a, 0x70,
	0x82, 0x51, 0x48, 0x88, 0x8b, 0x17, 0x2a, 0x10, 0xec, 0xe1, 0x68, 0x64, 0x6a, 0x26, 0xb0, 0x86,
	0x29, 0x89, 0x0d, 0xec, 0x66, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xf0, 0x35, 0x41,
	0xc8, 0x00, 0x00, 0x00,
}
