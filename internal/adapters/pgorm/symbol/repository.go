package symbol

import (
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

func (r *repository) Exists(name string) bool {
	symbol := &Symbol{Name: name}
	result := r.db.First(symbol)
	if result.Error != nil {
		return false
	}
	return true
}

func (r *repository) Create(name string) (*symbol.Symbol, error) {
	symbol := &Symbol{Name: name}
	result := r.db.Create(symbol)
	if result.Error != nil {
		return nil, result.Error
	}
	tableName := strings.ToLower(name)
	r.db.Exec(fmt.Sprintf("CREATE TABLE %vs (LIKE quotes INCLUDING ALL);", tableName))
	symbolDomain := symbol.ToDomain()
	return &symbolDomain, result.Error
}

func (r *repository) GetAll() ([]symbol.Symbol, error) {
	symbols := make([]Symbol, 0)
	result := r.db.Find(&symbols)
	symbolsDomain := make([]symbol.Symbol, 0, len(symbols))
	for _, symbol := range symbols {
		symbolsDomain = append(symbolsDomain, symbol.ToDomain())
	}
	return symbolsDomain, result.Error
}
