package handlers

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/jianhan/ms-bmp-products/db"
	pproducts "github.com/jianhan/ms-bmp-products/proto/products"
	"github.com/nats-io/go-nats-streaming"
)

const (
	TopicSyncProductsToElasticSearch = "products:sync-to-elastic-search"
)

type Products struct {
	db       db.Products
	stanConn stan.Conn
}

func NewProductsHandler(db db.Products, stanConn stan.Conn) *Products {
	return &Products{db: db, stanConn: stanConn}
}

func (h *Products) UpsertProducts(ctx context.Context, req *pproducts.UpsertProductsReq, rsp *pproducts.UpsertProductsRsp) (err error) {
	if rsp.Matched, rsp.Modified, err = h.db.UpsertProducts(req.Products); err != nil {
		return err
	}

	// get all products and construct response
	products, err := h.db.Products()
	if err != nil {
		return err
	}
	rsp.Products = products

	// publish message
	rspBytes, err := proto.Marshal(rsp)
	if err != nil {
		return err
	}
	if err := h.stanConn.Publish(TopicSyncProductsToElasticSearch, rspBytes); err != nil {
		return err
	}

	return nil
}

func (h *Products) Products(ctx context.Context, req *pproducts.ProductsReq, rsp *pproducts.ProductsRsp) (err error) {
	// get all products and construct response
	if rsp.Products, err = h.db.Products(); err != nil {
		return
	}

	return nil
}
