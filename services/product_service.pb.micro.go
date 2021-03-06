// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: product_service.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ProductService service

type ProductService interface {
	GetProductList(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductResponse, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "services"
	}
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) GetProductList(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductResponse, error) {
	req := c.c.NewRequest(c.name, "ProductService.GetProductList", in)
	out := new(ProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductService service

type ProductServiceHandler interface {
	GetProductList(context.Context, *ProductRequest, *ProductResponse) error
}

func RegisterProductServiceHandler(s server.Server, hdlr ProductServiceHandler, opts ...server.HandlerOption) error {
	type productService interface {
		GetProductList(ctx context.Context, in *ProductRequest, out *ProductResponse) error
	}
	type ProductService struct {
		productService
	}
	h := &productServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductService{h}, opts...))
}

type productServiceHandler struct {
	ProductServiceHandler
}

func (h *productServiceHandler) GetProductList(ctx context.Context, in *ProductRequest, out *ProductResponse) error {
	return h.ProductServiceHandler.GetProductList(ctx, in, out)
}
