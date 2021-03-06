// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/crawler/crawler.proto

package crawler

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetVkUsersPhotosHandlerRequest struct {
	Users                []uint64 `protobuf:"varint,1,rep,packed,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVkUsersPhotosHandlerRequest) Reset()         { *m = GetVkUsersPhotosHandlerRequest{} }
func (m *GetVkUsersPhotosHandlerRequest) String() string { return proto.CompactTextString(m) }
func (*GetVkUsersPhotosHandlerRequest) ProtoMessage()    {}
func (*GetVkUsersPhotosHandlerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_440e940d8a62a8c7, []int{0}
}

func (m *GetVkUsersPhotosHandlerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVkUsersPhotosHandlerRequest.Unmarshal(m, b)
}
func (m *GetVkUsersPhotosHandlerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVkUsersPhotosHandlerRequest.Marshal(b, m, deterministic)
}
func (m *GetVkUsersPhotosHandlerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVkUsersPhotosHandlerRequest.Merge(m, src)
}
func (m *GetVkUsersPhotosHandlerRequest) XXX_Size() int {
	return xxx_messageInfo_GetVkUsersPhotosHandlerRequest.Size(m)
}
func (m *GetVkUsersPhotosHandlerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVkUsersPhotosHandlerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVkUsersPhotosHandlerRequest proto.InternalMessageInfo

func (m *GetVkUsersPhotosHandlerRequest) GetUsers() []uint64 {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetVkUsersPhotosHandlerResponse struct {
	Photos               []*GetVkUsersPhotosHandlerResponse_UserPhotos `protobuf:"bytes,1,rep,name=photos,proto3" json:"photos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                      `json:"-"`
	XXX_unrecognized     []byte                                        `json:"-"`
	XXX_sizecache        int32                                         `json:"-"`
}

func (m *GetVkUsersPhotosHandlerResponse) Reset()         { *m = GetVkUsersPhotosHandlerResponse{} }
func (m *GetVkUsersPhotosHandlerResponse) String() string { return proto.CompactTextString(m) }
func (*GetVkUsersPhotosHandlerResponse) ProtoMessage()    {}
func (*GetVkUsersPhotosHandlerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_440e940d8a62a8c7, []int{1}
}

func (m *GetVkUsersPhotosHandlerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVkUsersPhotosHandlerResponse.Unmarshal(m, b)
}
func (m *GetVkUsersPhotosHandlerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVkUsersPhotosHandlerResponse.Marshal(b, m, deterministic)
}
func (m *GetVkUsersPhotosHandlerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVkUsersPhotosHandlerResponse.Merge(m, src)
}
func (m *GetVkUsersPhotosHandlerResponse) XXX_Size() int {
	return xxx_messageInfo_GetVkUsersPhotosHandlerResponse.Size(m)
}
func (m *GetVkUsersPhotosHandlerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVkUsersPhotosHandlerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetVkUsersPhotosHandlerResponse proto.InternalMessageInfo

func (m *GetVkUsersPhotosHandlerResponse) GetPhotos() []*GetVkUsersPhotosHandlerResponse_UserPhotos {
	if m != nil {
		return m.Photos
	}
	return nil
}

type GetVkUsersPhotosHandlerResponse_UserPhotos struct {
	User                 uint64   `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Photos               [][]byte `protobuf:"bytes,2,rep,name=photos,proto3" json:"photos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetVkUsersPhotosHandlerResponse_UserPhotos) Reset() {
	*m = GetVkUsersPhotosHandlerResponse_UserPhotos{}
}
func (m *GetVkUsersPhotosHandlerResponse_UserPhotos) String() string {
	return proto.CompactTextString(m)
}
func (*GetVkUsersPhotosHandlerResponse_UserPhotos) ProtoMessage() {}
func (*GetVkUsersPhotosHandlerResponse_UserPhotos) Descriptor() ([]byte, []int) {
	return fileDescriptor_440e940d8a62a8c7, []int{1, 0}
}

func (m *GetVkUsersPhotosHandlerResponse_UserPhotos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVkUsersPhotosHandlerResponse_UserPhotos.Unmarshal(m, b)
}
func (m *GetVkUsersPhotosHandlerResponse_UserPhotos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVkUsersPhotosHandlerResponse_UserPhotos.Marshal(b, m, deterministic)
}
func (m *GetVkUsersPhotosHandlerResponse_UserPhotos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVkUsersPhotosHandlerResponse_UserPhotos.Merge(m, src)
}
func (m *GetVkUsersPhotosHandlerResponse_UserPhotos) XXX_Size() int {
	return xxx_messageInfo_GetVkUsersPhotosHandlerResponse_UserPhotos.Size(m)
}
func (m *GetVkUsersPhotosHandlerResponse_UserPhotos) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVkUsersPhotosHandlerResponse_UserPhotos.DiscardUnknown(m)
}

