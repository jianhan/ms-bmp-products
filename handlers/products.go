package handlers

import (
	"context"

	"github.com/jianhan/ms-bmp-products/db"
	pproducts "github.com/jianhan/ms-bmp-products/proto/product"
)

type Products struct {
	db db.Products
}

func NewProductsHandler(db db.Products) *Products {
	return &Products{db: db}
}

func (s *Products) UpsertProducts(ctx context.Context, req *pproducts.UpsertProductsReq, rsp *pproducts.UpsertProductsRsp) error {
	if err := s.db.UpsertProducts(ctx, req.Products); err != nil {
		return err
	}

	return nil
}
