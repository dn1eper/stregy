package composites

import (
	"stregy/internal/adapters/pgorm/migration"
	"stregy/pkg/client/pgorm"

	"gorm.io/gorm"
)

type PGormComposite struct {
	db *gorm.DB
}

func NewPGormComposite(Host, Port, Username, Password, Database string) (*PGormComposite, error) {
	client, _ := pgorm.NewClient(Host, Port, Username, Password, Database)
	err := migration.Migrate(client)
	if err != nil {
		return nil, err
	}
	return &PGormComposite{db: client}, nil
}
