// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filter_message.proto

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

type Filter struct {
	MaxPriceUsd          float64  `protobuf:"fixed64,1,opt,name=max_price_usd,json=maxPriceUsd,proto3" json:"max_price_usd,omitempty"`
	MinCpuCores          uint32   `protobuf:"varint,2,opt,name=min_cpu_cores,json=minCpuCores,proto3" json:"min_cpu_cores,omitempty"`
	MinCpuGhz            float64  `protobuf:"fixed64,3,opt,name=min_cpu_ghz,json=minCpuGhz,proto3" json:"min_cpu_ghz,omitempty"`
	MinRam               *Memory  `protobuf:"bytes,4,opt,name=min_ram,json=minRam,proto3" json:"min_ram,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_02dd12c5821a9fa1, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetMaxPriceUsd() float64 {
	if m != nil {
		return m.MaxPriceUsd
	}
	return 0
}

func (m *Filter) GetMinCpuCores() uint32 {
	if m != nil {
		return m.MinCpuCores
	}
	return 0
}

func (m *Filter) GetMinCpuGhz() float64 {
	if m != nil {
		return m.MinCpuGhz
	}
	return 0
}

func (m *Filter) GetMinRam() *Memory {
	if m != nil {
		return m.MinRam
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "grpc.pcbook.Filter")
}

func init() {
	proto.RegisterFile("filter_message.proto", fileDescriptor_02dd12c5821a9fa1)
}

var fileDescriptor_02dd12c5821a9fa1 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xcf, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0xa2, 0xa3, 0x62, 0xc2, 0x2e, 0x75, 0x87, 0xe2, 0x41, 0xcb, 0x4e, 0x15, 0x24,
	0xa0, 0xe2, 0xc9, 0x9b, 0x03, 0x3d, 0x09, 0xa3, 0xe0, 0xc5, 0x4b, 0x48, 0xb2, 0x98, 0x06, 0xf7,
	0x9a, 0x47, 0xd2, 0xe0, 0xb6, 0x0f, 0xe3, 0x67, 0x95, 0x46, 0x0a, 0xee, 0xfa, 0x7f, 0xbf, 0xf7,
	0xf8, 0x3f, 0xba, 0xf8, 0x74, 0xdb, 0xc1, 0x04, 0x01, 0x26, 0x46, 0x69, 0x0d, 0xc7, 0xe0, 0x07,
	0x5f, 0x32, 0x1b, 0x50, 0x73, 0xd4, 0xca, 0xfb, 0xaf, 0xcb, 0x05, 0x18, 0xf0, 0x61, 0x7f, 0x4c,
	0x96, 0x3f, 0x84, 0x16, 0x2f, 0x79, 0xb7, 0x5c, 0xd2, 0x39, 0xc8, 0x9d, 0xc0, 0xe0, 0xb4, 0x11,
	0x29, 0x6e, 0x2a, 0x52, 0x93, 0x86, 0xb4, 0x0c, 0xe4, 0x6e, 0x3d, 0x66, 0xef, 0x71, 0x93, 0x8d,
	0xeb, 0x85, 0xc6, 0x24, 0xb4, 0x0f, 0x26, 0x56, 0x27, 0x35, 0x69, 0xe6, 0x2d, 0x03, 0xd7, 0xaf,
	0x30, 0xad, 0xc6, 0xa8, 0xbc, 0xa2, 0x6c, 0x32, 0xb6, 0x3b, 0x54, 0xa7, 0xf9, 0xca, 0xf9, 0x9f,
	0x78, 0xed, 0x0e, 0xe5, 0x2d, 0x3d, 0x1b, 0xe7, 0x41, 0x42, 0x35, 0xab, 0x49, 0xc3, 0xee, 0x2f,
	0xf8, 0xbf, 0x9e, 0xfc, 0x2d, 0xd7, 0x6c, 0x0b, 0x70, 0x7d, 0x2b, 0xe1, 0xf9, 0x86, 0x5e, 0x6b,
	0x0f, 0xdc, 0xba, 0xa1, 0x4b, 0x8a, 0xcb, 0xad, 0x4a, 0xd6, 0x7f, 0xef, 0xef, 0x1e, 0x27, 0x8e,
	0x6a, 0x4d, 0x3e, 0x66, 0xfc, 0x09, 0x95, 0x2a, 0xf2, 0x4b, 0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xeb, 0xea, 0xad, 0x79, 0x0d, 0x01, 0x00, 0x00,
}
