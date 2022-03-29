package composites

import (
	quote1 "stregy/internal/adapters/pgorm/quote"
	"stregy/internal/domain/quote"
)

type QuoteComposite struct {
	Repository quote.Repository
}

func NewQuoteComposite(composite *PGormComposite) (*QuoteComposite, error) {
	repository := quote1.NewRepository(composite.db)

	return &QuoteComposite{
		Repository: repository,
	}, nil
}
