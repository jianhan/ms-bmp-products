// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/supplier/supplier.proto

/*
Package products is a generated protocol buffer package.

It is generated from these files:
	proto/supplier/supplier.proto

It has these top-level messages:
	UpsertSuppliersReq
	UpsertSuppliersRsp
	Supplier
*/
package products

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
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

// Client API for SuppliersService service

type SuppliersServiceClient interface {
	UpsertSuppliers(ctx context.Context, in *UpsertSuppliersReq, opts ...client.CallOption) (*UpsertSuppliersRsp, error)
}

type suppliersServiceClient struct {
	c           client.Client
	serviceName string
}

func NewSuppliersServiceClient(serviceName string, c client.Client) SuppliersServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.products"
	}
	return &suppliersServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *suppliersServiceClient) UpsertSuppliers(ctx context.Context, in *UpsertSuppliersReq, opts ...client.CallOption) (*UpsertSuppliersRsp, error) {
	req := c.c.NewRequest(c.serviceName, "SuppliersService.UpsertSuppliers", in)
	out := new(UpsertSuppliersRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SuppliersService service

type SuppliersServiceHandler interface {
	UpsertSuppliers(context.Context, *UpsertSuppliersReq, *UpsertSuppliersRsp) error
}

func RegisterSuppliersServiceHandler(s server.Server, hdlr SuppliersServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&SuppliersService{hdlr}, opts...))
}

type SuppliersService struct {
	SuppliersServiceHandler
}

func (h *SuppliersService) UpsertSuppliers(ctx context.Context, in *UpsertSuppliersReq, out *UpsertSuppliersRsp) error {
	return h.SuppliersServiceHandler.UpsertSuppliers(ctx, in, out)
}
