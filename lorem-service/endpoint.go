package lorem_service

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"errors"
)

//request
type LoremRequest struct {
	RequestType string
	Min         int32
	Max         int32
}

//response
type LoremResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

func MakeLoremEndpoint(ls Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(LoremRequest)
		min := int(req.Min)
		max := int(req.Max)
		message, err := ls.Lorem(ctx, req.RequestType, min, max)
		return LoremResponse{Message: message}, err
	}
}

// endpoints wrapper
type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

// Wrapping Endpoints as a Service implementation.
// Will be used in gRPC client
func (e Endpoints) Lorem(ctx context.Context, requestType string, min, max int) (string, error) {
	req := LoremRequest{
		RequestType: requestType,
		Min:         int32(min),
		Max:         int32(max),
	}
	resp, err := e.LoremEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	loremResp := resp.(LoremResponse)
	if loremResp.Err != "" {
		return "", errors.New(loremResp.Err)
	}
	return loremResp.Message, nil
}
