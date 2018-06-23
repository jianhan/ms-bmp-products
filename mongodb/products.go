package mongodb

import (
	"github.com/asaskevich/govalidator"
	"github.com/gosimple/slug"
	"github.com/jianhan/ms-bmp-products/db"
	pproducts "github.com/jianhan/ms-bmp-products/proto/products"
	"github.com/leebenson/conform"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Products struct {
	base
}

func NewProducts(session *mgo.Session, collection string) db.Products {
	c := session.DB(viper.GetString("mongodb.db")).C(collection)
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
	priceIndex := mgo.Index{
		Key: []string{"price"},
	}
	if err := c.EnsureIndex(priceIndex); err != nil {
		panic(err)
	}
	ratingIndex := mgo.Index{
		Key: []string{"rating"},
	}
	if err := c.EnsureIndex(ratingIndex); err != nil {
		panic(err)
	}
	hiddenIndex := mgo.Index{
		Key: []string{"hidden"},
	}
	if err := c.EnsureIndex(hiddenIndex); err != nil {
		panic(err)
	}
	b := base{
		session:    session,
		db:         viper.GetString("mongodb.db"),
		collection: collection,
	}

	return &Products{
		base: b,
	}
}

func (p *Products) UpsertProducts(products []*pproducts.Product) (int64, int64, error) {
	bulk := p.session.DB(p.db).C(p.collection).Bulk()
	for _, product := range products {
		conform.Strings(product)
		p.beforeUpsert(product)
		product.Slug = slug.Make(product.Name)
		if _, err := govalidator.ValidateStruct(product); err != nil {
			return 0, 0, err
		}
		bulk.Upsert(
			bson.M{"_id": product.ID},
			bson.M{"$set": product},
		)
	}
	r, err := bulk.Run()
	if err != nil {
		return 0, 0, err
	}

	return int64(r.Matched), int64(r.Modified), nil
}

func (p *Products) Products() (products []*pproducts.Product, err error) {
	var r []*pproducts.Product
	if err = p.session.DB(p.db).C(p.collection).Find(nil).All(&r); err != nil {
		return
	}

	return r, nil
}
