package migration

import (
	"stregy/internal/adapters/pgorm/symbol"
	"stregy/internal/adapters/pgorm/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&user.User{}, &symbol.Symbol{})
}
