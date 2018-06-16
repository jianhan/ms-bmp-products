package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/asaskevich/govalidator"
	"github.com/fatih/structs"
	"github.com/gosimple/slug"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
	"github.com/leebenson/conform"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Suppliers interface {
	UpsertSuppliers(ctx context.Context, suppliers []*psuppliers.Supplier) error
	GetAllSuppliers(ctx context.Context) (suppliers []*psuppliers.Supplier, err error)
}

type suppliers struct {
	Base
}

func NewSuppliers(path string, client *firestore.Client) Suppliers {
	base := Base{
		path:   path,
		client: client,
	}

	return &suppliers{base}
}

func (d *suppliers) UpsertSuppliers(ctx context.Context, suppliers []*psuppliers.Supplier) error {
	// validation
	if err := validateSuppliers(suppliers); err != nil {
		return err
	}

	// get all first
	existingSuppliers, err := d.GetAllSuppliers(ctx)
	if err != nil {
		return err
	}

	// setup batch
	batch := d.client.Batch()
	now := time.Now()
	for _, s := range suppliers {
		for _, e := range existingSuppliers {
			// start validation for unique constraints
			if e.Name == s.Name && s.ID == "" {
				// duplication for inserting, checking name
				return fmt.Errorf("error trying to insert new suppliers %v with duplicate name %s", s, s.Name)
			}

			if s.ID != "" && s.ID != e.ID && e.Name == s.Name {
				// duplication for updating, checking name
				return fmt.Errorf("error trying to update suppliers %v with duplicate name %s", s, s.Name)
			}

			if e.Slug == s.Slug && s.ID == "" {
				// duplication for inserting, checking slug
				return fmt.Errorf("error trying to insert new suppliers %v with duplicate slug %s", s, s.Slug)
			}

			if s.ID != "" && s.ID != e.ID && e.Slug == s.Slug {
				// duplication for updating, checking slug
				return fmt.Errorf("error trying to update suppliers %v with duplicate slug %s", s, s.Slug)
			}

			if e.HomePageUrl == s.HomePageUrl && s.ID == "" {
				// duplication for inserting, checking homepage url
				return fmt.Errorf("error trying to insert new suppliers %v with duplicate homepage url %s", s, s.HomePageUrl)
			}

			if s.ID != "" && s.ID != e.ID && e.HomePageUrl == s.HomePageUrl {
				// duplication for updating, checking name
				return fmt.Errorf("error trying to update suppliers %v with duplicate homepage url %s", s, s.HomePageUrl)
			}
		}

		// auto fill IDs
		for _, s := range suppliers {
			if s.ID == "" || s.CreatedAt == 0 {
				s.CreatedAt = now.Unix()
				s.ID = uuid.Must(uuid.NewV4()).String()
			}
			s.UpdatedAt = now.Unix()
		}
		batch.Set(d.client.Collection(d.path).Doc(s.ID), structs.Map(s), firestore.MergeAll)
	}

	// Commit the batch.
	_, err = batch.Commit(ctx)
	if err != nil {
		logrus.Errorf("error occur while batch set suppliers @UpsertSuppliers: %s", err.Error())
		return err
	}

	return nil
}

func (d *suppliers) GetAllSuppliers(ctx context.Context) (suppliers []*psuppliers.Supplier, err error) {
	t, _ := d.client.Collection(d.path).Documents(ctx).GetAll()
	for _, v := range t {
		var s psuppliers.Supplier
		if err = v.DataTo(&s); err != nil {
			return nil, err
		}
		suppliers = append(suppliers, &s)
	}

	return
}

func validateSuppliers(suppliers []*psuppliers.Supplier) (err error) {
	if len(suppliers) == 0 {
		return errors.New("empty suppliers")
	}
	for _, v := range suppliers {
		conform.Strings(v)
		// UUID checking
		if v.ID != "" && !govalidator.IsUUID(v.ID) {
			return fmt.Errorf("invalid UUID %s", v.ID)
		}

		if v.Slug == "" {
			v.Slug = slug.Make(v.Name)
		}
		_, err := govalidator.ValidateStruct(v)
		if err != nil {
			return err
		}
	}

	return nil
}
