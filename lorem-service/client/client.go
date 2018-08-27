package client

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"go-kit-service/lorem-service"
	"go-kit-service/lorem-service/proto"
)

// Return new lorem_grpc service
func New(conn *grpc.ClientConn) lorem_service.Service {
	var loremEndpoint = grpctransport.NewClient(
		conn, "Lorem", "Lorem",
		lorem_service.EncodeGRPCLoremRequest,
		lorem_service.DecodeGRPCLoremResponse,
		pb.LoremResponse{},
	).Endpoint()

	return lorem_service.Endpoints{LoremEndpoint: loremEndpoint}

}
