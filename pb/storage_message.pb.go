// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage_message.proto

package pb

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

type Storage_Driver int32

const (
	Storage_UNKNOWN Storage_Driver = 0
	Storage_HDD     Storage_Driver = 1
	Storage_SSD     Storage_Driver = 2
)

var Storage_Driver_name = map[int32]string{
	0: "UNKNOWN",
	1: "HDD",
	2: "SSD",
}

var Storage_Driver_value = map[string]int32{
	"UNKNOWN": 0,
	"HDD":     1,
	"SSD":     2,
}

func (x Storage_Driver) String() string {
	return proto.EnumName(Storage_Driver_name, int32(x))
}

func (Storage_Driver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_170f09d838bd8a04, []int{0, 0}
}

type Storage struct {
	Driver               Storage_Driver `protobuf:"varint,1,opt,name=driver,proto3,enum=grpc.pcbook.Storage_Driver" json:"driver,omitempty"`
	Memory               *Memory        `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_170f09d838bd8a04, []int{0}
}

func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetDriver() Storage_Driver {
	if m != nil {
		return m.Driver
	}
	return Storage_UNKNOWN
}

func (m *Storage) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterEnum("grpc.pcbook.Storage_Driver", Storage_Driver_name, Storage_Driver_value)
	proto.RegisterType((*Storage)(nil), "grpc.pcbook.Storage")
}

func init() {
	proto.RegisterFile("storage_message.proto", fileDescriptor_170f09d838bd8a04)
}

var fileDescriptor_170f09d838bd8a04 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0x8d, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x4e, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x48, 0x4e, 0xca, 0xcf, 0xcf, 0x96, 0x12,
	0xc9, 0x4d, 0xcd, 0xcd, 0x2f, 0xaa, 0x44, 0x55, 0xa2, 0x34, 0x85, 0x91, 0x8b, 0x3d, 0x18, 0xa2,
	0x59, 0xc8, 0x98, 0x8b, 0x2d, 0xa5, 0x28, 0xb3, 0x2c, 0xb5, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xcf, 0x48, 0x5a, 0x0f, 0x49, 0xbf, 0x1e, 0x54, 0x95, 0x9e, 0x0b, 0x58, 0x49, 0x10, 0x54, 0xa9,
	0x90, 0x36, 0x17, 0x1b, 0xc4, 0x60, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x61, 0x14, 0x4d,
	0xbe, 0x60, 0xa9, 0x20, 0xa8, 0x12, 0x25, 0x75, 0x2e, 0x36, 0x88, 0x76, 0x21, 0x6e, 0x2e, 0xf6,
	0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f, 0x01, 0x06, 0x21, 0x76, 0x2e, 0x66, 0x0f, 0x17, 0x17,
	0x01, 0x46, 0x10, 0x23, 0x38, 0xd8, 0x45, 0x80, 0xc9, 0x49, 0x93, 0x4b, 0x3e, 0x39, 0x3f, 0x57,
	0x2f, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x31, 0x27, 0xa9, 0x34, 0x3d, 0xbf, 0xbc, 0xd2,
	0xd0, 0x14, 0x66, 0x6e, 0x41, 0x52, 0x00, 0x63, 0x14, 0x8b, 0x9e, 0x75, 0x41, 0x52, 0x12, 0x1b,
	0xd8, 0x23, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x20, 0x23, 0xae, 0x04, 0x01, 0x00,
	0x00,
}
