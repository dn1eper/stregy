package composites

import (
	symbol1 "stregy/internal/adapters/pgorm/symbol"
	"stregy/internal/domain/symbol"
)

type SymbolComposite struct {
	Repository symbol.Repository
}

func NewSymbolComposite(composite *PGormComposite) (*SymbolComposite, error) {
	repository := symbol1.NewRepository(composite.db)

	return &SymbolComposite{
		Repository: repository,
	}, nil
}
