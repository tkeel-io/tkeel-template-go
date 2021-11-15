// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 3.3.27

package v1

import (
	context "context"
	json "encoding/json"
	go_restful "github.com/emicklei/go-restful"
	v1 "github.com/tkeel-io/tkeel-interface/openapi/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.go_restful.json.

type OpenapiHTTPServer interface {
	AddonsIdentify(context.Context, *v1.AddonsIdentifyRequest) (*v1.AddonsIdentifyResponse, error)
	Identify(context.Context, *emptypb.Empty) (*v1.IdentifyResponse, error)
	Tatus(context.Context, *emptypb.Empty) (*v1.StatusResponse, error)
	TenantBind(context.Context, *v1.TenantBindRequst) (*v1.TenantBindResponse, error)
	TenantUnbind(context.Context, *v1.TenantUnbindRequst) (*v1.TenantUnbindResponse, error)
}

type OpenapiHTTPHandler struct {
	srv OpenapiHTTPServer
}

func newOpenapiHTTPHandler(s OpenapiHTTPServer) *OpenapiHTTPHandler {
	return &OpenapiHTTPHandler{srv: s}
}

func (h *OpenapiHTTPHandler) AddonsIdentify(req *go_restful.Request, resp *go_restful.Response) {
	in := &v1.AddonsIdentifyRequest{}
	if err := transportHTTP.GetBody(req, in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.AddonsIdentify(req.Request.Context(), in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *OpenapiHTTPHandler) Identify(req *go_restful.Request, resp *go_restful.Response) {
	in := &emptypb.Empty{}
	if err := transportHTTP.GetQuery(req, in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.Identify(req.Request.Context(), in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *OpenapiHTTPHandler) Tatus(req *go_restful.Request, resp *go_restful.Response) {
	in := &emptypb.Empty{}
	if err := transportHTTP.GetQuery(req, in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.Tatus(req.Request.Context(), in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *OpenapiHTTPHandler) TenantBind(req *go_restful.Request, resp *go_restful.Response) {
	in := &v1.TenantBindRequst{}
	if err := transportHTTP.GetBody(req, in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.TenantBind(req.Request.Context(), in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *OpenapiHTTPHandler) TenantUnbind(req *go_restful.Request, resp *go_restful.Response) {
	in := &v1.TenantUnbindRequst{}
	if err := transportHTTP.GetBody(req, in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.TenantUnbind(req.Request.Context(), in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func RegisterOpenapiHTTPServer(container *go_restful.Container, srv OpenapiHTTPServer) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/v1")
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	handler := newOpenapiHTTPHandler(srv)
	ws.Route(ws.GET("/identify").
		To(handler.Identify))
	ws.Route(ws.POST("/addons/identify").
		To(handler.AddonsIdentify))
	ws.Route(ws.POST("/tenant/bind").
		To(handler.TenantBind))
	ws.Route(ws.POST("/tenant/unbind").
		To(handler.TenantUnbind))
	ws.Route(ws.GET("/status").
		To(handler.Tatus))
}
