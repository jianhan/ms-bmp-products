package db

import (
	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
)

type Suppliers interface {
	UpsertSuppliers(suppliers []*psuppliers.Supplier) (int64, int64, error)
	Suppliers() (suppliers []*psuppliers.Supplier, err error)
}
