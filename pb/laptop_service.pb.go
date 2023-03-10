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

type UploadImageRequest struct {
	// Types that are valid to be assigned to Data:
	//	*UploadImageRequest_Info
	//	*UploadImageRequest_ChunkData
	Data                 isUploadImageRequest_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UploadImageRequest) Reset()         { *m = UploadImageRequest{} }
func (m *UploadImageRequest) String() string { return proto.CompactTextString(m) }
func (*UploadImageRequest) ProtoMessage()    {}
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{4}
}

func (m *UploadImageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageRequest.Unmarshal(m, b)
}
func (m *UploadImageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageRequest.Marshal(b, m, deterministic)
}
func (m *UploadImageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageRequest.Merge(m, src)
}
func (m *UploadImageRequest) XXX_Size() int {
	return xxx_messageInfo_UploadImageRequest.Size(m)
}
func (m *UploadImageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageRequest proto.InternalMessageInfo

type isUploadImageRequest_Data interface {
	isUploadImageRequest_Data()
}

type UploadImageRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadImageRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadImageRequest_Info) isUploadImageRequest_Data() {}

func (*UploadImageRequest_ChunkData) isUploadImageRequest_Data() {}

func (m *UploadImageRequest) GetData() isUploadImageRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UploadImageRequest) GetInfo() *ImageInfo {
	if x, ok := m.GetData().(*UploadImageRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (m *UploadImageRequest) GetChunkData() []byte {
	if x, ok := m.GetData().(*UploadImageRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UploadImageRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UploadImageRequest_Info)(nil),
		(*UploadImageRequest_ChunkData)(nil),
	}
}

type ImageInfo struct {
	LaptopId             string   `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	ImageType            string   `protobuf:"bytes,2,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageInfo) Reset()         { *m = ImageInfo{} }
func (m *ImageInfo) String() string { return proto.CompactTextString(m) }
func (*ImageInfo) ProtoMessage()    {}
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{5}
}

func (m *ImageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageInfo.Unmarshal(m, b)
}
func (m *ImageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageInfo.Marshal(b, m, deterministic)
}
func (m *ImageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageInfo.Merge(m, src)
}
func (m *ImageInfo) XXX_Size() int {
	return xxx_messageInfo_ImageInfo.Size(m)
}
func (m *ImageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ImageInfo proto.InternalMessageInfo

func (m *ImageInfo) GetLaptopId() string {
	if m != nil {
		return m.LaptopId
	}
	return ""
}

func (m *ImageInfo) GetImageType() string {
	if m != nil {
		return m.ImageType
	}
	return ""
}

type UploadImageResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Size                 uint32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadImageResponse) Reset()         { *m = UploadImageResponse{} }
func (m *UploadImageResponse) String() string { return proto.CompactTextString(m) }
func (*UploadImageResponse) ProtoMessage()    {}
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{6}
}

