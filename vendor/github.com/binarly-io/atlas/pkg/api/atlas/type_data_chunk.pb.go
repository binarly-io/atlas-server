// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type_data_chunk.proto

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

type DataChunkType int32

const (
	// Due to first enum value has to be zero in proto3
	DataChunkType_DATA_CHUNK_TYPE_RESERVED    DataChunkType = 0
	DataChunkType_DATA_CHUNK_TYPE_UNSPECIFIED DataChunkType = 100
	DataChunkType_DATA_CHUNK_TYPE_DATA        DataChunkType = 200
)

var DataChunkType_name = map[int32]string{
	0:   "DATA_CHUNK_TYPE_RESERVED",
	100: "DATA_CHUNK_TYPE_UNSPECIFIED",
	200: "DATA_CHUNK_TYPE_DATA",
}

var DataChunkType_value = map[string]int32{
	"DATA_CHUNK_TYPE_RESERVED":    0,
	"DATA_CHUNK_TYPE_UNSPECIFIED": 100,
	"DATA_CHUNK_TYPE_DATA":        200,
}

func (x DataChunkType) String() string {
	return proto.EnumName(DataChunkType_name, int32(x))
}

func (DataChunkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c69e33da091858ca, []int{0}
}

// DataChunk is a chunk of data transferred as a single piece.
// Can be part of bigger data, transferred by chunks
type DataChunk struct {
	// Header describes this chunk
	Header *Metadata `protobuf:"bytes,100,opt,name=header,proto3" json:"header,omitempty"`
	// Bytes is the purpose of the whole data chunk type
	// May contain any arbitrary sequence of bytes no longer than 2^32
	Bytes []byte `protobuf:"bytes,200,opt,name=bytes,proto3" json:"bytes,omitempty"`
	// Types that are valid to be assigned to TransportMetadataOptional:
	//	*DataChunk_TransportMetadata
	TransportMetadataOptional isDataChunk_TransportMetadataOptional `protobuf_oneof:"transport_metadata_optional"`
	// Types that are valid to be assigned to PayloadMetadataOptional:
	//	*DataChunk_PayloadMetadata
	PayloadMetadataOptional isDataChunk_PayloadMetadataOptional `protobuf_oneof:"payload_metadata_optional"`
	XXX_NoUnkeyedLiteral    struct{}                            `json:"-"`
	XXX_unrecognized        []byte                              `json:"-"`
	XXX_sizecache           int32                               `json:"-"`
}

func (m *DataChunk) Reset()         { *m = DataChunk{} }
func (m *DataChunk) String() string { return proto.CompactTextString(m) }
func (*DataChunk) ProtoMessage()    {}
func (*DataChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69e33da091858ca, []int{0}
}

func (m *DataChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataChunk.Unmarshal(m, b)
}
func (m *DataChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataChunk.Marshal(b, m, deterministic)
}
func (m *DataChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataChunk.Merge(m, src)
}
func (m *DataChunk) XXX_Size() int {
	return xxx_messageInfo_DataChunk.Size(m)
}
func (m *DataChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_DataChunk.DiscardUnknown(m)
}

var xxx_messageInfo_DataChunk proto.InternalMessageInfo

func (m *DataChunk) GetHeader() *Metadata {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *DataChunk) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

type isDataChunk_TransportMetadataOptional interface {
	isDataChunk_TransportMetadataOptional()
}

type DataChunk_TransportMetadata struct {
	TransportMetadata *Metadata `protobuf:"bytes,300,opt,name=transport_metadata,json=transportMetadata,proto3,oneof"`
}

func (*DataChunk_TransportMetadata) isDataChunk_TransportMetadataOptional() {}

func (m *DataChunk) GetTransportMetadataOptional() isDataChunk_TransportMetadataOptional {
	if m != nil {
		return m.TransportMetadataOptional
	}
	return nil
}

func (m *DataChunk) GetTransportMetadata() *Metadata {
	if x, ok := m.GetTransportMetadataOptional().(*DataChunk_TransportMetadata); ok {
		return x.TransportMetadata
	}
	return nil
}

type isDataChunk_PayloadMetadataOptional interface {
	isDataChunk_PayloadMetadataOptional()
}

type DataChunk_PayloadMetadata struct {
	PayloadMetadata *Metadata `protobuf:"bytes,400,opt,name=payload_metadata,json=payloadMetadata,proto3,oneof"`
}

