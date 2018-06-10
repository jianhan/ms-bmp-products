package db

import (
	"context"
	"fmt"

	"strings"

	"cloud.google.com/go/firestore"
	"github.com/asaskevich/govalidator"
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
	// validation
	if err := validateSuppliers(suppliers); err != nil {
		return err
	}

	// get all first
	existingSuppliers, err := f.getAll(ctx)
	if err != nil {
		return err
	}

	// setup batch
	batch := f.client.Batch()
	for _, s := range suppliers {
		for _, e := range existingSuppliers {
			// start validation for unique constraints
			if e.Name == s.Name && s.ID == "" {
				// duplication for inserting, checking name
				return fmt.Errorf("error trying to insert new supplier %v with duplicate name %s", s, s.Name)
			}

			if s.ID != "" && s.ID != e.ID && e.Name == s.Name {
				// duplication for updating, checking name
				return fmt.Errorf("error trying to update supplier %v with duplicate name %s", s, s.Name)
			}

			if e.HomePageUrl == s.HomePageUrl && s.ID == "" {
				// duplication for inserting, checking homepage url
				return fmt.Errorf("error trying to insert new supplier %v with duplicate homepage url %s", s, s.HomePageUrl)
			}

			if s.ID != "" && s.ID != e.ID && e.HomePageUrl == s.HomePageUrl {
				// duplication for updating, checking name
				return fmt.Errorf("error trying to update supplier %v with duplicate homepage url %s", s, s.HomePageUrl)
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

func checkEmptyTrimmed(fieldName, str, id string) (string, error) {
	trimmed := strings.Trim(str, " ")
	if trimmed == "" {
		return "", fmt.Errorf("%s can not be empty for supplier %s", fieldName, id)
	}
	return trimmed, nil
}

func validateSuppliers(suppliers []*psuppliers.Supplier) (err error) {
	for _, v := range suppliers {
		// UUID checking
		if v.ID != "" && !govalidator.IsUUID(v.ID) {
			return fmt.Errorf("invalid UUID %s", v.ID)
		}

		// name checking
		if v.Name, err = checkEmptyTrimmed("name", v.Name, v.ID); err != nil {
			return err
		}

		// home page url checking
		if v.HomePageUrl, err = checkEmptyTrimmed("homepage url", v.HomePageUrl, v.ID); err != nil {
			return err
		}
		if !govalidator.IsURL(v.HomePageUrl) {
			return fmt.Errorf("invalid url %s for supplier %s", v.HomePageUrl, v.ID)
		}

		// currency checking
		if v.Currency, err = checkEmptyTrimmed("currency", v.Currency, v.ID); err != nil {
			return err
		}
	}

	return nil
}
