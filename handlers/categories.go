package handlers

import (
	"context"

	"github.com/jianhan/ms-bmp-products/db"
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
	"github.com/nats-io/go-nats-streaming"
)

const (
	TopicCategoriesUpserted = "categories:upserted"
)

type Categories struct {
	db       db.Categories
	stanConn stan.Conn
}

func NewCategoriesHandler(db db.Categories, stanConn stan.Conn) *Categories {
	return &Categories{db: db, stanConn: stanConn}
}

func (h *Categories) UpsertCategories(ctx context.Context, req *pcategories.UpsertCategoriesReq, rsp *pcategories.UpsertCategoriesRsp) (err error) {
	if rsp.Matched, rsp.Modified, err = h.db.UpsertCategories(req.Categories); err != nil {
		return
	}

	return
}

func (h *Categories) Categories(ctx context.Context, req *pcategories.CategoriesReq, rsp *pcategories.CategoriesRsp) (err error) {
	// get all products and construct response
	if rsp.Categories, err = h.db.Categories(); err != nil {
		return
	}

	return nil
}
