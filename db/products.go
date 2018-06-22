package db

import (
	pproducts "github.com/jianhan/ms-bmp-products/proto/products"
)

type Products interface {
	UpsertProducts(products []*pproducts.Product) error
	Products() (products []*pproducts.Product, err error)
}
