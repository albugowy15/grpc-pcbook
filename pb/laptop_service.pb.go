// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_service.proto

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

type CreateLaptopRequest struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopRequest) Reset()         { *m = CreateLaptopRequest{} }
func (m *CreateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopRequest) ProtoMessage()    {}
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{0}
}

func (m *CreateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopRequest.Unmarshal(m, b)
}
func (m *CreateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *CreateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopRequest.Merge(m, src)
}
func (m *CreateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopRequest.Size(m)
}
func (m *CreateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopRequest proto.InternalMessageInfo

func (m *CreateLaptopRequest) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopResponse) Reset()         { *m = CreateLaptopResponse{} }
func (m *CreateLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopResponse) ProtoMessage()    {}
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{1}
}

func (m *CreateLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopResponse.Unmarshal(m, b)
}
func (m *CreateLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopResponse.Marshal(b, m, deterministic)
}
func (m *CreateLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopResponse.Merge(m, src)
}
func (m *CreateLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopResponse.Size(m)
}
func (m *CreateLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopResponse proto.InternalMessageInfo

func (m *CreateLaptopResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type SearchLaptopRequest struct {
	Filter               *Filter  `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchLaptopRequest) Reset()         { *m = SearchLaptopRequest{} }
func (m *SearchLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*SearchLaptopRequest) ProtoMessage()    {}
func (*SearchLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{2}
}

func (m *SearchLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchLaptopRequest.Unmarshal(m, b)
}
func (m *SearchLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchLaptopRequest.Marshal(b, m, deterministic)
}
func (m *SearchLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchLaptopRequest.Merge(m, src)
}
func (m *SearchLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_SearchLaptopRequest.Size(m)
}
func (m *SearchLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchLaptopRequest proto.InternalMessageInfo

func (m *SearchLaptopRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

type SearchLaptopResponse struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchLaptopResponse) Reset()         { *m = SearchLaptopResponse{} }
func (m *SearchLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*SearchLaptopResponse) ProtoMessage()    {}
func (*SearchLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{3}
}

func (m *SearchLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchLaptopResponse.Unmarshal(m, b)
}
func (m *SearchLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchLaptopResponse.Marshal(b, m, deterministic)
}
func (m *SearchLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchLaptopResponse.Merge(m, src)
}
func (m *SearchLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_SearchLaptopResponse.Size(m)
}
func (m *SearchLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchLaptopResponse proto.InternalMessageInfo

func (m *SearchLaptopResponse) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateLaptopRequest)(nil), "grpc.pcbook.CreateLaptopRequest")
	proto.RegisterType((*CreateLaptopResponse)(nil), "grpc.pcbook.CreateLaptopResponse")
	proto.RegisterType((*SearchLaptopRequest)(nil), "grpc.pcbook.SearchLaptopRequest")
	proto.RegisterType((*SearchLaptopResponse)(nil), "grpc.pcbook.SearchLaptopResponse")
}

func init() {
	proto.RegisterFile("laptop_service.proto", fileDescriptor_240c60d9fb227e71)
}

var fileDescriptor_240c60d9fb227e71 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x49, 0x2c, 0x28,
	0xc9, 0x2f, 0x88, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x4e, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x48, 0x4e, 0xca, 0xcf, 0xcf, 0x96, 0x82, 0x29,
	0xc9, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x87, 0x2a, 0x91, 0x12, 0x49, 0xcb, 0xcc, 0x29, 0x49, 0x2d,
	0x42, 0x15, 0x55, 0x72, 0xe2, 0x12, 0x76, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0xf5, 0x01, 0xeb, 0x09,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xd2, 0xe6, 0x62, 0x83, 0x18, 0x22, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x6d, 0x24, 0xac, 0x87, 0x64, 0x81, 0x1e, 0x54, 0x2d, 0x54, 0x89, 0x92, 0x1a, 0x97,
	0x08, 0xaa, 0x19, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x60,
	0x03, 0x38, 0x83, 0x98, 0x32, 0x53, 0x40, 0x76, 0x05, 0xa7, 0x26, 0x16, 0x25, 0x67, 0x60, 0xd8,
	0x05, 0x71, 0x1a, 0x56, 0xbb, 0xdc, 0xc0, 0x52, 0x41, 0x50, 0x25, 0x4a, 0xce, 0x5c, 0x22, 0xa8,
	0x66, 0x40, 0xed, 0x22, 0xc5, 0xc1, 0x46, 0xfb, 0x19, 0xb9, 0x78, 0x21, 0x42, 0xc1, 0x90, 0x50,
	0x14, 0x0a, 0xe5, 0xe2, 0x41, 0xf6, 0x82, 0x90, 0x02, 0x8a, 0x76, 0x2c, 0x21, 0x24, 0xa5, 0x88,
	0x47, 0x05, 0xc4, 0x4d, 0x4a, 0x0c, 0x42, 0xe1, 0x5c, 0x3c, 0xc8, 0xae, 0x45, 0x33, 0x16, 0x4b,
	0x60, 0xa0, 0x19, 0x8b, 0xcd, 0xab, 0x4a, 0x0c, 0x06, 0x8c, 0x4e, 0x9a, 0x5c, 0xf2, 0xc9, 0xf9,
	0xb9, 0x7a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0x89, 0x39, 0x49, 0xa5, 0xe9, 0xf9, 0xe5,
	0x95, 0x86, 0xa6, 0x30, 0x8d, 0x05, 0x49, 0x01, 0x8c, 0x51, 0x2c, 0x7a, 0xd6, 0x05, 0x49, 0x49,
	0x6c, 0xe0, 0x88, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x68, 0x5e, 0x3a, 0x23, 0x39, 0x02,
	0x00, 0x00,
}
