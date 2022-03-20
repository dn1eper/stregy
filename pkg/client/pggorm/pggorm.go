package pggorm

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient(ctx context.Context, username, password, host, port, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", host, port, username, password, database)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
