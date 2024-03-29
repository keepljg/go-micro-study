// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fileinfo.proto

package go_micro_srv_fileinfo

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

type EntityId struct {
	EntityId             string   `protobuf:"bytes,1,opt,name=entityId,proto3" json:"entityId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityId) Reset()         { *m = EntityId{} }
func (m *EntityId) String() string { return proto.CompactTextString(m) }
func (*EntityId) ProtoMessage()    {}
func (*EntityId) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44c590a9180fd62, []int{0}
}

func (m *EntityId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityId.Unmarshal(m, b)
}
func (m *EntityId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityId.Marshal(b, m, deterministic)
}
func (m *EntityId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityId.Merge(m, src)
}
func (m *EntityId) XXX_Size() int {
	return xxx_messageInfo_EntityId.Size(m)
}
func (m *EntityId) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityId.DiscardUnknown(m)
}

var xxx_messageInfo_EntityId proto.InternalMessageInfo

func (m *EntityId) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

type Request struct {
	EntityIds            []*EntityId `protobuf:"bytes,1,rep,name=entityIds,proto3" json:"entityIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44c590a9180fd62, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetEntityIds() []*EntityId {
	if m != nil {
		return m.EntityIds
	}
	return nil
}

type AppInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IconUrl              string   `protobuf:"bytes,2,opt,name=iconUrl,proto3" json:"iconUrl,omitempty"`
	Screens              []string `protobuf:"bytes,3,rep,name=screens,proto3" json:"screens,omitempty"`
	DownloadUrl          string   `protobuf:"bytes,4,opt,name=downloadUrl,proto3" json:"downloadUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppInfo) Reset()         { *m = AppInfo{} }
func (m *AppInfo) String() string { return proto.CompactTextString(m) }
func (*AppInfo) ProtoMessage()    {}
func (*AppInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44c590a9180fd62, []int{2}
}

func (m *AppInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppInfo.Unmarshal(m, b)
}
func (m *AppInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppInfo.Marshal(b, m, deterministic)
}
func (m *AppInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppInfo.Merge(m, src)
}
func (m *AppInfo) XXX_Size() int {
	return xxx_messageInfo_AppInfo.Size(m)
}
func (m *AppInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AppInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AppInfo proto.InternalMessageInfo

func (m *AppInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AppInfo) GetIconUrl() string {
	if m != nil {
		return m.IconUrl
	}
	return ""
}

func (m *AppInfo) GetScreens() []string {
	if m != nil {
		return m.Screens
	}
	return nil
}

func (m *AppInfo) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

type Response struct {
	AppInfos             []*AppInfo `protobuf:"bytes,1,rep,name=appInfos,proto3" json:"appInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a44c590a9180fd62, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetAppInfos() []*AppInfo {
	if m != nil {
		return m.AppInfos
	}
	return nil
}

func init() {
	proto.RegisterType((*EntityId)(nil), "go.tutu.srv.fileinfo.EntityId")
	proto.RegisterType((*Request)(nil), "go.tutu.srv.fileinfo.Request")
	proto.RegisterType((*AppInfo)(nil), "go.tutu.srv.fileinfo.AppInfo")
	proto.RegisterType((*Response)(nil), "go.tutu.srv.fileinfo.Response")
}

func init() { proto.RegisterFile("fileinfo.proto", fileDescriptor_a44c590a9180fd62) }

var fileDescriptor_a44c590a9180fd62 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x85, 0xeb, 0x3a, 0xd4, 0xca, 0x05, 0x3a, 0x88, 0x0c, 0x22, 0xd0, 0x60, 0x34, 0x94, 0x4c,
	0x1a, 0xd2, 0xa9, 0xd0, 0xa5, 0x85, 0x34, 0x64, 0x15, 0x94, 0xce, 0xae, 0x7d, 0x2e, 0x02, 0x57,
	0x72, 0x2d, 0xb9, 0x25, 0xff, 0xbe, 0x58, 0x91, 0xdc, 0x0c, 0xf1, 0x76, 0x77, 0xef, 0xde, 0xd3,
	0xc7, 0x09, 0x6e, 0x6b, 0xd5, 0xa0, 0xd2, 0xb5, 0x11, 0x6d, 0x67, 0x9c, 0xa1, 0xcb, 0x4f, 0x23,
	0x5c, 0xef, 0x7a, 0x61, 0xbb, 0x1f, 0x11, 0x35, 0x7e, 0x0f, 0x64, 0xa7, 0x9d, 0x72, 0xc7, 0x43,
	0x45, 0x57, 0x40, 0x30, 0xd4, 0x2c, 0xc9, 0x93, 0xcd, 0x5c, 0x8e, 0x3d, 0xdf, 0x43, 0x26, 0xf1,
	0xbb, 0x47, 0xeb, 0xe8, 0x13, 0xcc, 0xe3, 0xd8, 0xb2, 0x24, 0x4f, 0x37, 0x8b, 0xed, 0x5a, 0x5c,
	0x0a, 0x17, 0x31, 0x59, 0xfe, 0x1b, 0xb8, 0x85, 0xec, 0xb9, 0x6d, 0x0f, 0xba, 0x36, 0x94, 0xc2,
	0x4c, 0x17, 0x5f, 0x18, 0xde, 0xf2, 0x35, 0x65, 0x90, 0xa9, 0xd2, 0xe8, 0xb7, 0xae, 0x61, 0xd7,
	0x7e, 0x1c, 0xdb, 0x41, 0xb1, 0x65, 0x87, 0xa8, 0x2d, 0x4b, 0xf3, 0x74, 0x50, 0x42, 0x4b, 0x73,
	0x58, 0x54, 0xe6, 0x57, 0x37, 0xa6, 0xa8, 0x06, 0xdf, 0xcc, 0xfb, 0xce, 0x47, 0x7c, 0x07, 0x44,
	0xa2, 0x6d, 0x8d, 0xb6, 0x48, 0x1f, 0x81, 0x14, 0x27, 0x80, 0x48, 0x7f, 0x77, 0x99, 0x3e, 0x60,
	0xca, 0x71, 0x7d, 0x5b, 0x02, 0x79, 0x55, 0x0d, 0x7a, 0xf8, 0x77, 0x58, 0xee, 0xd1, 0x85, 0x9d,
	0x97, 0xe3, 0x78, 0xc4, 0x89, 0xb0, 0x70, 0xbc, 0xd5, 0x7a, 0x4a, 0x3e, 0xd1, 0xf1, 0xab, 0x8f,
	0x1b, 0xff, 0x5d, 0x0f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xe1, 0x70, 0x76, 0xc0, 0x01,
	0x00, 0x00,
}
