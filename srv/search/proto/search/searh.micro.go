// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: searh.proto

/*
Package go_tutu_srv_search is a generated protocol buffer package.

It is generated from these files:
	searh.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_search

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Search service

type SearchService interface {
	SearchList(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type searchService struct {
	c    client.Client
	name string
}

func NewSearchService(name string, c client.Client) SearchService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.tutu.srv.search"
	}
	return &searchService{
		c:    c,
		name: name,
	}
}

func (c *searchService) SearchList(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Search.SearchList", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Search service

type SearchHandler interface {
	SearchList(context.Context, *Request, *Response) error
}

func RegisterSearchHandler(s server.Server, hdlr SearchHandler, opts ...server.HandlerOption) error {
	type search interface {
		SearchList(ctx context.Context, in *Request, out *Response) error
	}
	type Search struct {
		search
	}
	h := &searchHandler{hdlr}
	return s.Handle(s.NewHandler(&Search{h}, opts...))
}

type searchHandler struct {
	SearchHandler
}

func (h *searchHandler) SearchList(ctx context.Context, in *Request, out *Response) error {
	return h.SearchHandler.SearchList(ctx, in, out)
}
