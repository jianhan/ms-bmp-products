package firebase

import (
	"context"
	"sync"

	"fmt"

	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/jianhan/ms-bmp-products/db"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var (
	c    *auth.Client
	once sync.Once
)

func Repository() db.Database {
	once.Do(func() {
		opt := option.WithCredentialsFile("./admin_sdk.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			logrus.Fatal(err)
		}
		c, err = app.Auth(context.Background())
		if err != nil {
			logrus.Fatal(err)
			panic(fmt.Errorf("error getting Auth client: %v\n", err))
		}
	})

	return r
}
