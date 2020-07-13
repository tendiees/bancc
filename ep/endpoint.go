package ep

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/tendiees/bancc/svc"
)

type Set struct {
	Ping endpoint.Endpoint
}

func MakePing(s svc.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PingRequest)
		return PingResponse{Payload: s.Ping(ctx, req.Payload)}, nil
	}
}

type PingRequest struct {
	Payload string
}

type PingResponse struct {
	Payload string
}
