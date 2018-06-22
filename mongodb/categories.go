package mongodb

import (
	"github.com/asaskevich/govalidator"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/jianhan/ms-bmp-products/db"
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
	"github.com/leebenson/conform"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Categories struct {
	base
}

func NewCategories(session *mgo.Session, collection string) db.Categories {
	c := session.DB(viper.GetString("mongodb.db")).C(collection)
	hiddenIndex := mgo.Index{
		Key: []string{"hidden"},
	}
	if err := c.EnsureIndex(hiddenIndex); err != nil {
		panic(err)
	}
	displayOrderIndex := mgo.Index{
		Key: []string{"display_order"},
	}
	if err := c.EnsureIndex(displayOrderIndex); err != nil {
		panic(err)
	}
	slugIndex := mgo.Index{
		Key:    []string{"slug"},
		Unique: true,
	}
	if err := c.EnsureIndex(slugIndex); err != nil {
		panic(err)
	}
	urlIndex := mgo.Index{
		Key:    []string{"url"},
		Unique: true,
	}
	if err := c.EnsureIndex(urlIndex); err != nil {
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
	return &Categories{
		base: b,
	}
}

func (c *Categories) UpsertCategories(categories []*pcategories.Category) (int64, int64, error) {
	bulk := c.session.DB(c.db).C(c.collection).Bulk()
	for _, category := range categories {
		conform.Strings(category)
		category.Slug = slug.Make(category.Name)
		category.UpdatedAt = ptypes.TimestampNow()
		if category.CreatedAt == nil {
			category.CreatedAt = ptypes.TimestampNow()
		}
		if category.ID == "" {
			category.ID = uuid.New().String()
		}
		if _, err := govalidator.ValidateStruct(category); err != nil {
			return 0, 0, err
		}
		bulk.Upsert(
			bson.M{"_id": category.ID},
			bson.M{"$set": category},
		)
	}
	r, err := bulk.Run()
	if err != nil {
		return 0, 0, err
	}

	return int64(r.Matched), int64(r.Modified), nil
}

func (c *Categories) Categories() (categories []*pcategories.Category, err error) {
	var r []*pcategories.Category
	if err = c.session.DB(c.db).C(c.collection).Find(nil).All(&r); err != nil {
		return
	}
	return r, nil
}

//func (c *Categories) InsertCategories(categories *pcategory.Categories) (int64, error) {
//	bulk := c.session.DB(c.db).C(c.collection).Bulk()
//	for _, v := range categories.Categories {
//		bulk.Insert(v)
//	}
//	r, err := bulk.Run()
//	if err != nil {
//		return 0, err
//	}
//	return int64(r.Modified), nil
//}
//
//func (c *Categories) UpdateCategories(categories *pcategory.Categories) (int64, error) {
//	bulk := c.session.DB(c.db).C(c.collection).Bulk()
//	for _, v := range categories.Categories {
//		bulk.Update(bson.M{"_id": v.ID}, v)
//	}
//	r, err := bulk.Run()
//	if err != nil {
//		return 0, err
//	}
//	return int64(r.Modified), nil
//}
//
//func (c *Categories) FindCategories(req *pcategory.FindCategoriesRequest) (categories []*pcategory.Category, err error) {
//	// define query
//	query := bson.M{}
//	var r []*pcategory.Category
//	// set IDs query condition
//	if len(req.Ids) > 0 {
//		query["_id"] = bson.M{"$in": req.Ids}
//	}
//	// set hidden condition
//	query["hidden"] = bson.M{"$eq": req.Hidden}
//	// set search
//	if strings.TrimSpace(req.Query) != "" {
//		query["$text"] = bson.M{"$search": req.Query}
//	}
//	// set sort condition
//	sorts := []string{"display_order", "name"}
//	if req.Sort != nil {
//		sorts = req.Sort
//	}
//	// set pagination
//	currentPage := c.currentPage
//	if req.CurrentPage > 0 {
//		currentPage = int(req.CurrentPage)
//	}
//	perPage := c.perPage
//	if req.PerPage > 0 {
//		perPage = int(req.PerPage)
//	}
//	// get result
//	if err = c.session.DB(c.db).C(c.collection).Find(query).Sort(sorts...).Skip(perPage * (currentPage - 1)).Limit(perPage).All(&r); err != nil {
//		return
//	}
//	return r, nil
//}
//
//func (c *Categories) DeleteCategoriesByIDs(ids []string) error {
//	if len(ids) == 0 {
//		return nil
//	}
//	if err := c.session.DB(c.db).C(c.collection).Remove(bson.M{"_id": bson.M{"$in": ids}}); err != nil {
//		return err
//	}
//	return nil
//}
