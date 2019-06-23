package http

import (
	"context"
	"encoding/json"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"

	"github.com/wencan/kit-demo/go-service/cmd/endpoint"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

func RegisterHealthHTTPHandlers(router *httprouter.Router, service endpoint.HealthService) error {
	router.Handler(http.MethodGet, "/health", transport.NewServer(endpoint.NewHealthCheckEndpoint(service), decodeCheckRequest, encodeCommonResponse))
	return nil
}

func decodeCheckRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &protocol.HealthCheckRequest{
		Service: r.FormValue("service"),
	}
	return req, nil
}

func encodeCommonResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
