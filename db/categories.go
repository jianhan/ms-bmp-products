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
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
	"github.com/leebenson/conform"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Categories interface {
	UpsertCategories(ctx context.Context, categories []*pcategories.Category) error
	GetAllCategories(ctx context.Context) (categories []*pcategories.Category, err error)
}

type categories struct {
	Base
}

func NewCategories(path string, client *firestore.Client) Categories {
	base := Base{
		path:   path,
		client: client,
	}

	return &categories{base}
}

func (d *categories) UpsertCategories(ctx context.Context, categories []*pcategories.Category) error {
	// validation
	if err := validateCategories(categories); err != nil {
		return err
	}

	// get all first
	existingCategories, err := d.GetAllCategories(ctx)
	if err != nil {
		return err
	}

	// setup batch
	batch := d.client.Batch()
	now := time.Now()
	for _, s := range categories {
		for _, e := range existingCategories {
			// start validation for unique constraints
			if e.Name == s.Name && s.ID == "" {
				// duplication for inserting, checking name
				return fmt.Errorf("error trying to insert new categories %v with duplicate name %s", s, s.Name)
			}

			if s.ID != "" && s.ID != e.ID && e.Name == s.Name {
				// duplication for updating, checking name
				return fmt.Errorf("error trying to update categories %v with duplicate name %s, already exists %s", s, s.Name, e.ID)
			}

			if e.Slug == s.Slug && s.ID == "" {
				// duplication for inserting, checking slug
				return fmt.Errorf("error trying to insert new categories %v with duplicate slug %s", s, s.Slug)
			}

			if s.ID != "" && s.ID != e.ID && e.Slug == s.Slug {
				// duplication for updating, checking slug
				return fmt.Errorf("error trying to update categories %v with duplicate slug %s", s, s.Slug)
			}

			if e.Url == s.Url && s.ID == "" {
				// duplication for inserting, checking homepage url
				return fmt.Errorf("error trying to insert new categories %v with duplicate url %s", s, s.Url)
			}

			if s.ID != "" && s.ID != e.ID && e.Url == s.Url {
				// duplication for updating, checking name
				return fmt.Errorf("error trying to update categories %v with duplicate url %s", s, s.Url)
			}
		}

		// auto fill IDs and timestamp
		for _, s := range categories {
			if s.ID == "" {
				s.ID = uuid.Must(uuid.NewV4()).String()
			} else {
				for k := range existingCategories {
					if existingCategories[k].ID == s.ID {
						s.CreatedAt = existingCategories[k].CreatedAt
						break
					}
				}
			}
			if s.CreatedAt == 0 {
				s.CreatedAt = now.Unix()
			}
			s.UpdatedAt = now.Unix()
		}

		batch.Set(d.client.Collection(d.path).Doc(s.ID), structs.Map(s), firestore.MergeAll)
	}

	// Commit the batch.
	_, err = batch.Commit(ctx)
	if err != nil {
		logrus.Errorf("error occur while batch set categories: %s", err.Error())
		return err
	}

	return nil
}

func (d *categories) GetAllCategories(ctx context.Context) (categories []*pcategories.Category, err error) {
	t, _ := d.client.Collection(d.path).Documents(ctx).GetAll()
	for _, v := range t {
		var s pcategories.Category
		if err = v.DataTo(&s); err != nil {
			return nil, err
		}
		categories = append(categories, &s)
	}

	return
}

func validateCategories(categories []*pcategories.Category) (err error) {
	if len(categories) == 0 {
		return errors.New("empty categories")
	}
	for _, v := range categories {
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
