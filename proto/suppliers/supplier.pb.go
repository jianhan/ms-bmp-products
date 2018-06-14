// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/suppliers/suppliers.proto

/*
Package products is a generated protocol buffer package.

It is generated from these files:
	proto/suppliers/suppliers.proto

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
	// @inject_tag: valid:"uuidv4~suppliers ID must be a valid UUIDv4" conform:"trim"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" valid:"uuidv4~suppliers ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: valid:"required~Supplier name is required" conform:"trim"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" valid:"required~Supplier name is required" conform:"trim"`
	// @inject_tag: valid:"required~Supplier slug is required" conform:"trim,slug"
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty" valid:"required~Supplier slug is required" conform:"trim,slug"`
	// @inject_tag: valid:"url~Supplier logo must be a valid URL" conform:"trim"
	LogoUrl string `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty" valid:"url~Supplier logo must be a valid URL" conform:"trim"`
	// @inject_tag: valid:"required, url~Supplier home page must be a valid URL" conform:"trim"
	HomePageUrl string `protobuf:"bytes,5,opt,name=home_page_url,json=homePageUrl" json:"home_page_url,omitempty" valid:"required, url~Supplier home page must be a valid URL" conform:"trim"`
	// @inject_tag: conform:"trim"
	Currency     string `protobuf:"bytes,6,opt,name=currency" json:"currency,omitempty" conform:"trim"`
	DisplayOrder int64  `protobuf:"varint,7,opt,name=display_order,json=displayOrder" json:"display_order,omitempty"`
	CreatedAt    int64  `protobuf:"varint,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt    int64  `protobuf:"varint,9,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
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

func (m *Supplier) GetSlug() string {
	if m != nil {
		return m.Slug
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

func (m *Supplier) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Supplier) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*UpsertSuppliersReq)(nil), "go.micro.srv.products.UpsertSuppliersReq")
	proto.RegisterType((*UpsertSuppliersRsp)(nil), "go.micro.srv.products.UpsertSuppliersRsp")
	proto.RegisterType((*Supplier)(nil), "go.micro.srv.products.Supplier")
}

func init() { proto.RegisterFile("proto/suppliers/suppliers.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4b, 0xeb, 0x40,
	0x10, 0x87, 0x49, 0xda, 0xd7, 0x26, 0xd3, 0x57, 0x95, 0x05, 0x61, 0x2d, 0x14, 0x4b, 0xbd, 0xd4,
	0x4b, 0x84, 0x7a, 0xf6, 0x50, 0xe9, 0xa5, 0x27, 0x25, 0xa5, 0x17, 0x2f, 0x21, 0x26, 0x43, 0x0c,
	0xa4, 0xdd, 0x75, 0x76, 0xb7, 0x50, 0xbc, 0xf8, 0xa7, 0xcb, 0x6e, 0x93, 0x0a, 0x5a, 0x41, 0xf0,
	0x36, 0xf3, 0x7d, 0xbf, 0x99, 0xc3, 0xec, 0xc2, 0x50, 0x92, 0xd0, 0xe2, 0x46, 0x19, 0x29, 0xab,
	0x12, 0xe9, 0x50, 0x44, 0x8e, 0xb3, 0xf3, 0x42, 0x44, 0xeb, 0x32, 0x23, 0x11, 0x29, 0xda, 0x5a,
	0x96, 0x9b, 0x4c, 0xab, 0xf1, 0x12, 0xd8, 0x4a, 0x2a, 0x24, 0xbd, 0xac, 0xe3, 0x2a, 0xc6, 0x57,
	0x76, 0x07, 0x61, 0x33, 0xae, 0xb8, 0x37, 0x6a, 0x4d, 0x7a, 0xd3, 0xcb, 0xe8, 0xe8, 0x82, 0xa8,
	0x99, 0x8b, 0x3f, 0x27, 0x8e, 0x2d, 0x55, 0xf2, 0xaf, 0x4b, 0xdf, 0x7d, 0x08, 0x1a, 0xce, 0x4e,
	0xc0, 0x5f, 0xcc, 0xb9, 0x37, 0xf2, 0x26, 0x61, 0xec, 0x2f, 0xe6, 0x8c, 0x41, 0x7b, 0x93, 0xae,
	0x91, 0xfb, 0x8e, 0xb8, 0xda, 0x32, 0x55, 0x99, 0x82, 0xb7, 0xf6, 0xcc, 0xd6, 0xec, 0x02, 0x82,
	0x4a, 0x14, 0x22, 0x31, 0x54, 0xf1, 0xb6, 0xe3, 0x5d, 0xdb, 0xaf, 0xa8, 0x62, 0x63, 0xe8, 0xbf,
	0x88, 0x35, 0x26, 0x32, 0x2d, 0xd0, 0xf9, 0x7f, 0xce, 0xf7, 0x2c, 0x7c, 0x4c, 0x0b, 0xb4, 0x99,
	0x01, 0x04, 0x99, 0x21, 0xc2, 0x4d, 0xb6, 0xe3, 0x1d, 0xa7, 0x0f, 0x3d, 0xbb, 0x82, 0x7e, 0x5e,
	0x2a, 0x59, 0xa5, 0xbb, 0x44, 0x50, 0x8e, 0xc4, 0xbb, 0x23, 0x6f, 0xd2, 0x8a, 0xff, 0xd7, 0xf0,
	0xc1, 0x32, 0x36, 0x04, 0xc8, 0x08, 0x53, 0x8d, 0x79, 0x92, 0x6a, 0x1e, 0xb8, 0x44, 0x58, 0x93,
	0x99, 0xb6, 0xda, 0xc8, 0xbc, 0xd1, 0xe1, 0x5e, 0xd7, 0x64, 0xa6, 0xa7, 0x6f, 0x70, 0x76, 0xb8,
	0xe8, 0x12, 0x69, 0x5b, 0x66, 0xc8, 0x0a, 0x38, 0xfd, 0x72, 0x6b, 0x76, 0xfd, 0xc3, 0x55, 0xbf,
	0x3f, 0xf4, 0xe0, 0xb7, 0x51, 0x25, 0xef, 0xe1, 0x29, 0x68, 0xf4, 0x73, 0xc7, 0xfd, 0xa9, 0xdb,
	0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x79, 0xa3, 0xda, 0x74, 0x02, 0x00, 0x00,
}