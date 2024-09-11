package controllers

import (
	"context"
	"github.com/orewaee/quotes/internal/app/api"
	"github.com/orewaee/quotes/internal/handlers"
	"github.com/orewaee/quotes/internal/middlewares"
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

type RestController struct {
	server   *http.Server
	quoteApi api.QuoteApi
	log      *zerolog.Logger
}

func NewRestController(addr string, quoteApi api.QuoteApi, log *zerolog.Logger) *RestController {
	mux := http.NewServeMux()

	randomHandler := handlers.NewRandomHandler(quoteApi, log)
	mux.Handle("GET /random", middlewares.LogMiddleware(randomHandler, log))

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	return &RestController{
		server:   server,
		quoteApi: quoteApi,
		log:      log,
	}
}

func (controller *RestController) Run() error {
	controller.log.Info().Msgf("running app at addr %s", controller.server.Addr)
	return controller.server.ListenAndServe()
}

func (controller *RestController) Shutdown(ctx context.Context) error {
	controller.log.Info().Msg("shutting down the app...")
	return controller.server.Shutdown(ctx)
}
