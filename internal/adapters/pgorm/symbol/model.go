package symbol

import (
	"stregy/internal/domain/symbol"
)

type Symbol struct {
	Name string `gorm:"primaryKey;type:string;"`
}

func (s Symbol) ToDomain() symbol.Symbol {
	return symbol.Symbol{Name: s.Name}
}
