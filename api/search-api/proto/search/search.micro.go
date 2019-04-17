// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: search.proto

/*
Package go_micro_api_search is a generated protocol buffer package.

It is generated from these files:
	search.proto

It has these top-level messages:
*/
package go_micro_api_search

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_api "github.com/micro/go-api/proto"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = go_api.Response{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SearchEr service

type SearchErService interface {
	SearchList(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
}

type searchErService struct {
	c    client.Client
	name string
}

func NewSearchErService(name string, c client.Client) SearchErService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.api.search"
	}
	return &searchErService{
		c:    c,
		name: name,
	}
}

func (c *searchErService) SearchList(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "SearchEr.SearchList", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SearchEr service

type SearchErHandler interface {
	SearchList(context.Context, *go_api.Request, *go_api.Response) error
}

func RegisterSearchErHandler(s server.Server, hdlr SearchErHandler, opts ...server.HandlerOption) error {
	type searchEr interface {
		SearchList(ctx context.Context, in *go_api.Request, out *go_api.Response) error
	}
	type SearchEr struct {
		searchEr
	}
	h := &searchErHandler{hdlr}
	return s.Handle(s.NewHandler(&SearchEr{h}, opts...))
}

type searchErHandler struct {
	SearchErHandler
}

func (h *searchErHandler) SearchList(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.SearchErHandler.SearchList(ctx, in, out)
}