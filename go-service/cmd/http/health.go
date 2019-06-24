package http

import (
	"context"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/julienschmidt/httprouter"

	"github.com/wencan/kit-demo/go-service/cmd/endpoint"
	protocol "github.com/wencan/kit-demo/protocol/model"
)

func RegisterHealthHTTPHandlers(router *httprouter.Router, service endpoint.HealthService) error {
	router.Handler(http.MethodGet, "/health", transport.NewServer(endpoint.NewHealthCheckEndpoint(service), decodeCheckRequest, encodeResponse))
	return nil
}

func decodeCheckRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	req := &protocol.HealthCheckRequest{}
	err := decodeRequest(ctx, r, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
