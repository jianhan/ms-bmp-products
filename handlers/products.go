package handlers

import (
	"context"

	"github.com/jianhan/ms-bmp-products/db"
	pproducts "github.com/jianhan/ms-bmp-products/proto/products"
	"github.com/nats-io/go-nats-streaming"
)

const (
	TopicSyncProductsToElasticSearch = "products:sync-to-elastic-search"
)

type Products struct {
	categoriesDB db.Categories
	db           db.Products
	stanConn     stan.Conn
}

func NewProductsHandler(categoriesDB db.Categories, db db.Products, stanConn stan.Conn) *Products {
	return &Products{categoriesDB: categoriesDB, db: db, stanConn: stanConn}
}

func (h *Products) UpsertProducts(ctx context.Context, req *pproducts.UpsertProductsReq, rsp *pproducts.UpsertProductsRsp) (err error) {
	//	get all categories
	categories, cErr := h.categoriesDB.Categories()
	if cErr != nil {
		return cErr
	}

	if rsp.Matched, rsp.Modified, err = h.db.UpsertProducts(categories, req.Products); err != nil {
		return err
	}

	// publish event
	if err := h.stanConn.Publish(TopicSyncProductsToElasticSearch, nil); err != nil {
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
