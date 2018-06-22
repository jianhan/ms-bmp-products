package db

import (
	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
)

type Suppliers interface {
	UpsertSuppliers(suppliers []*psuppliers.Supplier) error
	Suppliers() (suppliers []*psuppliers.Supplier, err error)
}
