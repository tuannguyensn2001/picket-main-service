// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: answer_sheet.proto

package answersheetpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AnswerSheetServiceClient is the client API for AnswerSheetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnswerSheetServiceClient interface {
	StartDoTest(ctx context.Context, in *StartDoTestRequest, opts ...grpc.CallOption) (*StartDoTestResponse, error)
	CheckUserDoingTest(ctx context.Context, in *CheckUserDoingTestRequest, opts ...grpc.CallOption) (*CheckUserDoingTestResponse, error)
	UserAnswer(ctx context.Context, in *UserAnswerRequest, opts ...grpc.CallOption) (*UserAnswerResponse, error)
	GetTestContent(ctx context.Context, in *GetTestContentRequest, opts ...grpc.CallOption) (*GetTestContentResponse, error)
	GetLatestStartTime(ctx context.Context, in *GetLatestStartTimeRequest, opts ...grpc.CallOption) (*GetLatestStartTimeResponse, error)
	GetCurrentTest(ctx context.Context, in *GetCurrentTestRequest, opts ...grpc.CallOption) (*GetCurrentTestResponse, error)
	SubmitTest(ctx context.Context, in *SubmitTestRequest, opts ...grpc.CallOption) (*SubmitTestResponse, error)
	CheckUserSubmitted(ctx context.Context, in *CheckUserSubmittedRequest, opts ...grpc.CallOption) (*CheckUserSubmittedResponse, error)
}

type answerSheetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnswerSheetServiceClient(cc grpc.ClientConnInterface) AnswerSheetServiceClient {
	return &answerSheetServiceClient{cc}
}

