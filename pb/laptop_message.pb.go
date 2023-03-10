// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Laptop struct {
	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory   *Memory    `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storages []*Storage `protobuf:"bytes,7,rep,name=storages,proto3" json:"storages,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//	*Laptop_WeightKg
	//	*Laptop_WeightLbs
	Weight               isLaptop_Weight      `protobuf_oneof:"weight"`
	PriceUsd             float64              `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear          uint32               `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Laptop) Reset()         { *m = Laptop{} }
func (m *Laptop) String() string { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()    {}
func (*Laptop) Descriptor() ([]byte, []int) {
	return fileDescriptor_07a3824d46f4b731, []int{0}
}

func (m *Laptop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laptop.Unmarshal(m, b)
}
func (m *Laptop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laptop.Marshal(b, m, deterministic)
}
func (m *Laptop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laptop.Merge(m, src)
}
func (m *Laptop) XXX_Size() int {
	return xxx_messageInfo_Laptop.Size(m)
}
func (m *Laptop) XXX_DiscardUnknown() {
	xxx_messageInfo_Laptop.DiscardUnknown(m)
}

var xxx_messageInfo_Laptop proto.InternalMessageInfo

func (m *Laptop) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laptop) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Laptop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Laptop_WeightLbs struct {
	WeightLbs float64 `protobuf:"fixed64,11,opt,name=weight_lbs,json=weightLbs,proto3,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}

func (*Laptop_WeightLbs) isLaptop_Weight() {}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Laptop) GetWeightKg() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (m *Laptop) GetWeightLbs() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightLbs); ok {
		return x.WeightLbs
	}
	return 0
}

func (m *Laptop) GetPriceUsd() float64 {
	if m != nil {
		return m.PriceUsd
	}
	return 0
}

func (m *Laptop) GetReleaseYear() uint32 {
	if m != nil {
		return m.ReleaseYear
	}
	return 0
}

func (m *Laptop) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Laptop) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLbs)(nil),
	}
}

func init() {
	proto.RegisterType((*Laptop)(nil), "grpc.pcbook.Laptop")
}

func init() {
	proto.RegisterFile("laptop_message.proto", fileDescriptor_07a3824d46f4b731)
}

var fileDescriptor_07a3824d46f4b731 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xc9, 0xda, 0x95, 0xe4, 0x75, 0x9b, 0x90, 0xc9, 0xc0, 0x2a, 0x42, 0x0d, 0x13, 0x87,
	0x22, 0x24, 0x8f, 0x81, 0x38, 0x20, 0x4e, 0x8c, 0x03, 0x48, 0x1b, 0x52, 0xe5, 0xd1, 0x03, 0x5c,
	0x22, 0x3b, 0x79, 0x78, 0x51, 0x93, 0xda, 0xb2, 0x1d, 0x4d, 0xfd, 0x16, 0x7c, 0x64, 0x54, 0x27,
	0x9d, 0x68, 0xe1, 0x16, 0xff, 0x7f, 0xbf, 0x58, 0x7f, 0x3f, 0x3d, 0x48, 0x6b, 0x61, 0xbc, 0x36,
	0x79, 0x83, 0xce, 0x09, 0x85, 0xcc, 0x58, 0xed, 0x35, 0x19, 0x2b, 0x6b, 0x0a, 0x66, 0x0a, 0xa9,
	0xf5, 0x72, 0x92, 0x36, 0xd8, 0x68, 0xbb, 0xde, 0x55, 0x26, 0x4f, 0x8d, 0xd5, 0x05, 0x3a, 0xa7,
	0xed, 0x1e, 0x38, 0x75, 0x5e, 0x5b, 0xa1, 0x70, 0x2f, 0x4e, 0x5d, 0x61, 0x11, 0x57, 0x7b, 0xe9,
	0x93, 0x25, 0xae, 0xa5, 0x16, 0xb6, 0xdc, 0xcb, 0xa7, 0x4a, 0x6b, 0x55, 0xe3, 0x79, 0x38, 0xc9,
	0xf6, 0xd7, 0xb9, 0xaf, 0x1a, 0x74, 0x5e, 0x34, 0xa6, 0x13, 0xce, 0x7e, 0x0f, 0x61, 0x74, 0x1d,
	0xaa, 0x93, 0x13, 0x38, 0xa8, 0x4a, 0x1a, 0x65, 0xd1, 0x2c, 0xe1, 0x07, 0x55, 0x49, 0x52, 0x38,
	0x94, 0x56, 0xac, 0x4a, 0x7a, 0x10, 0xa2, 0xee, 0x40, 0x08, 0x0c, 0x57, 0xa2, 0x41, 0x3a, 0x08,
	0x61, 0xf8, 0x26, 0x67, 0x30, 0x28, 0x4c, 0x4b, 0x87, 0x59, 0x34, 0x1b, 0xbf, 0x7d, 0xc4, 0xfe,
	0x7a, 0x34, 0xfb, 0x3c, 0x5f, 0xf0, 0x0d, 0x24, 0xaf, 0x61, 0xd4, 0xbd, 0x9f, 0x1e, 0x06, 0xed,
	0xf1, 0x8e, 0xf6, 0x2d, 0x20, 0xde, 0x2b, 0xe4, 0x25, 0x0c, 0x95, 0x69, 0x1d, 0x1d, 0x65, 0x83,
	0x7f, 0x6e, 0xfc, 0x32, 0x5f, 0xf0, 0x40, 0xc9, 0x1b, 0x88, 0xfb, 0x19, 0x39, 0xfa, 0x30, 0x98,
	0xe9, 0x8e, 0x79, 0xd3, 0x41, 0x7e, 0x6f, 0x6d, 0x4a, 0x74, 0xe3, 0xa3, 0xf1, 0x7f, 0x4a, 0xdc,
	0x04, 0xc4, 0x7b, 0x85, 0x5c, 0x40, 0xbc, 0x9d, 0x2a, 0x4d, 0x82, 0x7e, 0xba, 0xa3, 0x5f, 0xf5,
	0x90, 0xdf, 0x6b, 0xe4, 0x39, 0x24, 0x77, 0x58, 0xa9, 0x5b, 0x9f, 0x2f, 0x15, 0x85, 0x2c, 0x9a,
	0x45, 0x5f, 0x1f, 0xf0, 0xb8, 0x8b, 0xae, 0x14, 0x99, 0x02, 0xf4, 0xb8, 0x96, 0x8e, 0x8e, 0x7b,
	0xde, 0xff, 0x72, 0x2d, 0x1d, 0x79, 0x06, 0x89, 0xb1, 0x55, 0x81, 0x79, 0xeb, 0x4a, 0x7a, 0xb4,
	0xe1, 0x3c, 0x0e, 0xc1, 0xc2, 0x95, 0xe4, 0x05, 0x1c, 0x59, 0xac, 0x51, 0x38, 0xcc, 0xd7, 0x28,
	0x2c, 0x3d, 0xce, 0xa2, 0xd9, 0x31, 0x1f, 0xf7, 0xd9, 0x0f, 0x14, 0x96, 0x7c, 0x00, 0x68, 0x4d,
	0x29, 0x3c, 0x96, 0xb9, 0xf0, 0xf4, 0x24, 0x94, 0x9e, 0xb0, 0x6e, 0x07, 0xd8, 0x76, 0x07, 0xd8,
	0xf7, 0xed, 0x0e, 0xf0, 0xa4, 0xb7, 0x3f, 0xf9, 0xcb, 0x18, 0x46, 0x5d, 0x8f, 0xcb, 0x57, 0x30,
	0x2d, 0x74, 0xc3, 0x54, 0xe5, 0x6f, 0x5b, 0xc9, 0x44, 0x2d, 0x5b, 0xa5, 0xef, 0xd6, 0x17, 0xef,
	0xb7, 0xef, 0x36, 0x72, 0x1e, 0xfd, 0x1c, 0xb2, 0x8f, 0x46, 0xca, 0x51, 0xb8, 0xf3, 0xdd, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x57, 0xb8, 0x3f, 0xfe, 0x02, 0x00, 0x00,
}
