package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/tkeel-io/kit/app"
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
	flag.StringVar(&HTTPAddr, "http_addr", "", "http listen address.")
	flag.StringVar(&GRPCAddr, "grpc_addr", "", "grpc listen address.")
}

func main() {
	flag.Parse()

	app := app.New(Name, HTTPAddr, GRPCAddr)
	if err := app.Run(context.TODO()); err != nil {
		panic(err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
	<-stop

	app.Stop(context.TODO())
	if err := app.Run(context.TODO()); err != nil {
		panic(err)
	}
}
