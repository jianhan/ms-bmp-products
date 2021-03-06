// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/products/product.proto

/*
Package products is a generated protocol buffer package.

It is generated from these files:
	proto/products/product.proto

It has these top-level messages:
	UpsertProductsReq
	UpsertProductsRsp
	ProductsReq
	ProductsRsp
	Product
*/
package products

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

type UpsertProductsReq struct {
	Products []*Product `protobuf:"bytes,1,rep,name=products" json:"products,omitempty"`
}

func (m *UpsertProductsReq) Reset()                    { *m = UpsertProductsReq{} }
func (m *UpsertProductsReq) String() string            { return proto.CompactTextString(m) }
func (*UpsertProductsReq) ProtoMessage()               {}
func (*UpsertProductsReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpsertProductsReq) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type UpsertProductsRsp struct {
	Matched  int64      `protobuf:"varint,1,opt,name=matched" json:"matched,omitempty"`
	Modified int64      `protobuf:"varint,2,opt,name=modified" json:"modified,omitempty"`
	Products []*Product `protobuf:"bytes,3,rep,name=products" json:"products,omitempty"`
}

func (m *UpsertProductsRsp) Reset()                    { *m = UpsertProductsRsp{} }
func (m *UpsertProductsRsp) String() string            { return proto.CompactTextString(m) }
func (*UpsertProductsRsp) ProtoMessage()               {}
func (*UpsertProductsRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpsertProductsRsp) GetMatched() int64 {
	if m != nil {
		return m.Matched
	}
	return 0
}

func (m *UpsertProductsRsp) GetModified() int64 {
	if m != nil {
		return m.Modified
	}
	return 0
}

func (m *UpsertProductsRsp) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type ProductsReq struct {
}

func (m *ProductsReq) Reset()                    { *m = ProductsReq{} }
func (m *ProductsReq) String() string            { return proto.CompactTextString(m) }
func (*ProductsReq) ProtoMessage()               {}
func (*ProductsReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ProductsRsp struct {
	Products []*Product `protobuf:"bytes,1,rep,name=products" json:"products,omitempty"`
}

func (m *ProductsRsp) Reset()                    { *m = ProductsRsp{} }
func (m *ProductsRsp) String() string            { return proto.CompactTextString(m) }
func (*ProductsRsp) ProtoMessage()               {}
func (*ProductsRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ProductsRsp) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type Product struct {
	// @inject_tag: bson:"_id" valid:"uuidv4~Product ID must be a valid UUIDv4" conform:"trim"
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty" bson:"_id" valid:"uuidv4~Product ID must be a valid UUIDv4" conform:"trim"`
	// @inject_tag: bson:"name" valid:"required~Product name is required" conform:"trim"
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty" bson:"name" valid:"required~Product name is required" conform:"trim"`
	// @inject_tag: bson:"slug" valid:"required~Product slug is required" conform:"trim,slug"
	Slug string `protobuf:"bytes,3,opt,name=slug" json:"slug,omitempty" bson:"slug" valid:"required~Product slug is required" conform:"trim,slug"`
	// @inject_tag: bson:"price" valid:"required~Product price is required"
	Price float32 `protobuf:"fixed32,4,opt,name=price" json:"price,omitempty" bson:"price" valid:"required~Product price is required"`
	// @inject_tag: bson:"currency" valid:"required~Product currency is required"
	Currency string `protobuf:"bytes,5,opt,name=currency" json:"currency,omitempty" bson:"currency" valid:"required~Product currency is required"`
	// @inject_tag: bson:"image_url" valid:"url~Product image must be a valid URL format" conform:"trim"
	ImageUrl string `protobuf:"bytes,6,opt,name=image_url,json=imageUrl" json:"image_url,omitempty" bson:"image_url" valid:"url~Product image must be a valid URL format" conform:"trim"`
	// @inject_tag: bson:"category_url" valid:"required~Product categories url is required,url~Product category url must be a valid URL format" conform:"trim"
	CategoryUrl string `protobuf:"bytes,7,opt,name=category_url,json=categoryUrl" json:"category_url,omitempty" bson:"category_url" valid:"required~Product categories url is required,url~Product category url must be a valid URL format" conform:"trim"`
	// @inject_tag: bson:"url" valid:"required~Product url is required,url~Product url must be a valid URL format" conform:"trim"
	Url string `protobuf:"bytes,8,opt,name=url" json:"url,omitempty" bson:"url" valid:"required~Product url is required,url~Product url must be a valid URL format" conform:"trim"`
	// @inject_tag: bson:"rating"
	Rating float32 `protobuf:"fixed32,9,opt,name=rating" json:"rating,omitempty" bson:"rating"`
	// @inject_tag: bson:"brand" conform:"trim"
	Brand string `protobuf:"bytes,10,opt,name=brand" json:"brand,omitempty" bson:"brand" conform:"trim"`
	// @inject_tag: bson:"hidden"
	Hidden bool `protobuf:"varint,11,opt,name=hidden" json:"hidden,omitempty" bson:"hidden"`
	// @inject_tag: bson:"category_id" valid:"required~Category ID is required for product"
	CategoryId string `protobuf:"bytes,12,opt,name=category_id,json=categoryId" json:"category_id,omitempty" bson:"category_id" valid:"required~Category ID is required for product"`
	// @inject_tag: bson:"created_at"
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,13,opt,name=created_at,json=createdAt" json:"created_at,omitempty" bson:"created_at"`
	// @inject_tag: bson:"updated_at"
	UpdatedAt *google_protobuf.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty" bson:"updated_at"`
}

func (m *Product) Reset()                    { *m = Product{} }
func (m *Product) String() string            { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()               {}
func (*Product) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Product) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Product) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Product) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Product) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *Product) GetCategoryUrl() string {
	if m != nil {
		return m.CategoryUrl
	}
	return ""
}

