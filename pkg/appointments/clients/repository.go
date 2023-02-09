package clients

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func (r *repository) FetchByUsername(ctx context.Context, username string) (ClientEntity, error) {
	return ClientEntity{}, nil
}
