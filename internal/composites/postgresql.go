package composites

import (
	"context"
	"stregy/pkg/client/postgresql"
)

type PostgreSQLComposite struct {
	db postgresql.Client
}

func NewPostgreSQLComposite(ctx context.Context, Host, Port, Username, Password, Database string) (*PostgreSQLComposite, error) {
	client, err := postgresql.NewClient(ctx, 1, Host, Port, Username, Password, Database)
	if err != nil {
		return nil, err
	}
	return &PostgreSQLComposite{db: client}, nil
}
