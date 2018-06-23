package handlers

import (
	"context"

	"github.com/jianhan/ms-bmp-products/db"
	"github.com/nats-io/go-nats-streaming"

	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
)

const (
	TopicSuppliersUpserted = "suppliers:upserted"
)

type Suppliers struct {
	db       db.Suppliers
	stanConn stan.Conn
}

func NewSuppliersHandler(db db.Suppliers, stanConn stan.Conn) *Suppliers {
	return &Suppliers{db: db, stanConn: stanConn}
}

func (h *Suppliers) UpsertSuppliers(req *psuppliers.UpsertSuppliersReq, rsp *psuppliers.UpsertSuppliersRsp) (err error) {
	// bulk upserts
	if rsp.Matched, rsp.Modified, err = h.db.UpsertSuppliers(req.Suppliers); err != nil {
		return
	}

	return nil
}

func (h *Suppliers) Suppliers(ctx context.Context, req *psuppliers.SuppliersReq, rsp *psuppliers.SuppliersRsp) (err error) {
	// get all suppliers and construct response
	if rsp.Suppliers, err = h.db.Suppliers(); err != nil {
		return
	}

	return
}
