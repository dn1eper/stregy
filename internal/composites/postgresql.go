package composites

import (
	"stregy/pkg/client/postgresql"
)

type PostgreSQLComposite struct {
	db postgresql.Client
}

func NewPostgreSQLComposite(Host, Port, Username, Password, Database string) (*PostgreSQLComposite, error) {
	client, err := postgresql.NewClient(1, Host, Port, Username, Password, Database)
	if err != nil {
		return nil, err
	}
	return &PostgreSQLComposite{db: client}, nil
}