func (m *Product) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Product) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Product) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Product) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *Product) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

func (m *Product) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Product) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*UpsertProductsReq)(nil), "go.micro.srv.products.UpsertProductsReq")
	proto.RegisterType((*UpsertProductsRsp)(nil), "go.micro.srv.products.UpsertProductsRsp")
	proto.RegisterType((*ProductsReq)(nil), "go.micro.srv.products.ProductsReq")
	proto.RegisterType((*ProductsRsp)(nil), "go.micro.srv.products.ProductsRsp")
	proto.RegisterType((*Product)(nil), "go.micro.srv.products.Product")
}

func init() { proto.RegisterFile("proto/products/product.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0x9a, 0xad, 0x4d, 0x5e, 0xb6, 0x02, 0x16, 0x20, 0xab, 0x20, 0x16, 0x7a, 0xca, 0x29,
	0x95, 0xca, 0x09, 0x6e, 0x43, 0xbb, 0xf4, 0x04, 0x32, 0xec, 0xc2, 0xa5, 0x72, 0x63, 0x2f, 0xb3,
	0xd4, 0xc4, 0xc6, 0x76, 0x26, 0xed, 0x07, 0xf0, 0x47, 0xf8, 0x3b, 0xfc, 0x29, 0x64, 0x27, 0x8e,
	0x36, 0x0d, 0x28, 0x88, 0xdb, 0xfb, 0xbe, 0xf7, 0x7d, 0xaf, 0xef, 0xc5, 0x5f, 0xe1, 0xa5, 0xd2,
	0xd2, 0xca, 0x95, 0xd2, 0x92, 0x75, 0x95, 0x35, 0xa1, 0x28, 0x3d, 0x8d, 0x9e, 0xd5, 0xb2, 0x6c,
	0x44, 0xa5, 0x65, 0x69, 0xf4, 0x4d, 0x19, 0x44, 0x8b, 0xb3, 0x5a, 0xca, 0x7a, 0xcf, 0x57, 0x5e,
	0xb4, 0xeb, 0xae, 0x56, 0x56, 0x34, 0xdc, 0x58, 0xda, 0xa8, 0xde, 0xb7, 0xfc, 0x00, 0x4f, 0x2e,
	0x95, 0xe1, 0xda, 0x7e, 0x1c, 0x2c, 0x84, 0x7f, 0x45, 0xef, 0x20, 0x09, 0x13, 0x70, 0x94, 0xc7,
	0x45, 0xb6, 0x7e, 0x55, 0xfe, 0x72, 0x7e, 0x39, 0xb8, 0xc8, 0xa8, 0x5f, 0x7e, 0x8b, 0x1e, 0x4c,
	0x34, 0x0a, 0x61, 0x98, 0x35, 0xd4, 0x56, 0xd7, 0x9c, 0xe1, 0x28, 0x8f, 0x8a, 0x98, 0x04, 0x88,
	0x16, 0x90, 0x34, 0x92, 0x89, 0x2b, 0xc1, 0x19, 0x9e, 0xf8, 0xd6, 0x88, 0xef, 0xed, 0x11, 0xff,
	0xe3, 0x1e, 0xa7, 0x90, 0xdd, 0x39, 0x69, 0xb9, 0xb9, 0x03, 0x8d, 0xfa, 0xaf, 0x0b, 0xbf, 0xc7,
	0x30, 0x1b, 0x58, 0x34, 0x87, 0xc9, 0xe6, 0xc2, 0x9f, 0x94, 0x92, 0xc9, 0xe6, 0x02, 0x21, 0x38,
	0x6a, 0x69, 0xc3, 0xfd, 0x25, 0x29, 0xf1, 0xb5, 0xe3, 0xcc, 0xbe, 0xab, 0x71, 0xdc, 0x73, 0xae,
	0x46, 0x4f, 0xe1, 0x58, 0x69, 0x51, 0x71, 0x7c, 0x94, 0x47, 0xc5, 0x84, 0xf4, 0xc0, 0x7d, 0x8b,
	0xaa, 0xd3, 0x9a, 0xb7, 0xd5, 0x2d, 0x3e, 0xf6, 0xea, 0x11, 0xa3, 0x17, 0x90, 0x8a, 0x86, 0xd6,
	0x7c, 0xdb, 0xe9, 0x3d, 0x9e, 0xf6, 0x4d, 0x4f, 0x5c, 0xea, 0x3d, 0x7a, 0x0d, 0x27, 0x15, 0xb5,
	0xbc, 0x96, 0xfa, 0xd6, 0xf7, 0x67, 0xbe, 0x9f, 0x05, 0xce, 0x49, 0x1e, 0x43, 0xec, 0x3a, 0x89,
	0xef, 0xb8, 0x12, 0x3d, 0x87, 0xa9, 0xa6, 0x56, 0xb4, 0x35, 0x4e, 0xfd, 0x12, 0x03, 0x72, 0xbb,
	0xed, 0x34, 0x6d, 0x19, 0x06, 0xaf, 0xed, 0x81, 0x53, 0x5f, 0x0b, 0xc6, 0x78, 0x8b, 0xb3, 0x3c,
	0x2a, 0x12, 0x32, 0x20, 0x74, 0x06, 0xe3, 0xcf, 0x6c, 0x05, 0xc3, 0x27, 0xde, 0x03, 0x81, 0xda,
	0x30, 0xf4, 0x16, 0xa0, 0xd2, 0x9c, 0x5a, 0xce, 0xb6, 0xd4, 0xe2, 0xd3, 0x3c, 0x2a, 0xb2, 0xf5,
	0xa2, 0xec, 0x73, 0x59, 0x86, 0x5c, 0x96, 0x9f, 0x43, 0x2e, 0x49, 0x3a, 0xa8, 0xcf, 0xad, 0xb3,
	0x76, 0x8a, 0x05, 0xeb, 0xfc, 0xb0, 0x75, 0x50, 0x9f, 0xdb, 0xf5, 0x8f, 0x08, 0x1e, 0x85, 0x07,
	0xff, 0xc4, 0xf5, 0x8d, 0xfb, 0xbc, 0x0c, 0xe6, 0xf7, 0x93, 0x89, 0x8a, 0xdf, 0x3c, 0xfa, 0x83,
	0xbf, 0xc4, 0xe2, 0x2f, 0x95, 0x46, 0x21, 0x02, 0xc9, 0x38, 0x7f, 0xf9, 0xe7, 0x50, 0xf9, 0xc9,
	0x07, 0x35, 0x46, 0xbd, 0x87, 0x2f, 0x63, 0xfc, 0x76, 0x53, 0x7f, 0xf8, 0x9b, 0x9f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0xb0, 0xab, 0xa1, 0x10, 0x04, 0x00, 0x00,
}
