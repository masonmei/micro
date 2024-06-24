// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/runtime/v1/source.proto

package v1

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/smart-echo/micro/api"
	client "github.com/smart-echo/micro/client"
	server "github.com/smart-echo/micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for SourceService service

func NewSourceServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SourceService service

type SourceService interface {
	Upload(ctx context.Context, opts ...client.CallOption) (SourceService_UploadService, error)
}

type sourceService struct {
	c    client.Client
	name string
}

func NewSourceService(name string, c client.Client) SourceService {
	return &sourceService{
		c:    c,
		name: name,
	}
}

func (c *sourceService) Upload(ctx context.Context, opts ...client.CallOption) (SourceService_UploadService, error) {
	req := c.c.NewRequest(c.name, "SourceService.Upload", &UploadRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &sourceServiceUpload{stream}, nil
}

type SourceService_UploadService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*UploadRequest) error
}

type sourceServiceUpload struct {
	stream client.Stream
}

func (x *sourceServiceUpload) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *sourceServiceUpload) Close() error {
	return x.stream.Close()
}

func (x *sourceServiceUpload) Context() context.Context {
	return x.stream.Context()
}

func (x *sourceServiceUpload) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceServiceUpload) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceServiceUpload) Send(m *UploadRequest) error {
	return x.stream.Send(m)
}

// Server API for SourceService service

type SourceServiceHandler interface {
	Upload(context.Context, SourceService_UploadStream) error
}

func RegisterSourceServiceHandler(s server.Server, hdlr SourceServiceHandler, opts ...server.HandlerOption) error {
	type sourceService interface {
		Upload(ctx context.Context, stream server.Stream) error
	}
	type SourceService struct {
		sourceService
	}
	h := &sourceServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SourceService{h}, opts...))
}

type sourceServiceHandler struct {
	SourceServiceHandler
}

func (h *sourceServiceHandler) Upload(ctx context.Context, stream server.Stream) error {
	return h.SourceServiceHandler.Upload(ctx, &sourceServiceUploadStream{stream})
}

type SourceService_UploadStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*UploadRequest, error)
}

type sourceServiceUploadStream struct {
	stream server.Stream
}

func (x *sourceServiceUploadStream) Close() error {
	return x.stream.Close()
}

func (x *sourceServiceUploadStream) Context() context.Context {
	return x.stream.Context()
}

func (x *sourceServiceUploadStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceServiceUploadStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceServiceUploadStream) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}