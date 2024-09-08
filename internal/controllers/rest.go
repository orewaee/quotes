package controllers

import (
	"context"
	"github.com/orewaee/quotes/internal/app/api"
	"github.com/orewaee/quotes/internal/dtos"
	"github.com/orewaee/quotes/internal/utils"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type RestController struct {
	server   *http.Server
	quoteApi api.QuoteApi
	logger   *zerolog.Logger
}

func NewRestController(addr string, quoteApi api.QuoteApi, logger *zerolog.Logger) *RestController {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /random", func(writer http.ResponseWriter, request *http.Request) {
		quote, err := quoteApi.GetRandomQuote()
		if err != nil {
			utils.MustWriteError(writer, err.Error(), http.StatusBadRequest)
			return
		}

		dtoQuote := &dtos.Quote{
			Text:   quote.Text,
			Author: quote.Author,
		}

		utils.MustWriteJson(writer, dtoQuote, http.StatusOK)
	})

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	return &RestController{
		server:   server,
		quoteApi: quoteApi,
		logger:   logger,
	}
}

func (controller *RestController) Run() error {
	controller.logger.Info().Msgf("running app at addr %s", controller.server.Addr)
	return controller.server.ListenAndServe()
}

func (controller *RestController) Shutdown(ctx context.Context) error {
	controller.logger.Info().Msg("shutting down the app...")
	return controller.server.Shutdown(ctx)
}
