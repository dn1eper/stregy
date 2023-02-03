package composites

import (
	symbol1 "stregy/internal/adapters/pgorm/symbol"
	"stregy/internal/domain/symbol"
)

type SymbolComposite struct {
	Service symbol.Service
}

func NewSymbolComposite(composite *PGormComposite) (*SymbolComposite, error) {
	repository := symbol1.NewRepository(composite.db)

	return &SymbolComposite{
		Service: symbol.NewService(repository),
	}, nil
}