func (*DataChunk_PayloadMetadata) isDataChunk_PayloadMetadataOptional() {}

func (m *DataChunk) GetPayloadMetadataOptional() isDataChunk_PayloadMetadataOptional {
	if m != nil {
		return m.PayloadMetadataOptional
	}
	return nil
}

func (m *DataChunk) GetPayloadMetadata() *Metadata {
	if x, ok := m.GetPayloadMetadataOptional().(*DataChunk_PayloadMetadata); ok {
		return x.PayloadMetadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DataChunk) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DataChunk_TransportMetadata)(nil),
		(*DataChunk_PayloadMetadata)(nil),
	}
}

func init() {
	proto.RegisterEnum("atlas.DataChunkType", DataChunkType_name, DataChunkType_value)
	proto.RegisterType((*DataChunk)(nil), "atlas.DataChunk")
}

func init() {
	proto.RegisterFile("type_data_chunk.proto", fileDescriptor_c69e33da091858ca)
}

var fileDescriptor_c69e33da091858ca = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0xa9, 0x2c, 0x48,
	0x8d, 0x4f, 0x49, 0x2c, 0x49, 0x8c, 0x4f, 0xce, 0x28, 0xcd, 0xcb, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x4d, 0x2c, 0xc9, 0x49, 0x2c, 0x96, 0x12, 0x06, 0xcb, 0xe6, 0xa6, 0x96, 0x24,
	0x82, 0x54, 0x40, 0xe4, 0x94, 0x3a, 0x99, 0xb8, 0x38, 0x5d, 0x12, 0x4b, 0x12, 0x9d, 0x41, 0xea,
	0x85, 0xd4, 0xb9, 0xd8, 0x32, 0x52, 0x13, 0x53, 0x52, 0x8b, 0x24, 0x52, 0x14, 0x18, 0x35, 0xb8,
	0x8d, 0xf8, 0xf5, 0xc0, 0x5a, 0xf5, 0x7c, 0xa1, 0x9a, 0x82, 0xa0, 0xd2, 0x42, 0xa2, 0x5c, 0xac,
	0x49, 0x95, 0x25, 0xa9, 0xc5, 0x12, 0x27, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x20, 0x3c, 0x21,
	0x47, 0x2e, 0xa1, 0x92, 0xa2, 0xc4, 0xbc, 0xe2, 0x82, 0xfc, 0xa2, 0x12, 0xb8, 0x4d, 0x12, 0x6b,
	0x98, 0xb0, 0x1a, 0xe6, 0xc1, 0x10, 0x24, 0x08, 0x57, 0x0d, 0x13, 0x14, 0xb2, 0xe5, 0x12, 0x28,
	0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x41, 0x18, 0x30, 0x81, 0x19, 0xbb, 0x01, 0x8c, 0x41, 0xfc,
	0x50, 0xb5, 0x30, 0x21, 0x27, 0x59, 0x2e, 0x69, 0x4c, 0x17, 0xc4, 0xe7, 0x17, 0x94, 0x64, 0xe6,
	0xe7, 0x25, 0xe6, 0x38, 0x49, 0x73, 0x49, 0xa2, 0x9b, 0x0e, 0x97, 0xd4, 0xca, 0xe4, 0xe2, 0x85,
	0x07, 0x45, 0x48, 0x65, 0x41, 0xaa, 0x90, 0x0c, 0x97, 0x84, 0x8b, 0x63, 0x88, 0x63, 0xbc, 0xb3,
	0x47, 0xa8, 0x9f, 0x77, 0x7c, 0x48, 0x64, 0x80, 0x6b, 0x7c, 0x90, 0x6b, 0xb0, 0x6b, 0x50, 0x98,
	0xab, 0x8b, 0x00, 0x83, 0x90, 0x3c, 0x97, 0x34, 0xba, 0x6c, 0xa8, 0x5f, 0x70, 0x80, 0xab, 0xb3,
	0xa7, 0x9b, 0xa7, 0xab, 0x8b, 0x40, 0x8a, 0x90, 0x24, 0x97, 0x08, 0xba, 0x02, 0x10, 0x5f, 0xe0,
	0x04, 0x63, 0x12, 0x1b, 0x38, 0xf4, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x65, 0x21,
	0x80, 0xb2, 0x01, 0x00, 0x00,
}
