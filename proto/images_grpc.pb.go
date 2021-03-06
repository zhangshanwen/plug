// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ImageClient is the client API for Image service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageClient interface {
	ImageList(ctx context.Context, in *ImageListRequest, opts ...grpc.CallOption) (*ImageListReply, error)
	ImageHistory(ctx context.Context, in *ImageHistoryRequest, opts ...grpc.CallOption) (*ImageHistoryReply, error)
	ImageSearch(ctx context.Context, in *ImageSearchRequest, opts ...grpc.CallOption) (*ImageSearchReply, error)
}

type imageClient struct {
	cc grpc.ClientConnInterface
}

func NewImageClient(cc grpc.ClientConnInterface) ImageClient {
	return &imageClient{cc}
}

func (c *imageClient) ImageList(ctx context.Context, in *ImageListRequest, opts ...grpc.CallOption) (*ImageListReply, error) {
	out := new(ImageListReply)
	err := c.cc.Invoke(ctx, "/proto.Image/ImageList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) ImageHistory(ctx context.Context, in *ImageHistoryRequest, opts ...grpc.CallOption) (*ImageHistoryReply, error) {
	out := new(ImageHistoryReply)
	err := c.cc.Invoke(ctx, "/proto.Image/ImageHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageClient) ImageSearch(ctx context.Context, in *ImageSearchRequest, opts ...grpc.CallOption) (*ImageSearchReply, error) {
	out := new(ImageSearchReply)
	err := c.cc.Invoke(ctx, "/proto.Image/ImageSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServer is the server API for Image service.
// All implementations must embed UnimplementedImageServer
// for forward compatibility
type ImageServer interface {
	ImageList(context.Context, *ImageListRequest) (*ImageListReply, error)
	ImageHistory(context.Context, *ImageHistoryRequest) (*ImageHistoryReply, error)
	ImageSearch(context.Context, *ImageSearchRequest) (*ImageSearchReply, error)
	mustEmbedUnimplementedImageServer()
}

// UnimplementedImageServer must be embedded to have forward compatible implementations.
type UnimplementedImageServer struct {
}

func (UnimplementedImageServer) ImageList(context.Context, *ImageListRequest) (*ImageListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageList not implemented")
}
func (UnimplementedImageServer) ImageHistory(context.Context, *ImageHistoryRequest) (*ImageHistoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageHistory not implemented")
}
func (UnimplementedImageServer) ImageSearch(context.Context, *ImageSearchRequest) (*ImageSearchReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImageSearch not implemented")
}
func (UnimplementedImageServer) mustEmbedUnimplementedImageServer() {}

// UnsafeImageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServer will
// result in compilation errors.
type UnsafeImageServer interface {
	mustEmbedUnimplementedImageServer()
}

func RegisterImageServer(s grpc.ServiceRegistrar, srv ImageServer) {
	s.RegisterService(&Image_ServiceDesc, srv)
}

func _Image_ImageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).ImageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Image/ImageList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).ImageList(ctx, req.(*ImageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_ImageHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).ImageHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Image/ImageHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).ImageHistory(ctx, req.(*ImageHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Image_ImageSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).ImageSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Image/ImageSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).ImageSearch(ctx, req.(*ImageSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Image_ServiceDesc is the grpc.ServiceDesc for Image service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Image_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Image",
	HandlerType: (*ImageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImageList",
			Handler:    _Image_ImageList_Handler,
		},
		{
			MethodName: "ImageHistory",
			Handler:    _Image_ImageHistory_Handler,
		},
		{
			MethodName: "ImageSearch",
			Handler:    _Image_ImageSearch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/images.proto",
}
