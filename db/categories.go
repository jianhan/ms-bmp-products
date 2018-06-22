package db

import (
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
)

type Categories interface {
	UpsertCategories(categories []*pcategories.Category) (int64, int64, error)
	Categories() (categories []*pcategories.Category, err error)
}
