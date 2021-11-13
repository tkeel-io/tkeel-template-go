package server

import (
	"github.com/tkeel-io/kit/transport/grpc"
	v1 "github.com/tkeel-io/tkeel-template-go/api/helloworld/v1"
	openapi_v1 "github.com/tkeel-io/tkeel-template-go/api/openapi/v1"
	"github.com/tkeel-io/tkeel-template-go/pkg/service"
)

// NewHTTPServer new a GRPC server.
func NewGRPCServer(addr string, greeter *service.GreeterService, openapi *service.OpenapiService) *grpc.Server {
	s := grpc.NewServer(addr)
	v1.RegisterGreeterServer(s.GetServe(), greeter)
	openapi_v1.RegisterOpenapiServer(s.GetServe(), openapi)
	return s
}
