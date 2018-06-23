package mongodb

import (
	"github.com/asaskevich/govalidator"
	"github.com/gosimple/slug"
	"github.com/jianhan/ms-bmp-products/db"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
	"github.com/leebenson/conform"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Suppliers struct {
	base
}

func NewSuppliers(session *mgo.Session, collection string) db.Suppliers {
	c := session.DB(viper.GetString("mongodb.db")).C(collection)
	displayOrderIndex := mgo.Index{
		Key: []string{"display_order"},
	}
	if err := c.EnsureIndex(displayOrderIndex); err != nil {
		panic(err)
	}
	nameIndex := mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	}
	if err := c.EnsureIndex(nameIndex); err != nil {
		panic(err)
	}
	slugIndex := mgo.Index{
		Key:    []string{"slug"},
		Unique: true,
	}
	if err := c.EnsureIndex(slugIndex); err != nil {
		panic(err)
	}
	homePageURLIndex := mgo.Index{
		Key:    []string{"home_page_url"},
		Unique: true,
	}
	if err := c.EnsureIndex(homePageURLIndex); err != nil {
		panic(err)
	}
	textIndex := mgo.Index{
		Key: []string{"$text:name"},
	}
	if err := c.EnsureIndex(textIndex); err != nil {
		panic(err)
	}
	b := base{
		session:    session,
		db:         viper.GetString("mongodb.db"),
		collection: collection,
	}

	return &Suppliers{
		base: b,
	}
}

func (s *Suppliers) UpsertSuppliers(suppliers []*psuppliers.Supplier) (int64, int64, error) {
	bulk := s.session.DB(s.db).C(s.collection).Bulk()
	for _, supplier := range suppliers {
		conform.Strings(supplier)
		s.beforeUpsert(supplier)
		supplier.Slug = slug.Make(supplier.Name)
		if _, err := govalidator.ValidateStruct(supplier); err != nil {
			return 0, 0, err
		}
		bulk.Upsert(
			bson.M{"_id": supplier.ID},
			bson.M{"$set": supplier},
		)
	}
	r, err := bulk.Run()
	if err != nil {
		return 0, 0, err
	}

	return int64(r.Matched), int64(r.Modified), nil
}

func (s *Suppliers) Suppliers() (suppliers []*psuppliers.Supplier, err error) {
	var r []*psuppliers.Supplier
	if err = s.session.DB(s.db).C(s.collection).Find(nil).All(&r); err != nil {
		return
	}

	return r, nil
}
