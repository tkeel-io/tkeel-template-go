package server

import (
	"github.com/tkeel-io/kit/transport/http"

	v1 "github.com/tkeel-io/tkeel-template-go/api/helloworld/v1"
	openapi_v1 "github.com/tkeel-io/tkeel-template-go/api/openapi/v1"
	"github.com/tkeel-io/tkeel-template-go/pkg/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(addr string, greeter *service.GreeterService, openapi *service.OpenapiService) *http.Server {
	srv := http.NewServer(addr)
	v1.RegisterGreeterHTTPServer(srv.Container, greeter)
	openapi_v1.RegisterOpenapiHTTPServer(srv.Container, openapi)
	return srv
}
