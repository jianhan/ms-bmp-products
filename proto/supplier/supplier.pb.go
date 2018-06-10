// Code generated by protoc-gen-go. DO NOT EDIT.
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

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpsertSuppliersReq struct {
	Suppliers []*Supplier `protobuf:"bytes,1,rep,name=suppliers" json:"suppliers,omitempty"`
}

func (m *UpsertSuppliersReq) Reset()                    { *m = UpsertSuppliersReq{} }
func (m *UpsertSuppliersReq) String() string            { return proto.CompactTextString(m) }
func (*UpsertSuppliersReq) ProtoMessage()               {}
func (*UpsertSuppliersReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpsertSuppliersReq) GetSuppliers() []*Supplier {
	if m != nil {
		return m.Suppliers
	}
	return nil
}

type UpsertSuppliersRsp struct {
	Suppliers []*Supplier `protobuf:"bytes,1,rep,name=suppliers" json:"suppliers,omitempty"`
}

func (m *UpsertSuppliersRsp) Reset()                    { *m = UpsertSuppliersRsp{} }
func (m *UpsertSuppliersRsp) String() string            { return proto.CompactTextString(m) }
func (*UpsertSuppliersRsp) ProtoMessage()               {}
func (*UpsertSuppliersRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpsertSuppliersRsp) GetSuppliers() []*Supplier {
	if m != nil {
		return m.Suppliers
	}
	return nil
}

type Supplier struct {
	ID           string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	LogoUrl      string `protobuf:"bytes,3,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
	HomePageUrl  string `protobuf:"bytes,4,opt,name=home_page_url,json=homePageUrl" json:"home_page_url,omitempty"`
	Currency     string `protobuf:"bytes,5,opt,name=currency" json:"currency,omitempty"`
	DisplayOrder int64  `protobuf:"varint,6,opt,name=display_order,json=displayOrder" json:"display_order,omitempty"`
}

func (m *Supplier) Reset()                    { *m = Supplier{} }
func (m *Supplier) String() string            { return proto.CompactTextString(m) }
func (*Supplier) ProtoMessage()               {}
func (*Supplier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Supplier) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Supplier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Supplier) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

func (m *Supplier) GetHomePageUrl() string {
	if m != nil {
		return m.HomePageUrl
	}
	return ""
}

func (m *Supplier) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Supplier) GetDisplayOrder() int64 {
	if m != nil {
		return m.DisplayOrder
	}
	return 0
}

func init() {
	proto.RegisterType((*UpsertSuppliersReq)(nil), "go.micro.srv.products.UpsertSuppliersReq")
	proto.RegisterType((*UpsertSuppliersRsp)(nil), "go.micro.srv.products.UpsertSuppliersRsp")
	proto.RegisterType((*Supplier)(nil), "go.micro.srv.products.Supplier")
}

func init() { proto.RegisterFile("proto/supplier/supplier.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x49, 0x52, 0x6b, 0x3a, 0xb5, 0x2a, 0x03, 0xc2, 0x5a, 0x10, 0x43, 0xbd, 0xc4, 0x4b,
	0x84, 0x7a, 0xf6, 0x22, 0xbd, 0xf4, 0xa4, 0x24, 0xf4, 0xe2, 0x25, 0xc4, 0x64, 0x88, 0x81, 0xa4,
	0xbb, 0xce, 0x26, 0x85, 0xe2, 0x3b, 0xf9, 0x8c, 0x92, 0xa5, 0x89, 0xa0, 0x15, 0x84, 0xde, 0x66,
	0x7e, 0xff, 0x0f, 0xc8, 0x64, 0xe1, 0x4a, 0xb1, 0xac, 0xe5, 0x9d, 0x6e, 0x94, 0x2a, 0x0b, 0xe2,
	0x7e, 0x08, 0x0c, 0xc7, 0x8b, 0x5c, 0x06, 0x55, 0x91, 0xb2, 0x0c, 0x34, 0x6f, 0x5a, 0x96, 0x35,
	0x69, 0xad, 0x67, 0x11, 0xe0, 0x4a, 0x69, 0xe2, 0x3a, 0xda, 0xd9, 0x75, 0x48, 0xef, 0xf8, 0x00,
	0xa3, 0x2e, 0xae, 0x85, 0xe5, 0x39, 0xfe, 0x78, 0x7e, 0x1d, 0xec, 0x2d, 0x08, 0xba, 0x5c, 0xf8,
	0x9d, 0xd8, 0x57, 0xaa, 0xd5, 0xa1, 0xa5, 0x9f, 0x16, 0xb8, 0x1d, 0xc7, 0x53, 0xb0, 0x97, 0x0b,
	0x61, 0x79, 0x96, 0x3f, 0x0a, 0xed, 0xe5, 0x02, 0x11, 0x06, 0xeb, 0xa4, 0x22, 0x61, 0x1b, 0x62,
	0x66, 0xbc, 0x04, 0xb7, 0x94, 0xb9, 0x8c, 0x1b, 0x2e, 0x85, 0x63, 0xf8, 0x71, 0xbb, 0xaf, 0xb8,
	0xc4, 0x19, 0x4c, 0xde, 0x64, 0x45, 0xb1, 0x4a, 0x72, 0x32, 0xfa, 0xc0, 0xe8, 0xe3, 0x16, 0x3e,
	0x27, 0x39, 0xb5, 0x9e, 0x29, 0xb8, 0x69, 0xc3, 0x4c, 0xeb, 0x74, 0x2b, 0x8e, 0x8c, 0xdc, 0xef,
	0x78, 0x03, 0x93, 0xac, 0xd0, 0xaa, 0x4c, 0xb6, 0xb1, 0xe4, 0x8c, 0x58, 0x0c, 0x3d, 0xcb, 0x77,
	0xc2, 0x93, 0x1d, 0x7c, 0x6a, 0xd9, 0xfc, 0x03, 0xce, 0xfb, 0xef, 0x8f, 0x88, 0x37, 0x45, 0x4a,
	0x98, 0xc3, 0xd9, 0x8f, 0xcb, 0xe0, 0xed, 0x1f, 0x37, 0xf8, 0xfd, 0x5b, 0xa6, 0xff, 0xb5, 0x6a,
	0xf5, 0x08, 0x2f, 0x6e, 0x27, 0xbf, 0x0e, 0xcd, 0x0b, 0xb8, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x9a, 0x86, 0x72, 0xa9, 0x22, 0x02, 0x00, 0x00,
}