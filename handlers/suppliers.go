package handlers

import (
	"github.com/jianhan/ms-bmp-products/db"
	"github.com/nats-io/go-nats-streaming"

	"context"

	"github.com/gogo/protobuf/proto"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
)

const (
	TopicSyncSuppliersToElasticSearch = "suppliers:sync-to-elastic-search"
)

type Suppliers struct {
	db       db.Suppliers
	stanConn stan.Conn
}

func NewSuppliersHandler(db db.Suppliers, stanConn stan.Conn) *Suppliers {
	return &Suppliers{db: db, stanConn: stanConn}
}

func (h *Suppliers) UpsertSuppliers(ctx context.Context, req *psuppliers.UpsertSuppliersReq, rsp *psuppliers.UpsertSuppliersRsp) (err error) {
	// bulk upserts
	if rsp.Matched, rsp.Modified, err = h.db.UpsertSuppliers(req.Suppliers); err != nil {
		return
	}

	// sync required, publish message
	if rsp.Suppliers, err = h.db.Suppliers(); err != nil {
		return
	}
	rspBytes, rErr := proto.Marshal(rsp)
	if rErr != nil {
		return rErr
	}
	h.stanConn.Publish(TopicSyncSuppliersToElasticSearch, rspBytes)

	return
}

func (h *Suppliers) Suppliers(ctx context.Context, req *psuppliers.SuppliersReq, rsp *psuppliers.SuppliersRsp) (err error) {
	// get all suppliers and construct response
	if rsp.Suppliers, err = h.db.Suppliers(); err != nil {
		return
	}

	return
}
