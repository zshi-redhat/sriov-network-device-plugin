// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/user_interest_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v0/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Request message for [UserInterestService.GetUserInterest][google.ads.googleads.v0.services.UserInterestService.GetUserInterest].
type GetUserInterestRequest struct {
	// Resource name of the UserInterest to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInterestRequest) Reset()         { *m = GetUserInterestRequest{} }
func (m *GetUserInterestRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInterestRequest) ProtoMessage()    {}
func (*GetUserInterestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_interest_service_6325edec6159d584, []int{0}
}
func (m *GetUserInterestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInterestRequest.Unmarshal(m, b)
}
func (m *GetUserInterestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInterestRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserInterestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInterestRequest.Merge(dst, src)
}
func (m *GetUserInterestRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInterestRequest.Size(m)
}
func (m *GetUserInterestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInterestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInterestRequest proto.InternalMessageInfo

func (m *GetUserInterestRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserInterestRequest)(nil), "google.ads.googleads.v0.services.GetUserInterestRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserInterestServiceClient is the client API for UserInterestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInterestServiceClient interface {
	// Returns the requested user interest in full detail
	GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error)
}

type userInterestServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserInterestServiceClient(cc *grpc.ClientConn) UserInterestServiceClient {
	return &userInterestServiceClient{cc}
}

func (c *userInterestServiceClient) GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error) {
	out := new(resources.UserInterest)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.UserInterestService/GetUserInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInterestServiceServer is the server API for UserInterestService service.
type UserInterestServiceServer interface {
	// Returns the requested user interest in full detail
	GetUserInterest(context.Context, *GetUserInterestRequest) (*resources.UserInterest, error)
}

func RegisterUserInterestServiceServer(s *grpc.Server, srv UserInterestServiceServer) {
	s.RegisterService(&_UserInterestService_serviceDesc, srv)
}

func _UserInterestService_GetUserInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.UserInterestService/GetUserInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, req.(*GetUserInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInterestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.UserInterestService",
	HandlerType: (*UserInterestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInterest",
			Handler:    _UserInterestService_GetUserInterest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/user_interest_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/user_interest_service.proto", fileDescriptor_user_interest_service_6325edec6159d584)
}

var fileDescriptor_user_interest_service_6325edec6159d584 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0xc7, 0x49, 0x05, 0xc1, 0x43, 0x11, 0x22, 0x48, 0x29, 0x0e, 0xa5, 0x3a, 0x48, 0x87, 0xbb,
	0xa8, 0x88, 0x82, 0x76, 0x48, 0x97, 0xea, 0x22, 0xa5, 0x62, 0x07, 0x09, 0x94, 0xb3, 0xf9, 0x38,
	0x02, 0xcd, 0x5d, 0xbd, 0xef, 0xd2, 0x45, 0x5c, 0x7c, 0x05, 0xdf, 0xc0, 0xd1, 0xdd, 0x97, 0x10,
	0x9c, 0x5c, 0x7c, 0x00, 0x1f, 0x44, 0xd2, 0xcb, 0x95, 0xaa, 0x0d, 0xdd, 0xfe, 0x5c, 0xbe, 0xdf,
	0xff, 0xbe, 0xff, 0x3f, 0x47, 0xce, 0x85, 0x52, 0x62, 0x04, 0x8c, 0xc7, 0xc8, 0xac, 0xcc, 0xd5,
	0x24, 0x60, 0x08, 0x7a, 0x92, 0x0c, 0x01, 0x59, 0x86, 0xa0, 0x07, 0x89, 0x34, 0xa0, 0x01, 0xcd,
	0xa0, 0x38, 0xa6, 0x63, 0xad, 0x8c, 0xf2, 0xeb, 0x16, 0xa1, 0x3c, 0x46, 0x3a, 0xa3, 0xe9, 0x24,
	0xa0, 0x8e, 0xae, 0x1d, 0x97, 0xf9, 0x6b, 0x40, 0x95, 0xe9, 0x7f, 0x17, 0x58, 0xe3, 0xda, 0x8e,
	0xc3, 0xc6, 0x09, 0xe3, 0x52, 0x2a, 0xc3, 0x4d, 0xa2, 0x24, 0xda, 0xaf, 0x8d, 0x16, 0xd9, 0xee,
	0x80, 0xb9, 0x41, 0xd0, 0x97, 0x05, 0xd6, 0x83, 0xfb, 0x0c, 0xd0, 0xf8, 0xbb, 0x64, 0xc3, 0x19,
	0x0f, 0x24, 0x4f, 0xa1, 0xea, 0xd5, 0xbd, 0xfd, 0xb5, 0xde, 0xba, 0x3b, 0xbc, 0xe2, 0x29, 0x1c,
	0x7e, 0x78, 0x64, 0x6b, 0x1e, 0xbe, 0xb6, 0xcb, 0xfa, 0x6f, 0x1e, 0xd9, 0xfc, 0xe3, 0xeb, 0x9f,
	0xd2, 0x65, 0x11, 0xe9, 0xe2, 0x55, 0x6a, 0xac, 0x94, 0x9c, 0x45, 0xa7, 0xf3, 0x5c, 0xe3, 0xe4,
	0xe9, 0xf3, 0xfb, 0xb9, 0x72, 0xe0, 0xb3, 0xbc, 0x9e, 0x87, 0x5f, 0x31, 0x5a, 0xc3, 0x0c, 0x8d,
	0x4a, 0x41, 0x23, 0x6b, 0x4e, 0xfb, 0x72, 0x10, 0xb2, 0xe6, 0x63, 0xfb, 0xcb, 0x23, 0x7b, 0x43,
	0x95, 0x2e, 0xdd, 0xb4, 0x5d, 0x5d, 0x90, 0xba, 0x9b, 0x37, 0xda, 0xf5, 0x6e, 0x2f, 0x0a, 0x5a,
	0xa8, 0x11, 0x97, 0x82, 0x2a, 0x2d, 0x98, 0x00, 0x39, 0xed, 0xdb, 0xfd, 0xb8, 0x71, 0x82, 0xe5,
	0xef, 0xe4, 0xcc, 0x89, 0x97, 0xca, 0x4a, 0x27, 0x0c, 0x5f, 0x2b, 0xf5, 0x8e, 0x35, 0x0c, 0x63,
	0xa4, 0x56, 0xe6, 0xaa, 0x1f, 0xd0, 0xe2, 0x62, 0x7c, 0x77, 0x23, 0x51, 0x18, 0x63, 0x34, 0x1b,
	0x89, 0xfa, 0x41, 0xe4, 0x46, 0xee, 0x56, 0xa7, 0x0b, 0x1c, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xb7, 0x66, 0x2b, 0x2d, 0xa7, 0x02, 0x00, 0x00,
}
