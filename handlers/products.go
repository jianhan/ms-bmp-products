package handlers

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/jianhan/ms-bmp-products/db"
	pproducts "github.com/jianhan/ms-bmp-products/proto/products"
	"github.com/nats-io/go-nats-streaming"
)

const (
	TopicProductsUpserted = "products:upserted"
)

type Products struct {
	db       db.Products
	stanConn stan.Conn
}

func NewProductsHandler(db db.Products, stanConn stan.Conn) *Products {
	return &Products{db: db, stanConn: stanConn}
}

func (h *Products) UpsertProducts(ctx context.Context, req *pproducts.UpsertProductsReq, rsp *pproducts.UpsertProductsRsp) error {
	if err := h.db.UpsertProducts(ctx, req.Products); err != nil {
		return err
	}

	// get all products and construct response
	products, err := h.db.GetAllProducts(ctx)
	if err != nil {
		return err
	}
	rsp.Products = products

	// publish message
	rspBytes, err := proto.Marshal(rsp)
	if err != nil {
		return err
	}
	if err := h.stanConn.Publish(TopicProductsUpserted, rspBytes); err != nil {
		return err
	}

	return nil
}
