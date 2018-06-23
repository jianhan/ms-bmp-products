package handlers

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/jianhan/ms-bmp-products/db"
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
	"github.com/nats-io/go-nats-streaming"
)

const (
	TopicSyncCategoriesToElasticSearch = "categories:sync-to-elastic-search"
)

type Categories struct {
	suppliersDB db.Suppliers
	db          db.Categories
	stanConn    stan.Conn
}

func NewCategoriesHandler(suppliersDB db.Suppliers, db db.Categories, stanConn stan.Conn) *Categories {
	return &Categories{suppliersDB: suppliersDB, db: db, stanConn: stanConn}
}

func (h *Categories) UpsertCategories(ctx context.Context, req *pcategories.UpsertCategoriesReq, rsp *pcategories.UpsertCategoriesRsp) (err error) {
	// get all suppliers first
	suppliers, sErr := h.suppliersDB.Suppliers()
	if sErr != nil {
		return sErr
	}

	if rsp.Matched, rsp.Modified, err = h.db.UpsertCategories(suppliers, req.Categories); err != nil {
		return
	}

	// sync required, publish message
	if rsp.Categories, err = h.db.Categories(); err != nil {
		return
	}
	rspBytes, rErr := proto.Marshal(rsp)
	if rErr != nil {
		return rErr
	}
	h.stanConn.Publish(TopicSyncCategoriesToElasticSearch, rspBytes)

	return
}

func (h *Categories) Categories(ctx context.Context, req *pcategories.CategoriesReq, rsp *pcategories.CategoriesRsp) (err error) {
	// get all products and construct response
	if rsp.Categories, err = h.db.Categories(); err != nil {
		return
	}

	return nil
}
