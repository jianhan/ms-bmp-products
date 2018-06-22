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
		return err
	}

	//// get all products and construct response
	//categories, err := h.db.GetAllCategories(ctx)
	//if err != nil {
	//	return err
	//}
	//rsp.Categories = categories
	//
	//// publish message
	//rspBytes, err := proto.Marshal(rsp)
	//if err != nil {
	//	return err
	//}
	//if err := h.stanConn.Publish(TopicCategoriesUpserted, rspBytes); err != nil {
	//	return err
	//}

	return nil
}

func (h *Categories) Categories(ctx context.Context, req *pcategories.CategoriesReq, rsp *pcategories.CategoriesRsp) error {
	// get all products and construct response
	categories, err := h.db.Categories()
	if err != nil {
		return err
	}
	rsp.Categories = categories

	return nil
}
