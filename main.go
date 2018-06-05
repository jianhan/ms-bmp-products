package main

import (
	"fmt"
	"log"
	"os"
	"time"

	cfgreader "github.com/jianhan/pkg/configs"
	"github.com/micro/go-micro"
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
	srv.Init()
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	if os.Getenv("environment") == "" {
		os.Setenv("environment", "development")
	}
}
