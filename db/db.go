package db

import (
	psuppliers "github.com/jianhan/ms-bmp-products/proto/supplier"
)

type Database interface {
	UpsertSuppliers(suppliers []psuppliers.Supplier) error
}
