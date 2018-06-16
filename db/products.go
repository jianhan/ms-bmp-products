package db

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/asaskevich/govalidator"
	"github.com/fatih/structs"
	"github.com/gosimple/slug"
	pproducts "github.com/jianhan/ms-bmp-products/proto/products"
	"github.com/leebenson/conform"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Products interface {
	UpsertProducts(ctx context.Context, products []*pproducts.Product) error
}

type products struct {
	Base
}

func NewProducts(path string, client *firestore.Client) Products {
	base := Base{
		path:   path,
		client: client,
	}
	return &products{base}
}

func (d *products) UpsertProducts(ctx context.Context, products []*pproducts.Product) error {
	// validation
	if err := validateProducts(products); err != nil {
		return err
	}

	// get all first
	existingProducts, err := d.getAllProducts(ctx)
	if err != nil {
		return err
	}

	// TODO: validate category urls

	// setup batch
	batch := d.client.Batch()
	now := time.Now()
	for _, p := range products {
		for _, e := range existingProducts {
			// start validation for unique constraints
			if e.Name == p.Name && p.ID == "" {
				// duplication for inserting, checking name
				return fmt.Errorf("error trying to insert new products %v with duplicate name %s", p, p.Name)
			}

			if p.ID != "" && p.ID != e.ID && e.Name == p.Name {
				// duplication for updating, checking name
				return fmt.Errorf("error trying to update products %v with duplicate name %s", p, p.Name)
			}

			if e.Slug == p.Slug && p.ID == "" {
				// duplication for inserting, checking slug
				return fmt.Errorf("error trying to insert new products %v with duplicate slug %s", p, p.Slug)
			}

			if p.ID != "" && p.ID != e.ID && e.Slug == p.Slug {
				// duplication for updating, checking slug
				return fmt.Errorf("error trying to update products %v with duplicate slug %s", p, p.Slug)
			}

			if e.Url == p.Url && p.ID == "" {
				// duplication for inserting, checking  url
				return fmt.Errorf("error trying to insert new products %v with duplicate homepage url %s", p, p.Url)
			}

			if p.ID != "" && p.ID != e.ID && e.Url == p.Url {
				// duplication for updating, checking name
				return fmt.Errorf("error trying to update products %v with duplicate homepage url %s", p, p.Url)
			}
		}

		// auto fill IDs and timestamp
		for _, s := range products {
			if s.ID == "" {
				s.ID = uuid.Must(uuid.NewV4()).String()
			} else {
				for k := range existingProducts {
					if existingProducts[k].ID == s.ID {
						s.CreatedAt = existingProducts[k].CreatedAt
						break
					}
				}
			}
			if s.CreatedAt == 0 {
				s.CreatedAt = now.Unix()
			}
			s.UpdatedAt = now.Unix()
		}
		batch.Set(d.client.Collection(d.path).Doc(p.ID), structs.Map(p), firestore.MergeAll)
	}

	// Commit the batch.
	_, err = batch.Commit(ctx)
	if err != nil {
		logrus.Errorf("error occur while batch set products: %s", err.Error())
		return err
	}

	return nil
}

func (f *products) getAllProducts(ctx context.Context) (products []*pproducts.Product, err error) {
	ps, _ := f.client.Collection(f.path).Documents(ctx).GetAll()
	for _, product := range ps {
		var p pproducts.Product
		if err = product.DataTo(&p); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}

	return
}

func validateProducts(products []*pproducts.Product) (err error) {
	for _, v := range products {
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
