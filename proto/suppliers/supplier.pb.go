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
	Suppliers         []*Supplier `protobuf:"bytes,1,rep,name=suppliers" json:"suppliers,omitempty"`
	SyncElasticSearch bool        `protobuf:"varint,2,opt,name=sync_elastic_search,json=syncElasticSearch" json:"sync_elastic_search,omitempty"`
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

func (m *UpsertSuppliersReq) GetSyncElasticSearch() bool {
	if m != nil {
		return m.SyncElasticSearch
	}
	return false
}

type UpsertSuppliersRsp struct {
	Matched   int64       `protobuf:"varint,1,opt,name=matched" json:"matched,omitempty"`
	Modified  int64       `protobuf:"varint,2,opt,name=modified" json:"modified,omitempty"`
	Suppliers []*Supplier `protobuf:"bytes,3,rep,name=suppliers" json:"suppliers,omitempty"`
}

func (m *UpsertSuppliersRsp) Reset()                    { *m = UpsertSuppliersRsp{} }
func (m *UpsertSuppliersRsp) String() string            { return proto.CompactTextString(m) }
func (*UpsertSuppliersRsp) ProtoMessage()               {}
func (*UpsertSuppliersRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpsertSuppliersRsp) GetMatched() int64 {
	if m != nil {
		return m.Matched
	}
	return 0
}

func (m *UpsertSuppliersRsp) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

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
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0x57, 0x9a, 0xb1, 0x26, 0xaf, 0xdb, 0x00, 0x23, 0x24, 0x93, 0x03, 0xab, 0xb2, 0x4b, 0xb9,
	0xb8, 0x52, 0x39, 0x71, 0xe0, 0x30, 0x34, 0x0e, 0x3b, 0x20, 0x50, 0x4a, 0x2f, 0x5c, 0x22, 0xcf,
	0x79, 0x4b, 0x23, 0x25, 0xb5, 0xb1, 0x9d, 0x49, 0x3d, 0x73, 0xe2, 0x1b, 0xf1, 0x39, 0xf8, 0x44,
	0xc8, 0x6e, 0x93, 0x0e, 0x28, 0x2a, 0x88, 0xdb, 0x7b, 0xbf, 0x3f, 0xef, 0x3d, 0x5b, 0x3f, 0x78,
	0xae, 0xb4, 0xb4, 0x72, 0x6a, 0x5a, 0xa5, 0xea, 0x0a, 0xb5, 0xe9, 0x2b, 0xe6, 0x09, 0xf2, 0xb4,
	0x94, 0xac, 0xa9, 0x84, 0x96, 0xcc, 0xe8, 0x3b, 0x87, 0x15, 0xad, 0xb0, 0x26, 0x39, 0x2f, 0xa5,
	0x2c, 0x6b, 0x9c, 0x7a, 0xd1, 0x4d, 0x7b, 0x3b, 0xb5, 0x55, 0x83, 0xc6, 0xf2, 0x46, 0x6d, 0x7c,
	0xe9, 0x97, 0x00, 0xc8, 0x42, 0x19, 0xd4, 0x76, 0xde, 0x8d, 0xce, 0xf0, 0x33, 0x79, 0x0d, 0x71,
	0xbf, 0x8a, 0x06, 0xe3, 0x70, 0x32, 0x9a, 0x9d, 0xb3, 0xbd, 0x2b, 0x58, 0xe7, 0xcb, 0x76, 0x0e,
	0xc2, 0xe0, 0x89, 0x59, 0xaf, 0x44, 0x8e, 0x35, 0x37, 0xb6, 0x12, 0xb9, 0x41, 0xae, 0xc5, 0x92,
	0x0e, 0xc6, 0xc1, 0x24, 0xca, 0x1e, 0x3b, 0xea, 0xed, 0x86, 0x99, 0x7b, 0x22, 0xfd, 0xba, 0xe7,
	0x0a, 0xa3, 0x08, 0x85, 0x61, 0xc3, 0xad, 0x58, 0x62, 0x41, 0x83, 0x71, 0x30, 0x09, 0xb3, 0xae,
	0x25, 0x09, 0x44, 0x8d, 0x2c, 0xaa, 0xdb, 0x0a, 0x0b, 0x3f, 0x35, 0xcc, 0xfa, 0xfe, 0xe7, 0xdb,
	0xc3, 0x7f, 0xbd, 0x3d, 0xfd, 0x36, 0x80, 0xa8, 0xc3, 0xc9, 0x19, 0x0c, 0xae, 0xaf, 0xfc, 0xf2,
	0x38, 0x1b, 0x5c, 0x5f, 0x11, 0x02, 0x47, 0x2b, 0xde, 0xa0, 0xdf, 0x19, 0x67, 0xbe, 0x76, 0x98,
	0xa9, 0xdb, 0x92, 0x86, 0x1b, 0xcc, 0xd5, 0xe4, 0x19, 0x44, 0xb5, 0x2c, 0x65, 0xde, 0xea, 0x9a,
	0x1e, 0x79, 0x7c, 0xe8, 0xfa, 0x85, 0xae, 0x49, 0x0a, 0xa7, 0x4b, 0xd9, 0x60, 0xae, 0x78, 0x89,
	0x9e, 0x7f, 0xe0, 0xf9, 0x91, 0x03, 0x3f, 0xf0, 0x12, 0x9d, 0x26, 0x81, 0x48, 0xb4, 0x5a, 0xe3,
	0x4a, 0xac, 0xe9, 0xb1, 0xa7, 0xfb, 0x9e, 0x5c, 0xc0, 0x69, 0x51, 0x19, 0x55, 0xf3, 0x75, 0x2e,
	0x75, 0x81, 0x9a, 0x0e, 0xfd, 0xfb, 0x4f, 0xb6, 0xe0, 0x7b, 0x87, 0x91, 0x57, 0x00, 0x42, 0x23,
	0xb7, 0x58, 0xe4, 0xdc, 0xd2, 0x68, 0x1c, 0x4c, 0x46, 0xb3, 0x84, 0x6d, 0xc2, 0xc0, 0xba, 0x30,
	0xb0, 0x8f, 0x5d, 0x18, 0xb2, 0x78, 0xab, 0xbe, 0xb4, 0xce, 0xda, 0xaa, 0xa2, 0xb3, 0xc6, 0x87,
	0xad, 0x5b, 0xf5, 0xa5, 0x4d, 0xcf, 0xe0, 0xe4, 0x7e, 0x8a, 0xd2, 0x77, 0xf7, 0x7b, 0xa3, 0xfe,
	0x33, 0x55, 0xb3, 0xef, 0x01, 0x3c, 0xea, 0xe7, 0xcd, 0x51, 0xdf, 0x55, 0x02, 0x49, 0x09, 0x0f,
	0x7f, 0x49, 0x0e, 0x79, 0xf1, 0x87, 0x99, 0xbf, 0xe7, 0x3c, 0xf9, 0x5b, 0xa9, 0x51, 0x64, 0x01,
	0xf1, 0x6e, 0xc5, 0xc5, 0x81, 0xb3, 0xfd, 0xf0, 0xc3, 0x22, 0xa3, 0xde, 0x8c, 0x3e, 0xed, 0x5e,
	0x78, 0x73, 0xec, 0xff, 0xf7, 0xe5, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x8f, 0xa1, 0x5b,
	0xee, 0x03, 0x00, 0x00,
}
