// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/suppliers/supplier.proto

/*
Package suppliers is a generated protocol buffer package.

It is generated from these files:
	proto/suppliers/supplier.proto

It has these top-level messages:
	UpsertSuppliersReq
	UpsertSuppliersRsp
	Supplier
	SuppliersReq
	SuppliersRsp
*/
package suppliers

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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
	// @inject_tag: bson:"_id" valid:"uuidv4~suppliers ID must be a valid UUIDv4" conform:"trim"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"uuidv4~suppliers ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"name" valid:"required~Supplier name is required" conform:"trim"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" bson:"name" valid:"required~Supplier name is required" conform:"trim"`
	// @inject_tag: bson:"slug" valid:"required~Supplier slug is required" conform:"trim,slug"
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty" bson:"slug" valid:"required~Supplier slug is required" conform:"trim,slug"`
	// @inject_tag: bson:"logo_url" valid:"url~Supplier logo must be a valid URL" conform:"trim"
	LogoUrl string `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty" bson:"logo_url" valid:"url~Supplier logo must be a valid URL" conform:"trim"`
	// @inject_tag: bson:"home_page_url" valid:"required, url~Supplier home page must be a valid URL" conform:"trim"
	HomePageUrl string `protobuf:"bytes,5,opt,name=home_page_url,json=homePageUrl" json:"home_page_url,omitempty" bson:"home_page_url" valid:"required, url~Supplier home page must be a valid URL" conform:"trim"`
	// @inject_tag: bson:"currency" conform:"trim"
	Currency string `protobuf:"bytes,6,opt,name=currency" json:"currency,omitempty" bson:"currency" conform:"trim"`
	// @inject_tag: bson:"display_order"
	DisplayOrder int64 `protobuf:"varint,7,opt,name=display_order,json=displayOrder" json:"display_order,omitempty" bson:"display_order"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at"`
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

func (m *Supplier) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Supplier) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type SuppliersReq struct {
}

