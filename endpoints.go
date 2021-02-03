package news

// endpoints.go contains the endpoint definitions, including per-method request
// and response structs. Endpoints are the binding between the service and
// transport.

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opentracing"
	stdopentracing "github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"
)

// Endpoints collects the endpoints that comprise the Service.
type Endpoints struct {
	ListEndpoint   endpoint.Endpoint
	HealthEndpoint endpoint.Endpoint
}

// MakeEndpoints returns an Endpoints structure, where each endpoint is
// backed by the given service.
func MakeEndpoints(s Service, tracer stdopentracing.Tracer) Endpoints {
	return Endpoints{
		ListEndpoint:   opentracing.TraceServer(tracer, "GET /news")(MakeListEndpoint(s)),
		HealthEndpoint: opentracing.TraceServer(tracer, "GET /health")(MakeHealthEndpoint(s)),
	}
}

// MakeListEndpoint returns an endpoint via the given service.
func MakeListEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		newsItems, err := s.List()
		return listResponse{NewsItems: newsItems, Err: err}, err
	}
}

// MakeHealthEndpoint returns current health of the given service.
func MakeHealthEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		health := s.Health()
		return healthResponse{Health: health}, nil
	}
}

type listRequest struct {
	//
}

type listResponse struct {
	NewsItems []NewsItem `json:"newsItem"`
	Err       error      `json:"err"`
}

type healthRequest struct {
	//
}

type healthResponse struct {
	Health []Health `json:"health"`
}
