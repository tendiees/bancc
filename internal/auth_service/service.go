package auth_service

import (
	"context"
	"errors"

	"github.com/tendiees/bancc/internal/crypt"
)

type Service interface {
	Register(ctx context.Context, username string, password string) (err error)
	Auth(ctx context.Context, username string, password string) (token string, err error)
	Verify(ctx context.Context, token string) (ok bool, err error)
}

type service struct {
	store Store
}

func (s service) Register(ctx context.Context, username string, password string) (err error) {
	panic("implement me")
}

var ErrUnauthorized = errors.New("unauthorized")

func (s service) Auth(ctx context.Context, username string, password string) (token string, err error) {
	h, err := s.store.Get(ctx, username)
	if err != nil {
		return "", err
	}

	if crypt.Verify([]byte(password), h) {
	}
}

func (s service) Verify(ctx context.Context, token string) (ok bool, err error) {
	panic("implement me")
}

func New(store Store) Service {
	return service{}
}