func (m *SuppliersReq) Reset()                    { *m = SuppliersReq{} }
func (m *SuppliersReq) String() string            { return proto.CompactTextString(m) }
func (*SuppliersReq) ProtoMessage()               {}
func (*SuppliersReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SuppliersRsp struct {
	Suppliers []*Supplier `protobuf:"bytes,1,rep,name=suppliers" json:"suppliers,omitempty"`
}

func (m *SuppliersRsp) Reset()                    { *m = SuppliersRsp{} }
func (m *SuppliersRsp) String() string            { return proto.CompactTextString(m) }
func (*SuppliersRsp) ProtoMessage()               {}
func (*SuppliersRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SuppliersRsp) GetSuppliers() []*Supplier {
	if m != nil {
		return m.Suppliers
	}
	return nil
}

func init() {
	proto.RegisterType((*UpsertSuppliersReq)(nil), "go.micro.srv.products.UpsertSuppliersReq")
	proto.RegisterType((*UpsertSuppliersRsp)(nil), "go.micro.srv.products.UpsertSuppliersRsp")
	proto.RegisterType((*Supplier)(nil), "go.micro.srv.products.Supplier")
	proto.RegisterType((*SuppliersReq)(nil), "go.micro.srv.products.SuppliersReq")
	proto.RegisterType((*SuppliersRsp)(nil), "go.micro.srv.products.SuppliersRsp")
}

func init() { proto.RegisterFile("proto/suppliers/supplier.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0xcf, 0xd2, 0x40,
	0x10, 0x4d, 0xcb, 0xe7, 0xf7, 0xb5, 0x53, 0x40, 0xb3, 0x89, 0xc9, 0xda, 0x83, 0x34, 0xe5, 0x52,
	0x2f, 0x4b, 0x82, 0x27, 0x0f, 0x1e, 0x30, 0x5c, 0x38, 0x18, 0x4d, 0x91, 0x8b, 0x97, 0xa6, 0xb4,
	0xeb, 0xda, 0xa4, 0x65, 0xd7, 0xdd, 0x2d, 0x09, 0xff, 0xcc, 0xdf, 0xe1, 0x2f, 0x32, 0xbb, 0xd0,
	0x42, 0x14, 0x83, 0x09, 0xb7, 0x99, 0x37, 0xef, 0xbd, 0x69, 0xf7, 0x0d, 0xbc, 0x16, 0x92, 0x6b,
	0x3e, 0x53, 0xad, 0x10, 0x75, 0x45, 0xa5, 0xea, 0x2b, 0x62, 0x07, 0xe8, 0x25, 0xe3, 0xa4, 0xa9,
	0x0a, 0xc9, 0x89, 0x92, 0x7b, 0x83, 0x95, 0x6d, 0xa1, 0x55, 0x38, 0x61, 0x9c, 0xb3, 0x9a, 0xce,
	0x2c, 0x69, 0xdb, 0x7e, 0x9b, 0xe9, 0xaa, 0xa1, 0x4a, 0xe7, 0x8d, 0x38, 0xea, 0xe2, 0x35, 0xa0,
	0x8d, 0x50, 0x54, 0xea, 0x75, 0xe7, 0x9c, 0xd2, 0x1f, 0xe8, 0x3d, 0xf8, 0xfd, 0x26, 0xec, 0x44,
	0x83, 0x24, 0x98, 0x4f, 0xc8, 0xd5, 0x0d, 0xa4, 0xd3, 0xa5, 0x67, 0xc5, 0x35, 0x53, 0x25, 0xee,
	0x35, 0xfd, 0xe9, 0x82, 0xd7, 0xe1, 0x68, 0x0c, 0xee, 0x6a, 0x89, 0x9d, 0xc8, 0x49, 0xfc, 0xd4,
	0x5d, 0x2d, 0x11, 0x82, 0x87, 0x5d, 0xde, 0x50, 0xec, 0x5a, 0xc4, 0xd6, 0x06, 0x53, 0x75, 0xcb,
	0xf0, 0xe0, 0x88, 0x99, 0x1a, 0xbd, 0x02, 0xaf, 0xe6, 0x8c, 0x67, 0xad, 0xac, 0xf1, 0x83, 0xc5,
	0x9f, 0x4c, 0xbf, 0x91, 0x35, 0x8a, 0x61, 0xf4, 0x9d, 0x37, 0x34, 0x13, 0x39, 0xa3, 0x76, 0xfe,
	0xcc, 0xce, 0x03, 0x03, 0x7e, 0xce, 0x19, 0x35, 0x9c, 0x10, 0xbc, 0xa2, 0x95, 0x92, 0xee, 0x8a,
	0x03, 0x7e, 0xb4, 0xe3, 0xbe, 0x47, 0x53, 0x18, 0x95, 0x95, 0x12, 0x75, 0x7e, 0xc8, 0xb8, 0x2c,
	0xa9, 0xc4, 0x4f, 0x91, 0x93, 0x0c, 0xd2, 0xe1, 0x09, 0xfc, 0x64, 0x30, 0xf4, 0x0e, 0xa0, 0x90,
	0x34, 0xd7, 0xb4, 0xcc, 0x72, 0x8d, 0xbd, 0xc8, 0x49, 0x82, 0x79, 0x48, 0x8e, 0x21, 0x91, 0x2e,
	0x24, 0xf2, 0xa5, 0x0b, 0x29, 0xf5, 0x4f, 0xec, 0x85, 0x36, 0xd2, 0x56, 0x94, 0x9d, 0xd4, 0xbf,
	0x2d, 0x3d, 0xb1, 0x17, 0x3a, 0x1e, 0xc3, 0xf0, 0x32, 0xde, 0xf8, 0xe3, 0x65, 0x7f, 0x77, 0x32,
	0xf3, 0x5f, 0x0e, 0xbc, 0xe8, 0xfd, 0xd6, 0x54, 0xee, 0xab, 0x82, 0x22, 0x06, 0xcf, 0xff, 0xb8,
	0x01, 0xf4, 0xe6, 0x1f, 0x9e, 0x7f, 0x1f, 0x60, 0xf8, 0xbf, 0x54, 0x25, 0xd0, 0x06, 0xfc, 0xf3,
	0x8a, 0xe9, 0x8d, 0xcf, 0xb6, 0xe6, 0xb7, 0x49, 0x4a, 0x7c, 0x08, 0xbe, 0x9e, 0xff, 0x70, 0xfb,
	0x68, 0xdf, 0xf7, 0xed, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x0e, 0x14, 0x4f, 0x86, 0x03,
	0x00, 0x00,
}
