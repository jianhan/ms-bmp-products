package handlers

import (
	"context"

	"github.com/jianhan/ms-bmp-products/db"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
)

type Suppliers struct {
	db db.Suppliers
}

func NewSuppliersHandler(db db.Suppliers) *Suppliers {
	return &Suppliers{db: db}
}

func (s *Suppliers) UpsertSuppliers(ctx context.Context, req *psuppliers.UpsertSuppliersReq, rsp *psuppliers.UpsertSuppliersRsp) error {
	if err := s.db.UpsertSuppliers(ctx, req.Suppliers); err != nil {
		return err
	}

	return nil
}
