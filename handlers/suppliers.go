package handlers

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/jianhan/ms-bmp-products/db"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
	"github.com/nats-io/go-nats-streaming"
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

func (h *Suppliers) UpsertSuppliers(ctx context.Context, req *psuppliers.UpsertSuppliersReq, rsp *psuppliers.UpsertSuppliersRsp) error {
	// bulk upserts
	if err := h.db.UpsertSuppliers(ctx, req.Suppliers); err != nil {
		return err
	}

	// get all suppliers and construct response
	suppliers, err := h.db.GetAllSuppliers(ctx)
	if err != nil {
		return err
	}
	rsp.Suppliers = suppliers

	// publish message
	rspBytes, err := proto.Marshal(rsp)
	if err != nil {
		return err
	}
	if err := h.stanConn.Publish(TopicSuppliersUpserted, rspBytes); err != nil {
		return err
	}

	return nil
}
