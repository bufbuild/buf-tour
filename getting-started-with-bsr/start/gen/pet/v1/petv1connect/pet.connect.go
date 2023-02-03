// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pet/v1/pet.proto

package petv1connect

import (
	context "context"
	errors "errors"
	v1 "github.com/bufbuild/buf-tour/petstore/gen/pet/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// PetStoreServiceName is the fully-qualified name of the PetStoreService service.
	PetStoreServiceName = "pet.v1.PetStoreService"
)

// PetStoreServiceClient is a client for the pet.v1.PetStoreService service.
type PetStoreServiceClient interface {
	GetPet(context.Context, *connect_go.Request[v1.GetPetRequest]) (*connect_go.Response[v1.GetPetResponse], error)
	PutPet(context.Context, *connect_go.Request[v1.PutPetRequest]) (*connect_go.Response[v1.PutPetResponse], error)
	DeletePet(context.Context, *connect_go.Request[v1.DeletePetRequest]) (*connect_go.Response[v1.DeletePetResponse], error)
}

// NewPetStoreServiceClient constructs a client for the pet.v1.PetStoreService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPetStoreServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PetStoreServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &petStoreServiceClient{
		getPet: connect_go.NewClient[v1.GetPetRequest, v1.GetPetResponse](
			httpClient,
			baseURL+"/pet.v1.PetStoreService/GetPet",
			opts...,
		),
		putPet: connect_go.NewClient[v1.PutPetRequest, v1.PutPetResponse](
			httpClient,
			baseURL+"/pet.v1.PetStoreService/PutPet",
			opts...,
		),
		deletePet: connect_go.NewClient[v1.DeletePetRequest, v1.DeletePetResponse](
			httpClient,
			baseURL+"/pet.v1.PetStoreService/DeletePet",
			opts...,
		),
	}
}

// petStoreServiceClient implements PetStoreServiceClient.
type petStoreServiceClient struct {
	getPet    *connect_go.Client[v1.GetPetRequest, v1.GetPetResponse]
	putPet    *connect_go.Client[v1.PutPetRequest, v1.PutPetResponse]
	deletePet *connect_go.Client[v1.DeletePetRequest, v1.DeletePetResponse]
}

// GetPet calls pet.v1.PetStoreService.GetPet.
func (c *petStoreServiceClient) GetPet(ctx context.Context, req *connect_go.Request[v1.GetPetRequest]) (*connect_go.Response[v1.GetPetResponse], error) {
	return c.getPet.CallUnary(ctx, req)
}

// PutPet calls pet.v1.PetStoreService.PutPet.
func (c *petStoreServiceClient) PutPet(ctx context.Context, req *connect_go.Request[v1.PutPetRequest]) (*connect_go.Response[v1.PutPetResponse], error) {
	return c.putPet.CallUnary(ctx, req)
}

// DeletePet calls pet.v1.PetStoreService.DeletePet.
func (c *petStoreServiceClient) DeletePet(ctx context.Context, req *connect_go.Request[v1.DeletePetRequest]) (*connect_go.Response[v1.DeletePetResponse], error) {
	return c.deletePet.CallUnary(ctx, req)
}

// PetStoreServiceHandler is an implementation of the pet.v1.PetStoreService service.
type PetStoreServiceHandler interface {
	GetPet(context.Context, *connect_go.Request[v1.GetPetRequest]) (*connect_go.Response[v1.GetPetResponse], error)
	PutPet(context.Context, *connect_go.Request[v1.PutPetRequest]) (*connect_go.Response[v1.PutPetResponse], error)
	DeletePet(context.Context, *connect_go.Request[v1.DeletePetRequest]) (*connect_go.Response[v1.DeletePetResponse], error)
}

// NewPetStoreServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPetStoreServiceHandler(svc PetStoreServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/pet.v1.PetStoreService/GetPet", connect_go.NewUnaryHandler(
		"/pet.v1.PetStoreService/GetPet",
		svc.GetPet,
		opts...,
	))
	mux.Handle("/pet.v1.PetStoreService/PutPet", connect_go.NewUnaryHandler(
		"/pet.v1.PetStoreService/PutPet",
		svc.PutPet,
		opts...,
	))
	mux.Handle("/pet.v1.PetStoreService/DeletePet", connect_go.NewUnaryHandler(
		"/pet.v1.PetStoreService/DeletePet",
		svc.DeletePet,
		opts...,
	))
	return "/pet.v1.PetStoreService/", mux
}

// UnimplementedPetStoreServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPetStoreServiceHandler struct{}

func (UnimplementedPetStoreServiceHandler) GetPet(context.Context, *connect_go.Request[v1.GetPetRequest]) (*connect_go.Response[v1.GetPetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("pet.v1.PetStoreService.GetPet is not implemented"))
}

func (UnimplementedPetStoreServiceHandler) PutPet(context.Context, *connect_go.Request[v1.PutPetRequest]) (*connect_go.Response[v1.PutPetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("pet.v1.PetStoreService.PutPet is not implemented"))
}

func (UnimplementedPetStoreServiceHandler) DeletePet(context.Context, *connect_go.Request[v1.DeletePetRequest]) (*connect_go.Response[v1.DeletePetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("pet.v1.PetStoreService.DeletePet is not implemented"))
}
