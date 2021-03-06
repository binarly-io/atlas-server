// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_user_id.proto

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

// Unique identifier of the user.
type UserID struct {
	// Any arbitrary sequence of bytes no longer than 2^32
	Data                 []byte   `protobuf:"bytes,100,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserID) Reset()         { *m = UserID{} }
func (m *UserID) String() string { return proto.CompactTextString(m) }
func (*UserID) ProtoMessage()    {}
func (*UserID) Descriptor() ([]byte, []int) {
	return fileDescriptor_df8c113849595fd9, []int{0}
}

func (m *UserID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserID.Unmarshal(m, b)
}
func (m *UserID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserID.Marshal(b, m, deterministic)
}
func (m *UserID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserID.Merge(m, src)
}
func (m *UserID) XXX_Size() int {
	return xxx_messageInfo_UserID.Size(m)
}
func (m *UserID) XXX_DiscardUnknown() {
	xxx_messageInfo_UserID.DiscardUnknown(m)
}

var xxx_messageInfo_UserID proto.InternalMessageInfo

func (m *UserID) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*UserID)(nil), "atlas.UserID")
}

func init() { proto.RegisterFile("type_user_id.proto", fileDescriptor_df8c113849595fd9) }

var fileDescriptor_df8c113849595fd9 = []byte{
	// 83 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xa9, 0x2c, 0x48,
	0x8d, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a, 0xcf, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x56, 0x92, 0xe1, 0x62, 0x0b, 0x2d, 0x4e, 0x2d, 0xf2, 0x74, 0x11,
	0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x48, 0x51, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3,
	0x93, 0xd8, 0xc0, 0x6a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x65, 0x62, 0x03, 0x41,
	0x00, 0x00, 0x00,
}
