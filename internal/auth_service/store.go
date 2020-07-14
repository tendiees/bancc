package auth_service

import (
	"context"

	"github.com/tendiees/bancc/internal/crypt"
)

type Store interface {
	Get(ctx context.Context, username string) ([crypt.Size]byte, error)
	Set(ctx context.Context, username string, hash [crypt.Size]byte) error
}

type StoreMiddleware func(store Store) Store

func NewPGStoreFromDSN(dsn string) {
}
