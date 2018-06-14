package handlers

import (
	"context"

	"github.com/jianhan/ms-bmp-products/db"
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
)

type Categories struct {
	db db.Categories
}

func NewCategoriesHandler(db db.Categories) *Categories {
	return &Categories{db: db}
}

func (s *Categories) UpsertCategories(ctx context.Context, req *pcategories.UpsertCategoriesReq, rsp *pcategories.UpsertCategoriesRsp) error {
	if err := s.db.UpsertCategories(ctx, req.Categories); err != nil {
		return err
	}

	return nil
}
