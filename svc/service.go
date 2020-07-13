package svc

import (
	"context"
)

type Service interface {
	Ping(ctx context.Context, str string) string
}

type service struct {
}

func New() Service {
	return &service{}
}

func (s service) Ping(ctx context.Context, str string) string {
	return str
}
