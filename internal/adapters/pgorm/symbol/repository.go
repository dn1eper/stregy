package symbol

import (
	"context"
	"fmt"
	"stregy/internal/domain/symbol"
	"strings"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) symbol.Repository {
	return &repository{db: client}
}

func (r *repository) Exists(ctx context.Context, name string) bool {
	symbol := &Symbol{Name: name}
	result := r.db.First(symbol)
	if result.Error != nil {
		return false
	}
	return true
}

func (r *repository) Create(ctx context.Context, name string) (*symbol.Symbol, error) {
	symbol := &Symbol{Name: name}
	result := r.db.Create(symbol)
	if result.Error != nil {
		return nil, result.Error
	}
	tableName := strings.ToLower(name) + "s"
	r.db.Exec(fmt.Sprintf("CREATE TABLE %vs (LIKE quotes INCLUDING ALL);", tableName))
	symbolDomain := symbol.ToDomain()
	return &symbolDomain, result.Error
}

func (r *repository) GetAll(ctx context.Context) ([]symbol.Symbol, error) {
	symbols := make([]Symbol, 0)
	result := r.db.Find(&symbols)
	symbolsDomain := make([]symbol.Symbol, 0, len(symbols))
	for _, symbol := range symbols {
		symbolsDomain = append(symbolsDomain, symbol.ToDomain())
	}
	return symbolsDomain, result.Error
}
