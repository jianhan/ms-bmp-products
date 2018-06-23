package db

import (
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
)

type Categories interface {
	UpsertCategories(suppliers []*psuppliers.Supplier, categories []*pcategories.Category) (int64, int64, error)
	Categories() (categories []*pcategories.Category, err error)
}