func (c *answerSheetServiceClient) StartDoTest(ctx context.Context, in *StartDoTestRequest, opts ...grpc.CallOption) (*StartDoTestResponse, error) {
	out := new(StartDoTestResponse)
	err := c.cc.Invoke(ctx, "/answer_sheet.AnswerSheetService/StartDoTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerSheetServiceClient) CheckUserDoingTest(ctx context.Context, in *CheckUserDoingTestRequest, opts ...grpc.CallOption) (*CheckUserDoingTestResponse, error) {
	out := new(CheckUserDoingTestResponse)
	err := c.cc.Invoke(ctx, "/answer_sheet.AnswerSheetService/CheckUserDoingTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerSheetServiceClient) UserAnswer(ctx context.Context, in *UserAnswerRequest, opts ...grpc.CallOption) (*UserAnswerResponse, error) {
	out := new(UserAnswerResponse)
	err := c.cc.Invoke(ctx, "/answer_sheet.AnswerSheetService/UserAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerSheetServiceClient) GetTestContent(ctx context.Context, in *GetTestContentRequest, opts ...grpc.CallOption) (*GetTestContentResponse, error) {
	out := new(GetTestContentResponse)
	err := c.cc.Invoke(ctx, "/answer_sheet.AnswerSheetService/GetTestContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerSheetServiceClient) GetLatestStartTime(ctx context.Context, in *GetLatestStartTimeRequest, opts ...grpc.CallOption) (*GetLatestStartTimeResponse, error) {
	out := new(GetLatestStartTimeResponse)
	err := c.cc.Invoke(ctx, "/answer_sheet.AnswerSheetService/GetLatestStartTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerSheetServiceClient) GetCurrentTest(ctx context.Context, in *GetCurrentTestRequest, opts ...grpc.CallOption) (*GetCurrentTestResponse, error) {
	out := new(GetCurrentTestResponse)
	err := c.cc.Invoke(ctx, "/answer_sheet.AnswerSheetService/GetCurrentTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerSheetServiceClient) SubmitTest(ctx context.Context, in *SubmitTestRequest, opts ...grpc.CallOption) (*SubmitTestResponse, error) {
	out := new(SubmitTestResponse)
	err := c.cc.Invoke(ctx, "/answer_sheet.AnswerSheetService/SubmitTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *answerSheetServiceClient) CheckUserSubmitted(ctx context.Context, in *CheckUserSubmittedRequest, opts ...grpc.CallOption) (*CheckUserSubmittedResponse, error) {
	out := new(CheckUserSubmittedResponse)
	err := c.cc.Invoke(ctx, "/answer_sheet.AnswerSheetService/CheckUserSubmitted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnswerSheetServiceServer is the server API for AnswerSheetService service.
// All implementations must embed UnimplementedAnswerSheetServiceServer
// for forward compatibility
type AnswerSheetServiceServer interface {
	StartDoTest(context.Context, *StartDoTestRequest) (*StartDoTestResponse, error)
	CheckUserDoingTest(context.Context, *CheckUserDoingTestRequest) (*CheckUserDoingTestResponse, error)
	UserAnswer(context.Context, *UserAnswerRequest) (*UserAnswerResponse, error)
	GetTestContent(context.Context, *GetTestContentRequest) (*GetTestContentResponse, error)
	GetLatestStartTime(context.Context, *GetLatestStartTimeRequest) (*GetLatestStartTimeResponse, error)
	GetCurrentTest(context.Context, *GetCurrentTestRequest) (*GetCurrentTestResponse, error)
	SubmitTest(context.Context, *SubmitTestRequest) (*SubmitTestResponse, error)
	CheckUserSubmitted(context.Context, *CheckUserSubmittedRequest) (*CheckUserSubmittedResponse, error)
	mustEmbedUnimplementedAnswerSheetServiceServer()
}

// UnimplementedAnswerSheetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAnswerSheetServiceServer struct {
}

func (UnimplementedAnswerSheetServiceServer) StartDoTest(context.Context, *StartDoTestRequest) (*StartDoTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartDoTest not implemented")
}
func (UnimplementedAnswerSheetServiceServer) CheckUserDoingTest(context.Context, *CheckUserDoingTestRequest) (*CheckUserDoingTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserDoingTest not implemented")
}
func (UnimplementedAnswerSheetServiceServer) UserAnswer(context.Context, *UserAnswerRequest) (*UserAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAnswer not implemented")
}
func (UnimplementedAnswerSheetServiceServer) GetTestContent(context.Context, *GetTestContentRequest) (*GetTestContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestContent not implemented")
}
func (UnimplementedAnswerSheetServiceServer) GetLatestStartTime(context.Context, *GetLatestStartTimeRequest) (*GetLatestStartTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestStartTime not implemented")
}
func (UnimplementedAnswerSheetServiceServer) GetCurrentTest(context.Context, *GetCurrentTestRequest) (*GetCurrentTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentTest not implemented")
}
func (UnimplementedAnswerSheetServiceServer) SubmitTest(context.Context, *SubmitTestRequest) (*SubmitTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTest not implemented")
}
func (UnimplementedAnswerSheetServiceServer) CheckUserSubmitted(context.Context, *CheckUserSubmittedRequest) (*CheckUserSubmittedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserSubmitted not implemented")
}
func (UnimplementedAnswerSheetServiceServer) mustEmbedUnimplementedAnswerSheetServiceServer() {}

// UnsafeAnswerSheetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnswerSheetServiceServer will
// result in compilation errors.
type UnsafeAnswerSheetServiceServer interface {
	mustEmbedUnimplementedAnswerSheetServiceServer()
}

func RegisterAnswerSheetServiceServer(s grpc.ServiceRegistrar, srv AnswerSheetServiceServer) {
	s.RegisterService(&AnswerSheetService_ServiceDesc, srv)
}

func _AnswerSheetService_StartDoTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartDoTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerSheetServiceServer).StartDoTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answer_sheet.AnswerSheetService/StartDoTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerSheetServiceServer).StartDoTest(ctx, req.(*StartDoTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerSheetService_CheckUserDoingTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserDoingTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerSheetServiceServer).CheckUserDoingTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answer_sheet.AnswerSheetService/CheckUserDoingTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerSheetServiceServer).CheckUserDoingTest(ctx, req.(*CheckUserDoingTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerSheetService_UserAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerSheetServiceServer).UserAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answer_sheet.AnswerSheetService/UserAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerSheetServiceServer).UserAnswer(ctx, req.(*UserAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerSheetService_GetTestContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTestContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerSheetServiceServer).GetTestContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answer_sheet.AnswerSheetService/GetTestContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerSheetServiceServer).GetTestContent(ctx, req.(*GetTestContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerSheetService_GetLatestStartTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestStartTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerSheetServiceServer).GetLatestStartTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answer_sheet.AnswerSheetService/GetLatestStartTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerSheetServiceServer).GetLatestStartTime(ctx, req.(*GetLatestStartTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerSheetService_GetCurrentTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerSheetServiceServer).GetCurrentTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answer_sheet.AnswerSheetService/GetCurrentTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerSheetServiceServer).GetCurrentTest(ctx, req.(*GetCurrentTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerSheetService_SubmitTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitTestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerSheetServiceServer).SubmitTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answer_sheet.AnswerSheetService/SubmitTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerSheetServiceServer).SubmitTest(ctx, req.(*SubmitTestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnswerSheetService_CheckUserSubmitted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserSubmittedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnswerSheetServiceServer).CheckUserSubmitted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/answer_sheet.AnswerSheetService/CheckUserSubmitted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnswerSheetServiceServer).CheckUserSubmitted(ctx, req.(*CheckUserSubmittedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnswerSheetService_ServiceDesc is the grpc.ServiceDesc for AnswerSheetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnswerSheetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "answer_sheet.AnswerSheetService",
	HandlerType: (*AnswerSheetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartDoTest",
			Handler:    _AnswerSheetService_StartDoTest_Handler,
		},
		{
			MethodName: "CheckUserDoingTest",
			Handler:    _AnswerSheetService_CheckUserDoingTest_Handler,
		},
		{
			MethodName: "UserAnswer",
			Handler:    _AnswerSheetService_UserAnswer_Handler,
		},
		{
			MethodName: "GetTestContent",
			Handler:    _AnswerSheetService_GetTestContent_Handler,
		},
		{
			MethodName: "GetLatestStartTime",
			Handler:    _AnswerSheetService_GetLatestStartTime_Handler,
		},
		{
			MethodName: "GetCurrentTest",
			Handler:    _AnswerSheetService_GetCurrentTest_Handler,
		},
		{
			MethodName: "SubmitTest",
			Handler:    _AnswerSheetService_SubmitTest_Handler,
		},
		{
			MethodName: "CheckUserSubmitted",
			Handler:    _AnswerSheetService_CheckUserSubmitted_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "answer_sheet.proto",
}
