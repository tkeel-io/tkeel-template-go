package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/tkeel-io/kit/app"
	"github.com/tkeel-io/kit/log"
	"github.com/tkeel-io/tkeel-template-go/pkg/server"
	"github.com/tkeel-io/tkeel-template-go/pkg/service"
)

var (
	// app name
	Name string
	// http addr
	HTTPAddr string
	// grpc addr
	GRPCAddr string
)

func init() {
	flag.StringVar(&Name, "name", "hello", "app name.")
	flag.StringVar(&HTTPAddr, "http_addr", ":31234", "http listen address.")
	flag.StringVar(&GRPCAddr, "grpc_addr", ":31233", "grpc listen address.")
}

func main() {
	flag.Parse()

	greeterSrv := service.NewGreeterService()
	openapiSrv := service.NewOpenapiService()

	app := app.New(Name,
		&log.Conf{
			App:   Name,
			Level: "debug",
			Dev:   true,
		},
		server.NewHTTPServer(HTTPAddr, greeterSrv, openapiSrv),
		server.NewGRPCServer(GRPCAddr, greeterSrv, openapiSrv),
	)
	if err := app.Run(context.TODO()); err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop

	if err := app.Stop(context.TODO()); err != nil {
		panic(err)
	}
}
