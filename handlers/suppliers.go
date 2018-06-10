package handlers

import (
	"context"

	"github.com/jianhan/ms-bmp-products/db"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/supplier"
)

type Suppliers struct {
	db db.Database
}

func (s *Suppliers) UpsertSuppliers(ctx context.Context, req *psuppliers.UpsertSuppliersReq, rsp *psuppliers.UpsertSuppliersRsp) error {
	if err := s.db.UpsertSuppliers(ctx, req.Suppliers); err != nil {
		return err
	}

	return nil
}
