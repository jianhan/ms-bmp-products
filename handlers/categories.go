package handlers

import (
	"context"

	"github.com/gogo/protobuf/proto"
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

func (s *Categories) UpsertCategories(ctx context.Context, req *pcategories.UpsertCategoriesReq, rsp *pcategories.UpsertCategoriesRsp) error {
	if err := s.db.UpsertCategories(ctx, req.Categories); err != nil {
		return err
	}

	// get all products and construct response
	categories, err := s.db.GetAllCategories(ctx)
	if err != nil {
		return err
	}
	rsp.Categories = categories

	// publish message
	rspBytes, err := proto.Marshal(rsp)
	if err != nil {
		return err
	}
	if err := s.stanConn.Publish(TopicCategoriesUpserted, rspBytes); err != nil {
		return err
	}

	return nil
}
