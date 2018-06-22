package main

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jianhan/ms-bmp-products/handlers"
	"github.com/jianhan/ms-bmp-products/mongodb"
	pcategories "github.com/jianhan/ms-bmp-products/proto/categories"
	cfgreader "github.com/jianhan/pkg/configs"
	"github.com/micro/go-micro"
	"github.com/nats-io/go-nats-streaming"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gopkg.in/mgo.v2"
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
	spew.Dump(viper.Get("mongodb.host"))
	session, err := mgo.Dial(viper.GetString("mongodb.host"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//// register suppliers handler
	//psuppliers.RegisterSuppliersServiceHandler(
	//	srv.Server(),
	//	handlers.NewSuppliersHandler(db.NewSuppliers("suppliers", firestoreClient), sc),
	//)
	//
	//// register products handler
	//pproducts.RegisterProductsServiceHandler(
	//	srv.Server(),
	//	handlers.NewProductsHandler(db.NewProducts("products", firestoreClient), sc),
	//)

	// register categories handler
	pcategories.RegisterCategoriesServiceHandler(
		srv.Server(),
		handlers.NewCategoriesHandler(mongodb.NewCategories(session, "categories"), sc),
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