var xxx_messageInfo_GetVkUsersPhotosHandlerResponse_UserPhotos proto.InternalMessageInfo

func (m *GetVkUsersPhotosHandlerResponse_UserPhotos) GetUser() uint64 {
	if m != nil {
		return m.User
	}
	return 0
}

func (m *GetVkUsersPhotosHandlerResponse_UserPhotos) GetPhotos() [][]byte {
	if m != nil {
		return m.Photos
	}
	return nil
}

func init() {
	proto.RegisterType((*GetVkUsersPhotosHandlerRequest)(nil), "crawler.GetVkUsersPhotosHandlerRequest")
	proto.RegisterType((*GetVkUsersPhotosHandlerResponse)(nil), "crawler.GetVkUsersPhotosHandlerResponse")
	proto.RegisterType((*GetVkUsersPhotosHandlerResponse_UserPhotos)(nil), "crawler.GetVkUsersPhotosHandlerResponse.UserPhotos")
}

func init() {
	proto.RegisterFile("proto/crawler/crawler.proto", fileDescriptor_440e940d8a62a8c7)
}

var fileDescriptor_440e940d8a62a8c7 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2e, 0x4a, 0x2c, 0xcf, 0x49, 0x2d, 0x82, 0xd1, 0x7a, 0x60, 0x51, 0x21, 0x76,
	0x28, 0x57, 0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55, 0x3f, 0xb1, 0x20, 0x53, 0x3f, 0x31,
	0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x18, 0xa2, 0x4c, 0xc9, 0x8c, 0x4b, 0xce,
	0x3d, 0xb5, 0x24, 0x2c, 0x3b, 0xb4, 0x38, 0xb5, 0xa8, 0x38, 0x20, 0x23, 0xbf, 0x24, 0xbf, 0xd8,
	0x23, 0x31, 0x2f, 0x25, 0x27, 0xb5, 0x28, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x84,
	0x8b, 0xb5, 0x14, 0x24, 0x29, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x12, 0x04, 0xe1, 0x28, 0xad, 0x60,
	0xe4, 0x92, 0xc7, 0xa9, 0xb1, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x9b, 0x8b, 0xad, 0x00,
	0x2c, 0x01, 0xd6, 0xca, 0x6d, 0x64, 0xac, 0x07, 0x73, 0x22, 0x01, 0x9d, 0x7a, 0x20, 0x29, 0x88,
	0x4c, 0x10, 0xd4, 0x08, 0x29, 0x0b, 0x2e, 0x2e, 0x84, 0xa8, 0x90, 0x10, 0x17, 0x0b, 0xc8, 0x1d,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x60, 0xb6, 0x90, 0x18, 0xdc, 0x3a, 0x26, 0x05, 0x66,
	0x0d, 0x1e, 0x98, 0x4e, 0xa3, 0x85, 0x8c, 0x5c, 0xec, 0xce, 0x10, 0x8b, 0x85, 0xa6, 0x32, 0x72,
	0x89, 0xe3, 0xb0, 0x5c, 0x48, 0x9d, 0xb0, 0xf3, 0xc0, 0x21, 0x22, 0xa5, 0x41, 0xac, 0x3f, 0x94,
	0xd4, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0xa4, 0xa8, 0x24, 0x03, 0x0e, 0xfd, 0x32, 0x43, 0x78, 0x5c,
	0x95, 0x65, 0xeb, 0x57, 0x83, 0x1c, 0x1d, 0x9f, 0x99, 0x52, 0x6b, 0xc5, 0xa8, 0xe5, 0x64, 0x17,
	0x65, 0x93, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef, 0x9a, 0x93, 0x92,
	0x5a, 0x94, 0x53, 0xa9, 0xeb, 0xe8, 0xa9, 0x9f, 0x9f, 0x54, 0x9c, 0x5a, 0x54, 0x86, 0x88, 0x5d,
	0x7d, 0x94, 0x38, 0xb7, 0x86, 0xd2, 0x49, 0x6c, 0x60, 0x61, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xfe, 0x72, 0x5d, 0x3a, 0x13, 0x02, 0x00, 0x00,
}
