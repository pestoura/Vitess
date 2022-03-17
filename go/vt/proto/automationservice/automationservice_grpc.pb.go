// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package automationservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	automation "vitess.io/vitess/go/vt/proto/automation"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AutomationClient is the client API for Automation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutomationClient interface {
	// Start a cluster operation.
	EnqueueClusterOperation(ctx context.Context, in *automation.EnqueueClusterOperationRequest, opts ...grpc.CallOption) (*automation.EnqueueClusterOperationResponse, error)
	// TODO(mberlin): Polling this is bad. Implement a subscribe mechanism to wait for changes?
	// Get all details of an active cluster operation.
	GetClusterOperationDetails(ctx context.Context, in *automation.GetClusterOperationDetailsRequest, opts ...grpc.CallOption) (*automation.GetClusterOperationDetailsResponse, error)
}

type automationClient struct {
	cc grpc.ClientConnInterface
}

func NewAutomationClient(cc grpc.ClientConnInterface) AutomationClient {
	return &automationClient{cc}
}

func (c *automationClient) EnqueueClusterOperation(ctx context.Context, in *automation.EnqueueClusterOperationRequest, opts ...grpc.CallOption) (*automation.EnqueueClusterOperationResponse, error) {
	out := new(automation.EnqueueClusterOperationResponse)
	err := c.cc.Invoke(ctx, "/automationservice.Automation/EnqueueClusterOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationClient) GetClusterOperationDetails(ctx context.Context, in *automation.GetClusterOperationDetailsRequest, opts ...grpc.CallOption) (*automation.GetClusterOperationDetailsResponse, error) {
	out := new(automation.GetClusterOperationDetailsResponse)
	err := c.cc.Invoke(ctx, "/automationservice.Automation/GetClusterOperationDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutomationServer is the server API for Automation service.
// All implementations must embed UnimplementedAutomationServer
// for forward compatibility
type AutomationServer interface {
	// Start a cluster operation.
	EnqueueClusterOperation(context.Context, *automation.EnqueueClusterOperationRequest) (*automation.EnqueueClusterOperationResponse, error)
	// TODO(mberlin): Polling this is bad. Implement a subscribe mechanism to wait for changes?
	// Get all details of an active cluster operation.
	GetClusterOperationDetails(context.Context, *automation.GetClusterOperationDetailsRequest) (*automation.GetClusterOperationDetailsResponse, error)
	mustEmbedUnimplementedAutomationServer()
}

// UnimplementedAutomationServer must be embedded to have forward compatible implementations.
type UnimplementedAutomationServer struct {
}

func (UnimplementedAutomationServer) EnqueueClusterOperation(context.Context, *automation.EnqueueClusterOperationRequest) (*automation.EnqueueClusterOperationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnqueueClusterOperation not implemented")
}
func (UnimplementedAutomationServer) GetClusterOperationDetails(context.Context, *automation.GetClusterOperationDetailsRequest) (*automation.GetClusterOperationDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterOperationDetails not implemented")
}
func (UnimplementedAutomationServer) mustEmbedUnimplementedAutomationServer() {}

// UnsafeAutomationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutomationServer will
// result in compilation errors.
type UnsafeAutomationServer interface {
	mustEmbedUnimplementedAutomationServer()
}

func RegisterAutomationServer(s grpc.ServiceRegistrar, srv AutomationServer) {
	s.RegisterService(&Automation_ServiceDesc, srv)
}

func _Automation_EnqueueClusterOperation_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(automation.EnqueueClusterOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServer).EnqueueClusterOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automationservice.Automation/EnqueueClusterOperation",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(AutomationServer).EnqueueClusterOperation(ctx, req.(*automation.EnqueueClusterOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Automation_GetClusterOperationDetails_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(automation.GetClusterOperationDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServer).GetClusterOperationDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/automationservice.Automation/GetClusterOperationDetails",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(AutomationServer).GetClusterOperationDetails(ctx, req.(*automation.GetClusterOperationDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Automation_ServiceDesc is the grpc.ServiceDesc for Automation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Automation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "automationservice.Automation",
	HandlerType: (*AutomationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnqueueClusterOperation",
			Handler:    _Automation_EnqueueClusterOperation_Handler,
		},
		{
			MethodName: "GetClusterOperationDetails",
			Handler:    _Automation_GetClusterOperationDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "automationservice.proto",
}
