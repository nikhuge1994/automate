// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/rules.proto

package v2beta

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2beta/rules.proto", fileDescriptor_471a411ff4d64939)
}

var fileDescriptor_471a411ff4d64939 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xbb, 0x6e, 0xd4, 0x4c,
	0x14, 0xc7, 0xe5, 0x8d, 0x92, 0xef, 0x63, 0x40, 0x4a, 0x72, 0x36, 0x08, 0xc7, 0x44, 0x42, 0xb8,
	0x81, 0x84, 0xac, 0x0d, 0xe1, 0x52, 0x6c, 0xc1, 0x2d, 0x5c, 0x1a, 0x0a, 0x94, 0x88, 0x86, 0x06,
	0x4d, 0x9c, 0xc1, 0x71, 0x64, 0x7b, 0x26, 0x9e, 0x31, 0x68, 0x15, 0xd1, 0xb8, 0xa0, 0x58, 0x3a,
	0x68, 0x28, 0x68, 0x78, 0x08, 0xbf, 0x06, 0x0d, 0x6f, 0x80, 0x90, 0x78, 0x00, 0x2a, 0x3a, 0x34,
	0x33, 0x7b, 0xf1, 0x6e, 0x36, 0x9b, 0x31, 0x45, 0xba, 0xb5, 0xfd, 0x3b, 0x67, 0xce, 0xef, 0x3f,
	0xa3, 0xd5, 0xa0, 0x3b, 0x01, 0x4d, 0x18, 0x4d, 0x49, 0x2a, 0xb8, 0x8f, 0x73, 0x41, 0x13, 0x2c,
	0x48, 0x2b, 0xc4, 0x82, 0xbc, 0xc5, 0x1d, 0x1f, 0xb3, 0xc8, 0x8f, 0x70, 0xe2, 0xbf, 0xd9, 0xd8,
	0x21, 0x02, 0xfb, 0x59, 0x1e, 0x13, 0xee, 0xb1, 0x8c, 0x0a, 0x0a, 0x2b, 0xc1, 0x1e, 0x79, 0xed,
	0xf5, 0x2b, 0x3c, 0xcc, 0x22, 0x2f, 0xc2, 0x89, 0xa7, 0x49, 0x67, 0x25, 0xa4, 0x34, 0x8c, 0x89,
	0x6a, 0x80, 0xd3, 0x94, 0x0a, 0x2c, 0x22, 0x9a, 0xf6, 0x6a, 0x9d, 0xbb, 0x75, 0xd6, 0x24, 0x07,
	0x39, 0xe1, 0xa2, 0xba, 0xb6, 0x73, 0xaf, 0x56, 0x3d, 0x67, 0x34, 0xe5, 0x64, 0xa4, 0xc1, 0xfd,
	0x89, 0x0d, 0x32, 0x16, 0xf8, 0xea, 0x7b, 0xd0, 0x0a, 0x49, 0xda, 0x62, 0x34, 0x8e, 0x82, 0xce,
	0x31, 0x0a, 0x75, 0x3a, 0xc8, 0x69, 0x8e, 0x74, 0xd8, 0xf8, 0x36, 0x8f, 0x66, 0xb7, 0xe4, 0x4c,
	0xf0, 0xbe, 0x81, 0xd0, 0x66, 0x46, 0xb0, 0x20, 0xf2, 0x19, 0xae, 0x79, 0xd3, 0xa2, 0xf5, 0x86,
	0xe4, 0x16, 0x39, 0x70, 0xd6, 0xcd, 0x61, 0xce, 0xdc, 0x2f, 0x56, 0x51, 0xda, 0x2b, 0xc8, 0xc1,
	0xb9, 0xd8, 0x6b, 0xb3, 0x8c, 0xee, 0x93, 0x40, 0xf0, 0xf6, 0x61, 0xef, 0xd7, 0xab, 0x68, 0xf7,
	0x5d, 0x51, 0xda, 0xff, 0xc3, 0x5c, 0xce, 0x76, 0xb1, 0x20, 0xdd, 0xd2, 0xbe, 0x88, 0x96, 0x23,
	0x9c, 0x4c, 0x46, 0xbb, 0xa5, 0x7d, 0x1e, 0x9a, 0x23, 0x9f, 0x75, 0x5d, 0xf1, 0xfd, 0xe7, 0xa7,
	0xc6, 0xaa, 0x7b, 0xa5, 0xba, 0x0d, 0x7d, 0xc4, 0xaf, 0x76, 0xd0, 0x7b, 0xa2, 0xf0, 0x99, 0xb6,
	0xb5, 0x06, 0x1f, 0x1a, 0x08, 0xbd, 0x50, 0x2d, 0x4c, 0x82, 0x18, 0x92, 0x06, 0x41, 0x54, 0x61,
	0xce, 0xdc, 0xaf, 0xa7, 0x17, 0x84, 0xe7, 0xac, 0x1b, 0x06, 0xe1, 0x1f, 0xca, 0x75, 0x07, 0x69,
	0xfc, 0xb6, 0xd0, 0x7f, 0x4f, 0x89, 0x50, 0x51, 0x5c, 0x9d, 0x6e, 0xd7, 0xc3, 0x64, 0x0e, 0xab,
	0x86, 0x24, 0x67, 0xee, 0x47, 0x93, 0x10, 0x66, 0x61, 0x26, 0x24, 0xe2, 0xe4, 0x04, 0x00, 0x16,
	0x46, 0x3e, 0x87, 0x44, 0x68, 0x7d, 0xa8, 0xa5, 0x0f, 0x7f, 0x2c, 0xd4, 0x7c, 0x16, 0x71, 0x35,
	0x25, 0x7f, 0x42, 0xb3, 0xe7, 0x1a, 0x83, 0x5b, 0xd3, 0xbd, 0x26, 0x94, 0xc8, 0x34, 0x6e, 0xff,
	0x43, 0x15, 0x67, 0x6e, 0x5e, 0x94, 0xf6, 0x12, 0x82, 0xb1, 0x60, 0xc6, 0x02, 0x69, 0xa2, 0xc5,
	0xd1, 0x40, 0xa6, 0x07, 0x71, 0x19, 0x2e, 0x4d, 0x0e, 0x62, 0x10, 0x00, 0x14, 0x0d, 0x84, 0x1e,
	0x91, 0x98, 0x98, 0x9d, 0xff, 0x21, 0x69, 0x70, 0xfe, 0xab, 0x30, 0x67, 0xee, 0xe7, 0x53, 0x3c,
	0xff, 0x6b, 0xf5, 0x0e, 0xc0, 0x0f, 0x0b, 0xcd, 0x3f, 0x60, 0x2c, 0xee, 0xa8, 0x8d, 0xd9, 0x16,
	0x38, 0x13, 0x70, 0x7d, 0xba, 0xdc, 0x18, 0x2e, 0xe3, 0xb8, 0x51, 0xb3, 0x82, 0x33, 0x77, 0xbf,
	0x28, 0xed, 0x73, 0x08, 0xa9, 0x48, 0xf4, 0x1f, 0x53, 0x69, 0x37, 0x61, 0x11, 0x4b, 0xb4, 0xa5,
	0x5e, 0xb4, 0xb8, 0x84, 0xbb, 0xa5, 0x7d, 0x16, 0x9d, 0x91, 0xba, 0xea, 0x65, 0xb7, 0xb4, 0x17,
	0x61, 0x7e, 0xf0, 0xd8, 0x56, 0xbc, 0x32, 0x5f, 0x76, 0x2f, 0x54, 0xcd, 0x2b, 0x8d, 0xe0, 0x97,
	0x85, 0x16, 0x86, 0x33, 0x6c, 0xe2, 0x34, 0x20, 0x31, 0x18, 0xcf, 0xac, 0x79, 0xa9, 0xb9, 0x51,
	0xb7, 0x84, 0x33, 0x37, 0x9e, 0xe0, 0xb9, 0x04, 0x50, 0xf5, 0x0c, 0x14, 0x7d, 0x54, 0xb4, 0x77,
	0xb6, 0xb5, 0xa8, 0x86, 0xb4, 0xe9, 0x9a, 0xa1, 0xe9, 0xb6, 0xc0, 0x22, 0xe7, 0x50, 0x67, 0x77,
	0x44, 0xce, 0x6b, 0x99, 0xf6, 0x4b, 0xcc, 0x4c, 0xb9, 0xa2, 0x4f, 0x30, 0xd5, 0x90, 0x36, 0x85,
	0xe3, 0x4c, 0x1f, 0x3e, 0x7e, 0xb9, 0x19, 0x46, 0x62, 0x2f, 0xdf, 0xf1, 0x02, 0x9a, 0xf8, 0x72,
	0xda, 0xc1, 0xcd, 0xc0, 0x37, 0xbf, 0xb0, 0xec, 0xcc, 0xa9, 0xeb, 0xc1, 0xcd, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x20, 0xce, 0xc9, 0x36, 0x99, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RulesClient is the client API for Rules service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RulesClient interface {
	CreateRule(ctx context.Context, in *request.CreateRuleReq, opts ...grpc.CallOption) (*response.CreateRuleResp, error)
	UpdateRule(ctx context.Context, in *request.UpdateRuleReq, opts ...grpc.CallOption) (*response.UpdateRuleResp, error)
	GetRule(ctx context.Context, in *request.GetRuleReq, opts ...grpc.CallOption) (*response.GetRuleResp, error)
	ListRulesForProject(ctx context.Context, in *request.ListRulesForProjectReq, opts ...grpc.CallOption) (*response.ListRulesForProjectResp, error)
	DeleteRule(ctx context.Context, in *request.DeleteRuleReq, opts ...grpc.CallOption) (*response.DeleteRuleResp, error)
	ApplyRulesStart(ctx context.Context, in *request.ApplyRulesStartReq, opts ...grpc.CallOption) (*response.ApplyRulesStartResp, error)
	ApplyRulesCancel(ctx context.Context, in *request.ApplyRulesCancelReq, opts ...grpc.CallOption) (*response.ApplyRulesCancelResp, error)
	ApplyRulesStatus(ctx context.Context, in *request.ApplyRulesStatusReq, opts ...grpc.CallOption) (*response.ApplyRulesStatusResp, error)
}