func (m *UploadImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadImageResponse.Unmarshal(m, b)
}
func (m *UploadImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadImageResponse.Marshal(b, m, deterministic)
}
func (m *UploadImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadImageResponse.Merge(m, src)
}
func (m *UploadImageResponse) XXX_Size() int {
	return xxx_messageInfo_UploadImageResponse.Size(m)
}
func (m *UploadImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadImageResponse proto.InternalMessageInfo

func (m *UploadImageResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UploadImageResponse) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type RateLaptopRequest struct {
	LaptopId             string   `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	Score                float64  `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLaptopRequest) Reset()         { *m = RateLaptopRequest{} }
func (m *RateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*RateLaptopRequest) ProtoMessage()    {}
func (*RateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{7}
}

func (m *RateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLaptopRequest.Unmarshal(m, b)
}
func (m *RateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *RateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLaptopRequest.Merge(m, src)
}
func (m *RateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_RateLaptopRequest.Size(m)
}
func (m *RateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLaptopRequest proto.InternalMessageInfo

func (m *RateLaptopRequest) GetLaptopId() string {
	if m != nil {
		return m.LaptopId
	}
	return ""
}

func (m *RateLaptopRequest) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}

type RateLaptopResponse struct {
	LaptopId             string   `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	RatedCount           uint32   `protobuf:"varint,2,opt,name=rated_count,json=ratedCount,proto3" json:"rated_count,omitempty"`
	AverageScore         float64  `protobuf:"fixed64,3,opt,name=average_score,json=averageScore,proto3" json:"average_score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLaptopResponse) Reset()         { *m = RateLaptopResponse{} }
func (m *RateLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*RateLaptopResponse) ProtoMessage()    {}
func (*RateLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{8}
}

func (m *RateLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLaptopResponse.Unmarshal(m, b)
}
func (m *RateLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLaptopResponse.Marshal(b, m, deterministic)
}
func (m *RateLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLaptopResponse.Merge(m, src)
}
func (m *RateLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_RateLaptopResponse.Size(m)
}
func (m *RateLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLaptopResponse proto.InternalMessageInfo

func (m *RateLaptopResponse) GetLaptopId() string {
	if m != nil {
		return m.LaptopId
	}
	return ""
}

func (m *RateLaptopResponse) GetRatedCount() uint32 {
	if m != nil {
		return m.RatedCount
	}
	return 0
}

func (m *RateLaptopResponse) GetAverageScore() float64 {
	if m != nil {
		return m.AverageScore
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateLaptopRequest)(nil), "grpc.pcbook.CreateLaptopRequest")
	proto.RegisterType((*CreateLaptopResponse)(nil), "grpc.pcbook.CreateLaptopResponse")
	proto.RegisterType((*SearchLaptopRequest)(nil), "grpc.pcbook.SearchLaptopRequest")
	proto.RegisterType((*SearchLaptopResponse)(nil), "grpc.pcbook.SearchLaptopResponse")
	proto.RegisterType((*UploadImageRequest)(nil), "grpc.pcbook.UploadImageRequest")
	proto.RegisterType((*ImageInfo)(nil), "grpc.pcbook.ImageInfo")
	proto.RegisterType((*UploadImageResponse)(nil), "grpc.pcbook.UploadImageResponse")
	proto.RegisterType((*RateLaptopRequest)(nil), "grpc.pcbook.RateLaptopRequest")
	proto.RegisterType((*RateLaptopResponse)(nil), "grpc.pcbook.RateLaptopResponse")
}

func init() {
	proto.RegisterFile("laptop_service.proto", fileDescriptor_240c60d9fb227e71)
}

var fileDescriptor_240c60d9fb227e71 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x51, 0x8b, 0xd3, 0x40,
	0x18, 0x6c, 0x6a, 0x2d, 0xe6, 0x6b, 0x2b, 0xf8, 0x35, 0xc8, 0x11, 0xd1, 0xd4, 0x08, 0x52, 0x51,
	0xc2, 0x79, 0xe2, 0x83, 0xf8, 0xd6, 0xca, 0x79, 0x05, 0x1f, 0x74, 0xeb, 0x21, 0xf8, 0x52, 0x36,
	0xc9, 0x36, 0x0d, 0x6d, 0xb3, 0x6b, 0xb2, 0xb9, 0xa3, 0xfe, 0x60, 0x7f, 0x87, 0x64, 0x37, 0x3d,
	0x9b, 0x4b, 0xae, 0xe0, 0x5b, 0x3b, 0x3b, 0xdf, 0xcc, 0xb7, 0x3b, 0x43, 0xc0, 0xda, 0x50, 0x21,
	0xb9, 0x58, 0x64, 0x2c, 0xbd, 0x8a, 0x03, 0xe6, 0x89, 0x94, 0x4b, 0x8e, 0xbd, 0x28, 0x15, 0x81,
	0x27, 0x02, 0x9f, 0xf3, 0xb5, 0xbd, 0xa7, 0x6c, 0x59, 0x96, 0xd1, 0xa8, 0xa4, 0xd8, 0xd6, 0x32,
	0xde, 0x48, 0x96, 0x56, 0x51, 0x77, 0x02, 0xc3, 0x69, 0xca, 0xa8, 0x64, 0x5f, 0xd4, 0x0c, 0x61,
	0xbf, 0x72, 0x96, 0x49, 0x7c, 0x0d, 0x5d, 0x2d, 0x72, 0x62, 0x8c, 0x8c, 0x71, 0xef, 0x6c, 0xe8,
	0x1d, 0x18, 0x78, 0x25, 0xb7, 0xa4, 0xb8, 0x2f, 0xc1, 0xaa, 0x6a, 0x64, 0x82, 0x27, 0x19, 0xc3,
	0x87, 0xd0, 0x8e, 0x43, 0x25, 0x60, 0x92, 0x76, 0x1c, 0x16, 0x5e, 0x73, 0x46, 0xd3, 0x60, 0x55,
	0xf3, 0xd2, 0xab, 0x35, 0x7a, 0x9d, 0xab, 0x23, 0x52, 0x52, 0xdc, 0x29, 0x58, 0x55, 0x8d, 0xd2,
	0xeb, 0xbf, 0x16, 0x5e, 0x03, 0x5e, 0x8a, 0x0d, 0xa7, 0xe1, 0x6c, 0x4b, 0x23, 0xb6, 0xdf, 0xe3,
	0x0d, 0x74, 0xe2, 0x64, 0xc9, 0x4b, 0x81, 0xc7, 0x15, 0x01, 0x45, 0x9c, 0x25, 0x4b, 0x7e, 0xd1,
	0x22, 0x8a, 0x85, 0x0e, 0x40, 0xb0, 0xca, 0x93, 0xf5, 0x22, 0xa4, 0x92, 0x9e, 0xb4, 0x47, 0xc6,
	0xb8, 0x7f, 0xd1, 0x22, 0xa6, 0xc2, 0x3e, 0x51, 0x49, 0x27, 0x5d, 0xe8, 0x14, 0x47, 0xee, 0x67,
	0x30, 0x6f, 0xa6, 0xf1, 0x09, 0x98, 0x65, 0x38, 0x37, 0x2f, 0xf3, 0x40, 0x03, 0xb3, 0x10, 0x9f,
	0x02, 0xc4, 0x05, 0x73, 0x21, 0x77, 0x82, 0x29, 0x49, 0x93, 0x98, 0x0a, 0xf9, 0xbe, 0x13, 0xcc,
	0xfd, 0x00, 0xc3, 0xca, 0xd6, 0xcd, 0xaf, 0x8c, 0x08, 0x9d, 0x2c, 0xfe, 0xad, 0xe7, 0x07, 0x44,
	0xfd, 0x76, 0xcf, 0xe1, 0x11, 0xa9, 0x65, 0x7c, 0x74, 0x17, 0x0b, 0xee, 0x67, 0x01, 0x4f, 0xb5,
	0x8c, 0x41, 0xf4, 0x1f, 0xf7, 0x1a, 0x90, 0xd4, 0x73, 0x3e, 0x2a, 0xe4, 0x40, 0x2f, 0xa5, 0x92,
	0x85, 0x8b, 0x80, 0xe7, 0x89, 0x2c, 0xb7, 0x02, 0x05, 0x4d, 0x0b, 0x04, 0x5f, 0xc0, 0x80, 0x5e,
	0xb1, 0xb4, 0xb8, 0xb7, 0x76, 0xbc, 0xa7, 0x1c, 0xfb, 0x25, 0x38, 0x2f, 0xb0, 0xb3, 0x3f, 0x6d,
	0x18, 0x68, 0xd7, 0xb9, 0xee, 0x3d, 0x5e, 0x42, 0xff, 0xb0, 0x74, 0x38, 0xaa, 0xe4, 0xd5, 0xd0,
	0x69, 0xfb, 0xf9, 0x11, 0x86, 0xbe, 0x89, 0xdb, 0xc2, 0x1f, 0xd0, 0x3f, 0xec, 0xd7, 0x2d, 0xd9,
	0x86, 0xfa, 0xde, 0x92, 0x6d, 0x2a, 0xa7, 0xdb, 0x3a, 0x35, 0x90, 0x40, 0xef, 0x20, 0x3d, 0x74,
	0x2a, 0x53, 0xf5, 0x36, 0xda, 0xa3, 0xbb, 0x09, 0x5a, 0x75, 0x6c, 0xe0, 0x37, 0x80, 0x7f, 0x71,
	0xe0, 0xb3, 0xca, 0x44, 0x2d, 0x6f, 0xdb, 0xb9, 0xf3, 0x7c, 0x2f, 0x78, 0x6a, 0x4c, 0x5e, 0x81,
	0x13, 0xf0, 0xad, 0x17, 0xc5, 0x72, 0x95, 0xfb, 0x1e, 0xdd, 0xf8, 0x79, 0xc4, 0xaf, 0x77, 0x6f,
	0xdf, 0xef, 0xc7, 0x84, 0xff, 0xd5, 0xf8, 0xd9, 0xf1, 0x3e, 0x0a, 0xdf, 0xef, 0xaa, 0x2f, 0xc8,
	0xbb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xd7, 0x81, 0xcd, 0x92, 0x04, 0x00, 0x00,
}
