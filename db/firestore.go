package db

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/fatih/structs"
	"github.com/jianhan/ms-bmp-products/firebase"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/supplier"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type fireBase struct {
	client *firestore.Client
}

func NewFirebaseDB(ctx context.Context) Database {
	return &fireBase{
		client: firebase.NewFirestoreClient(ctx),
	}
}

func (f *fireBase) UpsertSuppliers(ctx context.Context, suppliers []*psuppliers.Supplier) error {
	// setup batch
	batch := f.client.Batch()
	for _, s := range suppliers {
		// new record
		if s.ID == "" {
			s.ID = uuid.Must(uuid.NewV4()).String()
		}
		batch.Set(f.client.Collection("suppliers").Doc(s.ID), structs.Map(s), firestore.MergeAll)
	}

	// Commit the batch.
	_, err := batch.Commit(ctx)
	if err != nil {
		logrus.Errorf("error occur while batch set suppliers @UpsertSuppliers: %s", err.Error())
		return err
	}

	return nil
}
