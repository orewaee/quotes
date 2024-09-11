package handlers

import (
	"github.com/orewaee/quotes/internal/app/api"
	"github.com/orewaee/quotes/internal/dtos"
	"github.com/orewaee/quotes/internal/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
)

type RandomHandler struct {
	quoteApi api.QuoteApi
	log      *zerolog.Logger
}

func NewRandomHandler(quoteApi api.QuoteApi, log *zerolog.Logger) *RandomHandler {
	return &RandomHandler{
		quoteApi: quoteApi,
		log:      log,
	}
}

func (handler *RandomHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	quote, err := handler.quoteApi.GetRandomQuote()
	if err != nil {
		log.Error().Err(err).Send()
		utils.MustWriteError(writer, err.Error(), http.StatusBadRequest)
		return
	}

	dtoQuote := &dtos.Quote{
		Text:   quote.Text,
		Author: quote.Author,
	}

	utils.MustWriteJson(writer, dtoQuote, http.StatusOK)
}
