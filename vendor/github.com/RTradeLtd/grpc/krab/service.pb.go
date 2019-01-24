// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package krab

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_41c0396827bd23f6, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type KeyGet struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyGet) Reset()         { *m = KeyGet{} }
func (m *KeyGet) String() string { return proto.CompactTextString(m) }
func (*KeyGet) ProtoMessage()    {}
func (*KeyGet) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_41c0396827bd23f6, []int{1}
}
func (m *KeyGet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyGet.Unmarshal(m, b)
}
func (m *KeyGet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyGet.Marshal(b, m, deterministic)
}
func (dst *KeyGet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyGet.Merge(dst, src)
}
func (m *KeyGet) XXX_Size() int {
	return xxx_messageInfo_KeyGet.Size(m)
}
func (m *KeyGet) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyGet.DiscardUnknown(m)
}

var xxx_messageInfo_KeyGet proto.InternalMessageInfo

func (m *KeyGet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type KeyPut struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PrivateKey           []byte   `protobuf:"bytes,2,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyPut) Reset()         { *m = KeyPut{} }
func (m *KeyPut) String() string { return proto.CompactTextString(m) }
func (*KeyPut) ProtoMessage()    {}
func (*KeyPut) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_41c0396827bd23f6, []int{2}
}
func (m *KeyPut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyPut.Unmarshal(m, b)
}
func (m *KeyPut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyPut.Marshal(b, m, deterministic)
}
func (dst *KeyPut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPut.Merge(dst, src)
}
func (m *KeyPut) XXX_Size() int {
	return xxx_messageInfo_KeyPut.Size(m)
}
func (m *KeyPut) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPut.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPut proto.InternalMessageInfo

func (m *KeyPut) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KeyPut) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

type KeyDelete struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyDelete) Reset()         { *m = KeyDelete{} }
func (m *KeyDelete) String() string { return proto.CompactTextString(m) }
func (*KeyDelete) ProtoMessage()    {}
func (*KeyDelete) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_41c0396827bd23f6, []int{3}
}
func (m *KeyDelete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyDelete.Unmarshal(m, b)
}
func (m *KeyDelete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyDelete.Marshal(b, m, deterministic)
}
func (dst *KeyDelete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyDelete.Merge(dst, src)
}
func (m *KeyDelete) XXX_Size() int {
	return xxx_messageInfo_KeyDelete.Size(m)
}
func (m *KeyDelete) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyDelete.DiscardUnknown(m)
}

var xxx_messageInfo_KeyDelete proto.InternalMessageInfo

func (m *KeyDelete) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PrivateKey           []byte   `protobuf:"bytes,2,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_41c0396827bd23f6, []int{4}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Response) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "krab.Empty")
	proto.RegisterType((*KeyGet)(nil), "krab.KeyGet")
	proto.RegisterType((*KeyPut)(nil), "krab.KeyPut")
	proto.RegisterType((*KeyDelete)(nil), "krab.KeyDelete")
	proto.RegisterType((*Response)(nil), "krab.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	GetPrivateKey(ctx context.Context, in *KeyGet, opts ...grpc.CallOption) (*Response, error)
	PutPrivateKey(ctx context.Context, in *KeyPut, opts ...grpc.CallOption) (*Response, error)
	DeletePrivateKey(ctx context.Context, in *KeyDelete, opts ...grpc.CallOption) (*Response, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetPrivateKey(ctx context.Context, in *KeyGet, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/krab.Service/GetPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) PutPrivateKey(ctx context.Context, in *KeyPut, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/krab.Service/PutPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeletePrivateKey(ctx context.Context, in *KeyDelete, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/krab.Service/DeletePrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	GetPrivateKey(context.Context, *KeyGet) (*Response, error)
	PutPrivateKey(context.Context, *KeyPut) (*Response, error)
	DeletePrivateKey(context.Context, *KeyDelete) (*Response, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_GetPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/krab.Service/GetPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetPrivateKey(ctx, req.(*KeyGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_PutPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyPut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).PutPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/krab.Service/PutPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).PutPrivateKey(ctx, req.(*KeyPut))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeletePrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeletePrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/krab.Service/DeletePrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeletePrivateKey(ctx, req.(*KeyDelete))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "krab.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPrivateKey",
			Handler:    _Service_GetPrivateKey_Handler,
		},
		{
			MethodName: "PutPrivateKey",
			Handler:    _Service_PutPrivateKey_Handler,
		},
		{
			MethodName: "DeletePrivateKey",
			Handler:    _Service_DeletePrivateKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_41c0396827bd23f6) }

var fileDescriptor_service_41c0396827bd23f6 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x2e, 0x4a, 0x4c, 0x52,
	0x62, 0xe7, 0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x54, 0x92, 0xe1, 0x62, 0xf3, 0x4e, 0xad, 0x74,
	0x4f, 0x2d, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x95, 0x6c, 0xc0, 0xb2, 0x01, 0xa5, 0x58, 0x65, 0x85, 0xe4, 0xb8, 0xb8, 0x0a,
	0x8a, 0x32, 0xcb, 0x12, 0x4b, 0x52, 0xbd, 0x53, 0x2b, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82,
	0x90, 0x44, 0x94, 0xe4, 0xb9, 0x38, 0xbd, 0x53, 0x2b, 0x5d, 0x52, 0x73, 0x52, 0x4b, 0x52, 0xb1,
	0x1a, 0xef, 0xc4, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5,
	0x56, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x0c, 0x55, 0x01, 0xe5, 0x11, 0xb2, 0xc4, 0x68, 0x11, 0x23,
	0x17, 0x7b, 0x30, 0xc4, 0x87, 0x42, 0xfa, 0x5c, 0xbc, 0xee, 0xa9, 0x25, 0x01, 0x70, 0x49, 0x21,
	0x1e, 0x3d, 0x90, 0x6f, 0xf5, 0x20, 0x3e, 0x94, 0xe2, 0x83, 0xf0, 0x60, 0x56, 0x2a, 0x31, 0x80,
	0x34, 0x04, 0x94, 0x62, 0xd7, 0x10, 0x50, 0x8a, 0x4d, 0x83, 0x29, 0x97, 0x00, 0xc4, 0x3f, 0x48,
	0x7a, 0xf8, 0xe1, 0x7a, 0x20, 0x52, 0x98, 0xda, 0x92, 0xd8, 0xc0, 0x61, 0x6f, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x09, 0xfe, 0xae, 0x76, 0x8c, 0x01, 0x00, 0x00,
}
