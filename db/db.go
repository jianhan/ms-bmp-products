package db

import (
	"context"

	pproducts "github.com/jianhan/ms-bmp-products/proto/product"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/supplier"
)

type Database interface {
	UpsertSuppliers(ctx context.Context, suppliers []*psuppliers.Supplier) error
	UpsertProducts(ctx context.Context, products []*pproducts.Product) error
	Close() error
}
