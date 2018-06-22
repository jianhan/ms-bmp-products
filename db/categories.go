package db

import (
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
)

type Categories interface {
	UpsertCategories(categories []*pcategories.Category) (int, int, error)
	Categories() (categories []*pcategories.Category, err error)
}
