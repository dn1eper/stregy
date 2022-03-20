package composites

import (
	"context"
	"stregy/internal/adapters/pggorm/migration"
	"stregy/pkg/client/pggorm"

	"gorm.io/gorm"
)

type PGGormComposite struct {
	db *gorm.DB
}

func NewPGGormComposite(ctx context.Context, Host, Port, Username, Password, Database string) (*PGGormComposite, error) {
	client, err := pggorm.NewClient(ctx, Host, Port, Username, Password, Database)
	migration.Migrate(client)
	if err != nil {
		return nil, err
	}
	return &PGGormComposite{db: client}, nil
}
