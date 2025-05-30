// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: note.proto

package note

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NoteManager_CreateNote_FullMethodName  = "/note.NoteManager/CreateNote"
	NoteManager_GetNoteById_FullMethodName = "/note.NoteManager/GetNoteById"
)

// NoteManagerClient is the client API for NoteManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoteManagerClient interface {
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	GetNoteById(ctx context.Context, in *GetNoteByIdRequest, opts ...grpc.CallOption) (*GetNoteByIdResponse, error)
}

type noteManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewNoteManagerClient(cc grpc.ClientConnInterface) NoteManagerClient {
	return &noteManagerClient{cc}
}

func (c *noteManagerClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, NoteManager_CreateNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteManagerClient) GetNoteById(ctx context.Context, in *GetNoteByIdRequest, opts ...grpc.CallOption) (*GetNoteByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNoteByIdResponse)
	err := c.cc.Invoke(ctx, NoteManager_GetNoteById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteManagerServer is the server API for NoteManager service.
// All implementations must embed UnimplementedNoteManagerServer
// for forward compatibility.
type NoteManagerServer interface {
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	GetNoteById(context.Context, *GetNoteByIdRequest) (*GetNoteByIdResponse, error)
	mustEmbedUnimplementedNoteManagerServer()
}

// UnimplementedNoteManagerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNoteManagerServer struct{}

func (UnimplementedNoteManagerServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNoteManagerServer) GetNoteById(context.Context, *GetNoteByIdRequest) (*GetNoteByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoteById not implemented")
}
func (UnimplementedNoteManagerServer) mustEmbedUnimplementedNoteManagerServer() {}
func (UnimplementedNoteManagerServer) testEmbeddedByValue()                     {}

// UnsafeNoteManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoteManagerServer will
// result in compilation errors.
type UnsafeNoteManagerServer interface {
	mustEmbedUnimplementedNoteManagerServer()
}

func RegisterNoteManagerServer(s grpc.ServiceRegistrar, srv NoteManagerServer) {
	// If the following call pancis, it indicates UnimplementedNoteManagerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NoteManager_ServiceDesc, srv)
}

func _NoteManager_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteManagerServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteManager_CreateNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteManagerServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteManager_GetNoteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNoteByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteManagerServer).GetNoteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteManager_GetNoteById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteManagerServer).GetNoteById(ctx, req.(*GetNoteByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NoteManager_ServiceDesc is the grpc.ServiceDesc for NoteManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoteManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "note.NoteManager",
	HandlerType: (*NoteManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNote",
			Handler:    _NoteManager_CreateNote_Handler,
		},
		{
			MethodName: "GetNoteById",
			Handler:    _NoteManager_GetNoteById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "note.proto",
}
