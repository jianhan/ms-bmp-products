package handlers

import (
	"context"

	"github.com/jianhan/ms-bmp-products/db"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
	"github.com/micro/go-micro/broker"
)

const (
	TopicSuppliersUpserted = "suppliers.upserted"
)

type Suppliers struct {
	db     db.Suppliers
	PubSub broker.Broker
}

func NewSuppliersHandler(db db.Suppliers) *Suppliers {
	return &Suppliers{db: db}
}

func (s *Suppliers) UpsertSuppliers(ctx context.Context, req *psuppliers.UpsertSuppliersReq, rsp *psuppliers.UpsertSuppliersRsp) error {
	// bulk upserts
	if err := s.db.UpsertSuppliers(ctx, req.Suppliers); err != nil {
		return err
	}

	// get all suppliers
	suppliers, err := s.db.GetAllSuppliers(ctx)
	if err != nil {
		return err
	}

	rsp.Suppliers = suppliers

	return nil
}

//func (s *Suppliers) publishUpsertSuppliers(ctx context.Context) error {
//	suppliers, err := s.db.GetAllSuppliers(ctx)
//	if err != nil {
//		return err
//	}
//
//	// Publish message to broker
//	if err := s.PubSub.Publish(TopicSuppliersUpserted, suppliers).Publish(topicSuppliersUpserted, msg); err != nil {
//		log.Printf("[pub] failed: %v", err)
//	}
//}
