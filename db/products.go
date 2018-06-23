package db

import (
	pproducts "github.com/jianhan/ms-bmp-products/proto/products"
)

type Products interface {
	UpsertProducts(products []*pproducts.Product) (int64, int64, error)
	Products() (products []*pproducts.Product, err error)
}
