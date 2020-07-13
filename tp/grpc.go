package tp

import (
	"context"

	"github.com/go-kit/kit/transport/grpc"

	"github.com/tendiees/bancc/bancc"
	"github.com/tendiees/bancc/ep"
)

type GRPCServer struct {
	bancc.UnimplementedBanccServer
	ping grpc.Handler
}

func (s *GRPCServer) Ping(ctx context.Context, request *bancc.PingRequest) (*bancc.PingResponse, error) {
	_, res, err := s.ping.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return res.(*bancc.PingResponse), nil
}

func NewGRPC(eps ep.Set) bancc.BanccServer {
	return &GRPCServer{
		ping: grpc.NewServer(eps.Ping, decGRPCPingRequest, encGRPCPingResponse),
	}
}

func decGRPCPingRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*bancc.PingRequest)
	return ep.PingRequest{Payload: req.Payload}, nil
}

func encGRPCPingResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(ep.PingResponse)
	return &bancc.PingResponse{Payload: res.Payload}, nil
}
