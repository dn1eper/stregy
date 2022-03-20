package composites

import (
	"context"
	"stregy/internal/adapters/pgorm/migration"
	"stregy/pkg/client/pgorm"

	"gorm.io/gorm"
)

type PGormComposite struct {
	db *gorm.DB
}

func NewPGormComposite(ctx context.Context, Host, Port, Username, Password, Database string) (*PGormComposite, error) {
	client, err := pgorm.NewClient(ctx, Host, Port, Username, Password, Database)
	migration.Migrate(client)
	if err != nil {
		return nil, err
	}
	return &PGormComposite{db: client}, nil
}
