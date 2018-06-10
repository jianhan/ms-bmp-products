package firebase

import (
	"context"
	"sync"

	"fmt"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

var (
	c    *firestore.Client
	once sync.Once
)

func NewFirestoreClient(ctx context.Context) *firestore.Client {
	once.Do(func() {
		opt := option.WithCredentialsFile("./admin_sdk.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			logrus.Fatal(err)
			panic(fmt.Errorf("can not initialize new firebase app : %v\n", err))
		}
		c, err = app.Firestore(ctx)
		if err != nil {
			logrus.Fatal(err)
			panic(fmt.Errorf("error getting Auth client: %v\n", err))
		}
	})

	return c
}
