package db

import (
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
	pproducts "github.com/jianhan/ms-bmp-products/proto/products"
)

type Products interface {
	UpsertProducts(categories []*pcategories.Category, products []*pproducts.Product) (int64, int64, error)
	Products() (products []*pproducts.Product, err error)
}
