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

	// publish message
	rspBytes, err := proto.Marshal(rsp)
	if err != nil {
		return err
	}
	if err := s.stanConn.Publish(TopicSuppliersUpserted, rspBytes); err != nil {
		return err
	}

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
