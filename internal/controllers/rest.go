package controllers

import (
	"context"
	"github.com/orewaee/quotes/internal/app/api"
	"github.com/orewaee/quotes/internal/dtos"
	"github.com/orewaee/quotes/internal/utils"
	"net/http"
	"time"
)

type RestController struct {
	server   *http.Server
	quoteApi api.QuoteApi
}

func NewRestController(addr string, quoteApi api.QuoteApi) *RestController {
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
	}
}

func (controller *RestController) Run() error {
	return controller.server.ListenAndServe()
}

func (controller *RestController) Shutdown(ctx context.Context) error {
	return controller.server.Shutdown(ctx)
}
