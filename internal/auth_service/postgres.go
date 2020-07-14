package auth_service

import (
	"context"
	"database/sql"

	"github.com/tendiees/bancc/internal/crypt"
)

type PGService struct {
	db *sql.DB
}

func (s PGService) Register(ctx context.Context, username string, password string) (err error) {
	panic("implement me")
}

func (s PGService) Auth(ctx context.Context, username string, password string) (token string, err error) {
	stmt, err := s.db.PrepareContext(ctx, "SELECT hash FROM credential WHERE username=?")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	r := stmt.QueryRowContext(ctx)
	buf := [crypt.Size]byte{}
	r.Scan(&buf)
	panic("implement me")
}

func (s PGService) Verify(ctx context.Context, token string) (ok bool, err error) {
	panic("implement me")
}

func NewPGFromDSN(dsn string) (Service, error) {
	db, err := sql.Open("postgres", dsn)
	return PGService{db: db}, err
}
