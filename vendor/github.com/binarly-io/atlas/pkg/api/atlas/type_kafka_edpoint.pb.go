// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_kafka_edpoint.proto

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

// KafkaEndpoint represents Kafka cluster access endpoint
type KafkaEndpoint struct {
	// Multiple brokers of Kafka's cluster
	Brokers              []string `protobuf:"bytes,100,rep,name=brokers,proto3" json:"brokers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaEndpoint) Reset()         { *m = KafkaEndpoint{} }
func (m *KafkaEndpoint) String() string { return proto.CompactTextString(m) }
func (*KafkaEndpoint) ProtoMessage()    {}
func (*KafkaEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_93eebabfa985a8d4, []int{0}
}

func (m *KafkaEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaEndpoint.Unmarshal(m, b)
}
func (m *KafkaEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaEndpoint.Marshal(b, m, deterministic)
}
func (m *KafkaEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaEndpoint.Merge(m, src)
}
func (m *KafkaEndpoint) XXX_Size() int {
	return xxx_messageInfo_KafkaEndpoint.Size(m)
}
func (m *KafkaEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaEndpoint proto.InternalMessageInfo

func (m *KafkaEndpoint) GetBrokers() []string {
	if m != nil {
		return m.Brokers
	}
	return nil
}

func init() {
	proto.RegisterType((*KafkaEndpoint)(nil), "atlas.KafkaEndpoint")
}

func init() { proto.RegisterFile("type_kafka_edpoint.proto", fileDescriptor_93eebabfa985a8d4) }

var fileDescriptor_93eebabfa985a8d4 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0xa9, 0x2c, 0x48,
	0x8d, 0xcf, 0x4e, 0x4c, 0xcb, 0x4e, 0x8c, 0x4f, 0x4d, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x56, 0xd2, 0xe4, 0xe2, 0xf5,
	0x06, 0xc9, 0xba, 0xe6, 0x41, 0x64, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x8a, 0xf2, 0xb3, 0x53, 0x8b,
	0x8a, 0x25, 0x52, 0x14, 0x98, 0x35, 0x38, 0x83, 0x60, 0xdc, 0x24, 0x36, 0xb0, 0x46, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x01, 0xe1, 0xf4, 0x54, 0x00, 0x00, 0x00,
}
