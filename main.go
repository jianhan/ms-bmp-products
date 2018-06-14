package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"context"

	"github.com/jianhan/ms-bmp-products/db"
	"github.com/jianhan/ms-bmp-products/firebase"
	"github.com/jianhan/ms-bmp-products/handlers"
	pproducts "github.com/jianhan/ms-bmp-products/proto/product"
	psuppliers "github.com/jianhan/ms-bmp-products/proto/supplier"
	cfgreader "github.com/jianhan/pkg/configs"
	"github.com/micro/go-micro"
	_ "github.com/spf13/viper/remote"
)

func main() {
	// Optional. Switch the session to a monotonic behavior.
	serviceConfigs, err := cfgreader.NewReader(os.Getenv("environment")).Read()
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
		handlers.NewSuppliersHandler(db.NewSuppliers("suppliers", firestoreClient)),
	)

	// register products handler
	pproducts.RegisterProductsServiceHandler(
		srv.Server(),
		handlers.NewProductsHandler(db.NewProducts("products", firestoreClient)),
	)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	// set default env if there is not one
	if os.Getenv("environment") == "" {
		os.Setenv("environment", "development")
	}
}
