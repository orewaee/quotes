package disk

import (
	"encoding/json"
	"github.com/orewaee/quotes/internal/app/domain"
	"github.com/orewaee/quotes/internal/models"
	"github.com/orewaee/quotes/internal/utils"
	"math/rand/v2"
	"os"
	"strings"
)

type QuoteRepo struct{}

func NewQuoteRepo() *QuoteRepo {
	return &QuoteRepo{}
}

func (repo *QuoteRepo) GetRandomQuote() (*domain.Quote, error) {
	entries, err := os.ReadDir("quotes")
	if err != nil {
		return nil, err
	}

	suitableEntries := utils.Filter(entries, func(entry os.DirEntry) bool {
		return strings.Contains(entry.Name(), ".json")
	})

	index := rand.IntN(len(suitableEntries))

	bytes, err := os.ReadFile("quotes/" + suitableEntries[index].Name())
	if err != nil {
		return nil, err
	}

	quote := new(models.Quote)
	if err := json.Unmarshal(bytes, quote); err != nil {
		return nil, err
	}

	return &domain.Quote{
		Text:   quote.Text,
		Author: quote.Author,
	}, nil
}
