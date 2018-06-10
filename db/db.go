package db

import (
	"context"

	psuppliers "github.com/jianhan/ms-bmp-products/proto/supplier"
)

type Database interface {
	UpsertSuppliers(ctx context.Context, suppliers []psuppliers.Supplier) error
}
