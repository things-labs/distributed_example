package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	consulKit "github.com/go-kit/kit/sd/consul"
	httptransport "github.com/go-kit/kit/transport/http"

	stdconsul "github.com/hashicorp/consul/api"
)

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")

// define sever

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int { return len(s) }

// define endpoint

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

// define encode and decode

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

const RegistryAddr = "127.0.0.1:8500"

func main() {
	svc := stringService{}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
	router := gin.Default()
	router.GET("/uppercase", gin.WrapH(uppercaseHandler))
	router.GET("/count", gin.WrapH(countHandler))

	asr := stdconsul.AgentServiceRegistration{
		Kind:              stdconsul.ServiceKindTypical,
		ID:                "1111",
		Name:              "stringsvc",
		Tags:              nil,
		Port:              8080,
		Address:           "127.0.0.1",
		TaggedAddresses:   nil,
		EnableTagOverride: false,
		Meta:              nil,
		Weights:           nil,
		Check:             nil,
		Checks:            nil,
		Proxy:             nil,
		Connect:           nil,
	}

	discoverClient, err := NewDiscoveryClient(RegistryAddr, &asr)
	if err != nil {
		log.Fatal(err)
	}

	if err = discoverClient.Register(); err != nil {
		log.Fatal(err)
	}
	log.Fatal(router.Run(":8080"))
}

type DiscoverClient struct {
	client       consulKit.Client
	registration *stdconsul.AgentServiceRegistration
}

func NewDiscoveryClient(addr string, registration *stdconsul.AgentServiceRegistration) (*DiscoverClient, error) {
	c := stdconsul.DefaultConfig()
	c.Address = addr
	consulClient, err := stdconsul.NewClient(c)
	if err != nil {
		return nil, err
	}
	return &DiscoverClient{
		consulKit.NewClient(consulClient),
		registration,
	}, nil
}

func (sf DiscoverClient) Register() error {
	return sf.client.Register(sf.registration)
}
func (sf DiscoverClient) Deregister() error {
	return sf.client.Deregister(sf.registration)
}
