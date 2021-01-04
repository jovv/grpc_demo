// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MovieCatalogueClient is the client API for MovieCatalogue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieCatalogueClient interface {
	GetMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieReply, error)
	GetMovies(ctx context.Context, in *MoviesRequest, opts ...grpc.CallOption) (*MoviesReply, error)
}

type movieCatalogueClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieCatalogueClient(cc grpc.ClientConnInterface) MovieCatalogueClient {
	return &movieCatalogueClient{cc}
}

func (c *movieCatalogueClient) GetMovie(ctx context.Context, in *MovieRequest, opts ...grpc.CallOption) (*MovieReply, error) {
	out := new(MovieReply)
	err := c.cc.Invoke(ctx, "/grpc.MovieCatalogue/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCatalogueClient) GetMovies(ctx context.Context, in *MoviesRequest, opts ...grpc.CallOption) (*MoviesReply, error) {
	out := new(MoviesReply)
	err := c.cc.Invoke(ctx, "/grpc.MovieCatalogue/GetMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieCatalogueServer is the server API for MovieCatalogue service.
// All implementations must embed UnimplementedMovieCatalogueServer
// for forward compatibility
type MovieCatalogueServer interface {
	GetMovie(context.Context, *MovieRequest) (*MovieReply, error)
	GetMovies(context.Context, *MoviesRequest) (*MoviesReply, error)
	mustEmbedUnimplementedMovieCatalogueServer()
}

// UnimplementedMovieCatalogueServer must be embedded to have forward compatible implementations.
type UnimplementedMovieCatalogueServer struct {
}

func (UnimplementedMovieCatalogueServer) GetMovie(context.Context, *MovieRequest) (*MovieReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedMovieCatalogueServer) GetMovies(context.Context, *MoviesRequest) (*MoviesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
}
func (UnimplementedMovieCatalogueServer) mustEmbedUnimplementedMovieCatalogueServer() {}

// UnsafeMovieCatalogueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieCatalogueServer will
// result in compilation errors.
type UnsafeMovieCatalogueServer interface {
	mustEmbedUnimplementedMovieCatalogueServer()
}

func RegisterMovieCatalogueServer(s grpc.ServiceRegistrar, srv MovieCatalogueServer) {
	s.RegisterService(&_MovieCatalogue_serviceDesc, srv)
}

func _MovieCatalogue_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCatalogueServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MovieCatalogue/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCatalogueServer).GetMovie(ctx, req.(*MovieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCatalogue_GetMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCatalogueServer).GetMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MovieCatalogue/GetMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCatalogueServer).GetMovies(ctx, req.(*MoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieCatalogue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MovieCatalogue",
	HandlerType: (*MovieCatalogueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _MovieCatalogue_GetMovie_Handler,
		},
		{
			MethodName: "GetMovies",
			Handler:    _MovieCatalogue_GetMovies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie_catalogue.proto",
}
