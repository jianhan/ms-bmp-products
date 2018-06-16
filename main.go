package main

import (
	"fmt"
	"os"
	"time"

	"context"

	"github.com/jianhan/ms-bmp-products/db"
	"github.com/jianhan/ms-bmp-products/firebase"
	"github.com/jianhan/ms-bmp-products/handlers"
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
	pproducts "github.com/jianhan/ms-bmp-products/proto/products"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/suppliers"
	cfgreader "github.com/jianhan/pkg/configs"
	"github.com/micro/go-micro"
	"github.com/nats-io/go-nats-streaming"
	_ "github.com/spf13/viper/remote"
)

func main() {
	// init nats streaming
	sc, err := stan.Connect(os.Getenv("NATS_STREAMING_CLUSTER"), os.Getenv("NATS_STREAMING_CLIENT_ID"))
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	serviceConfigs, err := cfgreader.NewReader(os.Getenv("ENVIRONMENT")).Read()
	if err != nil {
		panic(fmt.Sprintf("error while reading configurations: %s", err.Error()))
	}

	// initialize new service
	srv := micro.NewService(
		micro.Name(serviceConfigs.Name),
		micro.RegisterTTL(time.Duration(serviceConfigs.RegisterTTL)*time.Second),
		micro.RegisterInterval(time.Duration(serviceConfigs.RegisterInterval)*10),
		micro.Version(serviceConfigs.Version),
		micro.Metadata(serviceConfigs.Metadata),
	)

	// TODO: Added dynamic configs so do not have to rebuild service when configs changed
	// init service
	srv.Init()

	// init firestore and defer close it
	firestoreClient := firebase.NewFirestoreClient(context.Background())
	defer firestoreClient.Close()

	// register suppliers handler
	psuppliers.RegisterSuppliersServiceHandler(
		srv.Server(),
		handlers.NewSuppliersHandler(db.NewSuppliers("suppliers", firestoreClient), sc),
	)

	// register products handler
	pproducts.RegisterProductsServiceHandler(
		srv.Server(),
		handlers.NewProductsHandler(db.NewProducts("products", firestoreClient)),
	)

	// register categories handler
	pcategories.RegisterCategoriesServiceHandler(
		srv.Server(),
		handlers.NewCategoriesHandler(db.NewCategories("categories", firestoreClient)),
	)

	if err := srv.Run(); err != nil {
		panic(err)
	}

}

func init() {
	// set default env if there is not one
	if os.Getenv("ENVIRONMENT") == "" {
		os.Setenv("ENVIRONMENT", "development")
	}

	// set default nats configs
	if os.Getenv("NATS_STREAMING_CLUSTER") == "" {
		os.Setenv("NATS_STREAMING_CLUSTER", "test-cluster")
	}

	if os.Getenv("NATS_STREAMING_CLIENT_ID") == "" {
		os.Setenv("NATS_STREAMING_CLIENT_ID", "test-client")
	}
}
