package db

import (
	"context"
	"fmt"

	"time"

	"cloud.google.com/go/firestore"
	"github.com/asaskevich/govalidator"
	"github.com/fatih/structs"
	"github.com/gosimple/slug"
	"github.com/jianhan/ms-bmp-products/firebase"
	pproducts "github.com/jianhan/ms-bmp-products/proto/product"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/supplier"
	"github.com/leebenson/conform"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type firestoreDB struct {
	client         *firestore.Client
	categoriesPath string
	productsPath   string
}

func NewFirestoreDB(ctx context.Context) Database {
	return &firestoreDB{
		client:         firebase.NewFirestoreClient(ctx),
		categoriesPath: "suppliers",
		productsPath:   "products",
	}
}

func (f *firestoreDB) UpsertSuppliers(ctx context.Context, suppliers []*psuppliers.Supplier) error {
	// validation
	if err := validateSuppliers(suppliers); err != nil {
		return err
	}

	// get all first
	existingSuppliers, err := f.getAllSuppliers(ctx)
	if err != nil {
		return err
	}

	// setup batch
	batch := f.client.Batch()
	now := time.Now()
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

			if e.Slug == s.Slug && s.ID == "" {
				// duplication for inserting, checking slug
				return fmt.Errorf("error trying to insert new supplier %v with duplicate slug %s", s, s.Slug)
			}

			if s.ID != "" && s.ID != e.ID && e.Slug == s.Slug {
				// duplication for updating, checking slug
				return fmt.Errorf("error trying to update supplier %v with duplicate slug %s", s, s.Slug)
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
			if s.ID == "" || s.CreatedAt == 0 {
				s.CreatedAt = now.Unix()
				s.ID = uuid.Must(uuid.NewV4()).String()
			}
			s.UpdatedAt = now.Unix()
		}
		batch.Set(f.client.Collection(f.categoriesPath).Doc(s.ID), structs.Map(s), firestore.MergeAll)
	}

	// Commit the batch.
	_, err = batch.Commit(ctx)
	if err != nil {
		logrus.Errorf("error occur while batch set suppliers @UpsertSuppliers: %s", err.Error())
		return err
	}

	return nil
}

func (f *firestoreDB) getAllSuppliers(ctx context.Context) (suppliers []*psuppliers.Supplier, err error) {
	t, _ := f.client.Collection(f.categoriesPath).Documents(ctx).GetAll()
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

func (f *firestoreDB) UpsertProducts(ctx context.Context, products []*pproducts.Product) error {
	// get all first
	existingSuppliers, err := f.getAllProducts(ctx)
	if err != nil {
		return err
	}

	// setup batch
	batch := f.client.Batch()
	now := time.Now()
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

			if e.Slug == s.Slug && s.ID == "" {
				// duplication for inserting, checking slug
				return fmt.Errorf("error trying to insert new supplier %v with duplicate slug %s", s, s.Slug)
			}

			if s.ID != "" && s.ID != e.ID && e.Slug == s.Slug {
				// duplication for updating, checking slug
				return fmt.Errorf("error trying to update supplier %v with duplicate slug %s", s, s.Slug)
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
			if s.ID == "" || s.CreatedAt == 0 {
				s.CreatedAt = now.Unix()
				s.ID = uuid.Must(uuid.NewV4()).String()
			}
			s.UpdatedAt = now.Unix()
		}
		batch.Set(f.client.Collection(f.categoriesPath).Doc(s.ID), structs.Map(s), firestore.MergeAll)
	}

	// Commit the batch.
	_, err = batch.Commit(ctx)
	if err != nil {
		logrus.Errorf("error occur while batch set suppliers @UpsertSuppliers: %s", err.Error())
		return err
	}

	return nil
}

func (f *firestoreDB) getAllProducts(ctx context.Context) (products []*pproducts.Product, err error) {
	ps, _ := f.client.Collection(f.productsPath).Documents(ctx).GetAll()
	for _, product := range ps {
		var p pproducts.Product
		if err = product.DataTo(&p); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return
}

func (f *firestoreDB) Close() error {
	f.client.Close()
}
