package migration

import (
	"stregy/internal/adapters/pggorm/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&user.User{})
}
