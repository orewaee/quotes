package services

import (
	"github.com/orewaee/quotes/internal/app/domain"
	"github.com/orewaee/quotes/internal/app/repos"
)

type QuoteService struct {
	quoteRepo repos.QuoteRepo
}

func NewQuoteService(quoteRepo repos.QuoteRepo) *QuoteService {
	return &QuoteService{
		quoteRepo: quoteRepo,
	}
}

func (service *QuoteService) GetRandomQuote() (*domain.Quote, error) {
	return service.quoteRepo.GetRandomQuote()
}
