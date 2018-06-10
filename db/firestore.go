package db

import (
	"context"
	"errors"

	"cloud.google.com/go/firestore"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"
	"github.com/jianhan/ms-bmp-products/firebase"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/supplier"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type firestoreDB struct {
	client *firestore.Client
	path   string
}

func NewFirestoreDB(ctx context.Context) Database {
	return &firestoreDB{
		client: firebase.NewFirestoreClient(ctx),
		path:   "suppliers",
	}
}

func (f *firestoreDB) UpsertSuppliers(ctx context.Context, suppliers []*psuppliers.Supplier) error {
	// get all first
	existingSuppliers, err := f.getAll(ctx)
	if err != nil {
		return err
	}
	spew.Dump(existingSuppliers)
	// setup batch
	batch := f.client.Batch()
	for _, s := range suppliers {
		for _, e := range existingSuppliers {
			if e.Name == s.Name && s.ID == "" {
				return errors.New("duplicate name")
			}
		}
		// auto fill IDs
		for _, s := range suppliers {
			if s.ID == "" {
				s.ID = uuid.Must(uuid.NewV4()).String()
			}
		}
		batch.Set(f.client.Collection(f.path).Doc(s.ID), structs.Map(s), firestore.MergeAll)
	}

	// Commit the batch.
	_, err = batch.Commit(ctx)
	if err != nil {
		logrus.Errorf("error occur while batch set suppliers @UpsertSuppliers: %s", err.Error())
		return err
	}

	return nil
}

func (f *firestoreDB) getAll(ctx context.Context) (suppliers []*psuppliers.Supplier, err error) {
	t, _ := f.client.Collection(f.path).Documents(ctx).GetAll()
	for _, v := range t {
		var s psuppliers.Supplier
		if err = v.DataTo(&s); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, &s)
	}

	return
}

func validateSuppliers(suppliers []*psuppliers.Supplier) error {
	return nil
}
