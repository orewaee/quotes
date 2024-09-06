package api

import "github.com/orewaee/quotes/internal/app/domain"

type QuoteApi interface {
	GetRandomQuote() (*domain.Quote, error)
}
