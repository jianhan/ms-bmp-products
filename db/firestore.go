package db

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jianhan/ms-bmp-products/firebase"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/supplier"
)

type fireBase struct {
	client *firestore.Client
}

func NewFirebaseDB(ctx context.Context) Database {
	return &fireBase{
		client: firebase.NewFirestoreClient(ctx),
	}
}

func (f *fireBase) UpsertSuppliers(suppliers []psuppliers.Supplier) error {
	batch := f.client.Batch()
}
