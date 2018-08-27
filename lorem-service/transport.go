package lorem_service

import (
	"context"
	"go-kit-service/lorem-service/proto"
)

//Encode and Decode Lorem Request
func EncodeGRPCLoremRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(LoremRequest)
	return &pb.LoremRequest{
		RequestType: req.RequestType,
		Max: req.Max,
		Min: req.Min,
	} , nil
}

func DecodeGRPCLoremRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoremRequest)
	return LoremRequest{
		RequestType: req.RequestType,
		Max: req.Max,
		Min: req.Min,
	}, nil
}

// Encode and Decode Lorem Response
func EncodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(LoremResponse)
	return &pb.LoremResponse{
		Message: resp.Message,
		Err: resp.Err,
	}, nil
}

func DecodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.LoremResponse)
	return LoremResponse{
		Message: resp.Message,
		Err: resp.Err,
	}, nil
}