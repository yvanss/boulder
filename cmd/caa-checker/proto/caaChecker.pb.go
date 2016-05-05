// Code generated by protoc-gen-go.
// source: caaChecker.proto
// DO NOT EDIT!

/*
Package caaChecker is a generated protocol buffer package.

It is generated from these files:
	caaChecker.proto

It has these top-level messages:
	Check
	Result
*/
package caaChecker

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
const _ = proto.ProtoPackageIsVersion1

type Check struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	IssuerDomain     *string `protobuf:"bytes,2,opt,name=issuerDomain" json:"issuerDomain,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Check) Reset()                    { *m = Check{} }
func (m *Check) String() string            { return proto.CompactTextString(m) }
func (*Check) ProtoMessage()               {}
func (*Check) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Check) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Check) GetIssuerDomain() string {
	if m != nil && m.IssuerDomain != nil {
		return *m.IssuerDomain
	}
	return ""
}

type Result struct {
	Present          *bool  `protobuf:"varint,1,opt,name=present" json:"present,omitempty"`
	Valid            *bool  `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetPresent() bool {
	if m != nil && m.Present != nil {
		return *m.Present
	}
	return false
}

func (m *Result) GetValid() bool {
	if m != nil && m.Valid != nil {
		return *m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*Check)(nil), "Check")
	proto.RegisterType((*Result)(nil), "Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for CAAChecker service

type CAACheckerClient interface {
	ValidForIssuance(ctx context.Context, in *Check, opts ...grpc.CallOption) (*Result, error)
}

type cAACheckerClient struct {
	cc *grpc.ClientConn
}

func NewCAACheckerClient(cc *grpc.ClientConn) CAACheckerClient {
	return &cAACheckerClient{cc}
}

func (c *cAACheckerClient) ValidForIssuance(ctx context.Context, in *Check, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/CAAChecker/ValidForIssuance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CAAChecker service

type CAACheckerServer interface {
	ValidForIssuance(context.Context, *Check) (*Result, error)
}

func RegisterCAACheckerServer(s *grpc.Server, srv CAACheckerServer) {
	s.RegisterService(&_CAAChecker_serviceDesc, srv)
}

func _CAAChecker_ValidForIssuance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Check)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAACheckerServer).ValidForIssuance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CAAChecker/ValidForIssuance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAACheckerServer).ValidForIssuance(ctx, req.(*Check))
	}
	return interceptor(ctx, in, info, handler)
}

var _CAAChecker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CAAChecker",
	HandlerType: (*CAACheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidForIssuance",
			Handler:    _CAAChecker_ValidForIssuance_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4e, 0x4c, 0x74,
	0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe6, 0x62,
	0x05, 0x0b, 0x08, 0xf1, 0x70, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x0a, 0x89, 0x70, 0xf1, 0x64, 0x16, 0x17, 0x97, 0xa6, 0x16, 0xb9, 0xe4, 0xe7, 0x26, 0x66, 0xe6,
	0x49, 0x30, 0x81, 0x44, 0x95, 0x34, 0xb8, 0xd8, 0x82, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x84, 0xf8,
	0xb9, 0xd8, 0x0b, 0x8a, 0x52, 0x8b, 0x53, 0xf3, 0x4a, 0xc0, 0x1a, 0x38, 0x84, 0x78, 0xb9, 0x58,
	0xcb, 0x12, 0x73, 0x32, 0x53, 0xc0, 0x2a, 0x39, 0x8c, 0x8c, 0xb9, 0xb8, 0x9c, 0x1d, 0x1d, 0xa1,
	0x56, 0x09, 0xa9, 0x72, 0x09, 0x84, 0x81, 0x24, 0xdd, 0xf2, 0x8b, 0x3c, 0x81, 0xa6, 0x26, 0xe6,
	0x25, 0xa7, 0x0a, 0xb1, 0xe9, 0x81, 0x65, 0xa5, 0xd8, 0xf5, 0x20, 0x46, 0x2a, 0x31, 0x00, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa4, 0xed, 0x44, 0x02, 0x9e, 0x00, 0x00, 0x00,
}
