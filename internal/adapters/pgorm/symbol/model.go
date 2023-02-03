package symbol

import (
	"stregy/internal/domain/symbol"
)

type Symbol struct {
	Name      string `gorm:"primaryKey;type:string"`
	Precision int    `gorm:"type:int"`
}

func (s *Symbol) ToDomain() *symbol.Symbol {
	res := symbol.Symbol(*s)
	return &res
}
