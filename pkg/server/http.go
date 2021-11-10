package server

import (
	"github.com/tkeel-io/kit/transport/http"

	v1 "github.com/tkeel-io/tkeel-template-go/api/helloworld/v1"
	"github.com/tkeel-io/tkeel-template-go/pkg/service"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(addr string, greeter *service.GreeterService) *http.Server {
	srv := http.NewServer(addr)
	v1.RegisterGreeterHTTPServer(srv.Container, greeter)
	return srv
}
