package server

import (
	"github.com/tkeel-io/kit/transport/grpc"
	v1 "github.com/tkeel-io/tkeel-template-go/api/helloworld/v1"
	"github.com/tkeel-io/tkeel-template-go/pkg/service"
)

// NewHTTPServer new a GRPC server.
func NewGRPCServer(addr string, greeter *service.GreeterService) *grpc.Server {
	s := grpc.NewServer(addr)
	v1.RegisterGreeterServer(s.GetServe(), greeter)
	return s
}
