package svc

import (
	"context"
	"database/sql"
)

type Service interface {
	Ping(ctx context.Context, str string) string
	Auth(ctx context.Context, username string, password string) (token string, err error)
}

type service struct {
	db *sql.DB
}

func New(dsn string) (Service, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &service{db: db}, nil
}

func (s service) Ping(ctx context.Context, str string) string {
	return str
}

func (s service) Auth(ctx context.Context, username string, password string) (string, error) {
	return username, nil
}
