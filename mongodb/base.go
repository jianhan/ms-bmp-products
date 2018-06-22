package mongodb

import "gopkg.in/mgo.v2"

type base struct {
	session    *mgo.Session
	db         string
	collection string
}
