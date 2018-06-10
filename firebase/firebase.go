package firebase

import (
	"context"
	"sync"

	"fmt"

	"firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var (
	c    *db.Client
	once sync.Once
)

func NewFirebaseClient() *db.Client {
	once.Do(func() {
		opt := option.WithCredentialsFile("./admin_sdk.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			logrus.Fatal(err)
			panic(fmt.Errorf("can not initialize new firebase app : %v\n", err))
		}
		c, err = app.Database(context.Background())
		if err != nil {
			logrus.Fatal(err)
			panic(fmt.Errorf("error getting Auth client: %v\n", err))
		}
	})

	return c
}
