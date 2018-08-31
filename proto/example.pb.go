// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example.proto

package example

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

type PhoneReq struct {
	Imei                 string   `protobuf:"bytes,1,opt,name=imei,proto3" json:"imei,omitempty"`
	Modelname            string   `protobuf:"bytes,2,opt,name=modelname,proto3" json:"modelname,omitempty"`
	Unlockcode           string   `protobuf:"bytes,3,opt,name=unlockcode,proto3" json:"unlockcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhoneReq) Reset()         { *m = PhoneReq{} }
func (m *PhoneReq) String() string { return proto.CompactTextString(m) }
func (*PhoneReq) ProtoMessage()    {}
func (*PhoneReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_2a84d8b71000ab4d, []int{0}
}
func (m *PhoneReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneReq.Unmarshal(m, b)
}
func (m *PhoneReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneReq.Marshal(b, m, deterministic)
}
func (dst *PhoneReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneReq.Merge(dst, src)
}
func (m *PhoneReq) XXX_Size() int {
	return xxx_messageInfo_PhoneReq.Size(m)
}
func (m *PhoneReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneReq.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneReq proto.InternalMessageInfo

func (m *PhoneReq) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *PhoneReq) GetModelname() string {
	if m != nil {
		return m.Modelname
	}
	return ""
}

func (m *PhoneReq) GetUnlockcode() string {
	if m != nil {
		return m.Unlockcode
	}
	return ""
}

type PhoneResp struct {
	Imei                 string   `protobuf:"bytes,1,opt,name=imei,proto3" json:"imei,omitempty"`
	Modelname            string   `protobuf:"bytes,2,opt,name=modelname,proto3" json:"modelname,omitempty"`
	Unlockcode           string   `protobuf:"bytes,3,opt,name=unlockcode,proto3" json:"unlockcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhoneResp) Reset()         { *m = PhoneResp{} }
func (m *PhoneResp) String() string { return proto.CompactTextString(m) }
func (*PhoneResp) ProtoMessage()    {}
func (*PhoneResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_2a84d8b71000ab4d, []int{1}
}
func (m *PhoneResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneResp.Unmarshal(m, b)
}
func (m *PhoneResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneResp.Marshal(b, m, deterministic)
}
func (dst *PhoneResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneResp.Merge(dst, src)
}
func (m *PhoneResp) XXX_Size() int {
	return xxx_messageInfo_PhoneResp.Size(m)
}
func (m *PhoneResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneResp.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneResp proto.InternalMessageInfo

func (m *PhoneResp) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *PhoneResp) GetModelname() string {
	if m != nil {
		return m.Modelname
	}
	return ""
}

func (m *PhoneResp) GetUnlockcode() string {
	if m != nil {
		return m.Unlockcode
	}
	return ""
}

func init() {
	proto.RegisterType((*PhoneReq)(nil), "example.PhoneReq")
	proto.RegisterType((*PhoneResp)(nil), "example.PhoneResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PhoneTrackerClient is the client API for PhoneTracker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PhoneTrackerClient interface {
	AddPhone(ctx context.Context, in *PhoneReq, opts ...grpc.CallOption) (*PhoneResp, error)
}

type phoneTrackerClient struct {
	cc *grpc.ClientConn
}

func NewPhoneTrackerClient(cc *grpc.ClientConn) PhoneTrackerClient {
	return &phoneTrackerClient{cc}
}

func (c *phoneTrackerClient) AddPhone(ctx context.Context, in *PhoneReq, opts ...grpc.CallOption) (*PhoneResp, error) {
	out := new(PhoneResp)
	err := c.cc.Invoke(ctx, "/example.PhoneTracker/AddPhone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhoneTrackerServer is the server API for PhoneTracker service.
type PhoneTrackerServer interface {
	AddPhone(context.Context, *PhoneReq) (*PhoneResp, error)
}

func RegisterPhoneTrackerServer(s *grpc.Server, srv PhoneTrackerServer) {
	s.RegisterService(&_PhoneTracker_serviceDesc, srv)
}

func _PhoneTracker_AddPhone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhoneTrackerServer).AddPhone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.PhoneTracker/AddPhone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhoneTrackerServer).AddPhone(ctx, req.(*PhoneReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhoneTracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.PhoneTracker",
	HandlerType: (*PhoneTrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPhone",
			Handler:    _PhoneTracker_AddPhone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/example.proto",
}

func init() { proto.RegisterFile("proto/example.proto", fileDescriptor_example_2a84d8b71000ab4d) }

var fileDescriptor_example_2a84d8b71000ab4d = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x03, 0xf3, 0x84, 0xd8, 0xa1, 0x5c,
	0xa5, 0x18, 0x2e, 0x8e, 0x80, 0x8c, 0xfc, 0xbc, 0xd4, 0xa0, 0xd4, 0x42, 0x21, 0x21, 0x2e, 0x96,
	0xcc, 0xdc, 0xd4, 0x4c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x86, 0x8b,
	0x33, 0x37, 0x3f, 0x25, 0x35, 0x27, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x09, 0x2c, 0x81, 0x10, 0x10,
	0x92, 0xe3, 0xe2, 0x2a, 0xcd, 0xcb, 0xc9, 0x4f, 0xce, 0x4e, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x06,
	0x4b, 0x23, 0x89, 0x28, 0xc5, 0x72, 0x71, 0x42, 0x4d, 0x2f, 0x2e, 0xa0, 0xbe, 0xf1, 0x46, 0xce,
	0x5c, 0x3c, 0x60, 0xe3, 0x43, 0x8a, 0x12, 0x93, 0xb3, 0x53, 0x8b, 0x84, 0x8c, 0xb9, 0x38, 0x1c,
	0x53, 0x52, 0xc0, 0x42, 0x42, 0x82, 0x7a, 0x30, 0x1f, 0xc3, 0xfc, 0x27, 0x25, 0x84, 0x2e, 0x54,
	0x5c, 0xa0, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x11, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d,
	0x29, 0x56, 0x5e, 0x28, 0x01, 0x00, 0x00,
}