type rulesClient struct {
	cc *grpc.ClientConn
}

func NewRulesClient(cc *grpc.ClientConn) RulesClient {
	return &rulesClient{cc}
}

func (c *rulesClient) CreateRule(ctx context.Context, in *request.CreateRuleReq, opts ...grpc.CallOption) (*response.CreateRuleResp, error) {
	out := new(response.CreateRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/CreateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) UpdateRule(ctx context.Context, in *request.UpdateRuleReq, opts ...grpc.CallOption) (*response.UpdateRuleResp, error) {
	out := new(response.UpdateRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/UpdateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) GetRule(ctx context.Context, in *request.GetRuleReq, opts ...grpc.CallOption) (*response.GetRuleResp, error) {
	out := new(response.GetRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/GetRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ListRulesForProject(ctx context.Context, in *request.ListRulesForProjectReq, opts ...grpc.CallOption) (*response.ListRulesForProjectResp, error) {
	out := new(response.ListRulesForProjectResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/ListRulesForProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) DeleteRule(ctx context.Context, in *request.DeleteRuleReq, opts ...grpc.CallOption) (*response.DeleteRuleResp, error) {
	out := new(response.DeleteRuleResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/DeleteRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ApplyRulesStart(ctx context.Context, in *request.ApplyRulesStartReq, opts ...grpc.CallOption) (*response.ApplyRulesStartResp, error) {
	out := new(response.ApplyRulesStartResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/ApplyRulesStart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ApplyRulesCancel(ctx context.Context, in *request.ApplyRulesCancelReq, opts ...grpc.CallOption) (*response.ApplyRulesCancelResp, error) {
	out := new(response.ApplyRulesCancelResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/ApplyRulesCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) ApplyRulesStatus(ctx context.Context, in *request.ApplyRulesStatusReq, opts ...grpc.CallOption) (*response.ApplyRulesStatusResp, error) {
	out := new(response.ApplyRulesStatusResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2beta.Rules/ApplyRulesStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RulesServer is the server API for Rules service.
type RulesServer interface {
	CreateRule(context.Context, *request.CreateRuleReq) (*response.CreateRuleResp, error)
	UpdateRule(context.Context, *request.UpdateRuleReq) (*response.UpdateRuleResp, error)
	GetRule(context.Context, *request.GetRuleReq) (*response.GetRuleResp, error)
	ListRulesForProject(context.Context, *request.ListRulesForProjectReq) (*response.ListRulesForProjectResp, error)
	DeleteRule(context.Context, *request.DeleteRuleReq) (*response.DeleteRuleResp, error)
	ApplyRulesStart(context.Context, *request.ApplyRulesStartReq) (*response.ApplyRulesStartResp, error)
	ApplyRulesCancel(context.Context, *request.ApplyRulesCancelReq) (*response.ApplyRulesCancelResp, error)
	ApplyRulesStatus(context.Context, *request.ApplyRulesStatusReq) (*response.ApplyRulesStatusResp, error)
}

func RegisterRulesServer(s *grpc.Server, srv RulesServer) {
	s.RegisterService(&_Rules_serviceDesc, srv)
}

func _Rules_CreateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).CreateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/CreateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).CreateRule(ctx, req.(*request.CreateRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/UpdateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).UpdateRule(ctx, req.(*request.UpdateRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_GetRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).GetRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/GetRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).GetRule(ctx, req.(*request.GetRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ListRulesForProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListRulesForProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ListRulesForProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/ListRulesForProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ListRulesForProject(ctx, req.(*request.ListRulesForProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteRuleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/DeleteRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).DeleteRule(ctx, req.(*request.DeleteRuleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ApplyRulesStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ApplyRulesStartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ApplyRulesStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/ApplyRulesStart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ApplyRulesStart(ctx, req.(*request.ApplyRulesStartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ApplyRulesCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ApplyRulesCancelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ApplyRulesCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/ApplyRulesCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ApplyRulesCancel(ctx, req.(*request.ApplyRulesCancelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_ApplyRulesStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ApplyRulesStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).ApplyRulesStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2beta.Rules/ApplyRulesStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).ApplyRulesStatus(ctx, req.(*request.ApplyRulesStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rules_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2beta.Rules",
	HandlerType: (*RulesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRule",
			Handler:    _Rules_CreateRule_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _Rules_UpdateRule_Handler,
		},
		{
			MethodName: "GetRule",
			Handler:    _Rules_GetRule_Handler,
		},
		{
			MethodName: "ListRulesForProject",
			Handler:    _Rules_ListRulesForProject_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _Rules_DeleteRule_Handler,
		},
		{
			MethodName: "ApplyRulesStart",
			Handler:    _Rules_ApplyRulesStart_Handler,
		},
		{
			MethodName: "ApplyRulesCancel",
			Handler:    _Rules_ApplyRulesCancel_Handler,
		},
		{
			MethodName: "ApplyRulesStatus",
			Handler:    _Rules_ApplyRulesStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/iam/v2beta/rules.proto",
}
