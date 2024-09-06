package repos

import "github.com/orewaee/quotes/internal/app/domain"

type QuoteRepo interface {
	GetRandomQuote() (*domain.Quote, error)
}
