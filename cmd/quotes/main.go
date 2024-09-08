package main

import (
	"context"
	"errors"
	"flag"
	"github.com/orewaee/quotes/internal/controllers"
	"github.com/orewaee/quotes/internal/disk"
	"github.com/orewaee/quotes/internal/logger"
	"github.com/orewaee/quotes/internal/services"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	addr := ""

	flag.StringVar(&addr, "addr", ":8080", "--addr=:8081")
	flag.Parse()

	quoteRepo := disk.NewQuoteRepo()
	quoteApi := services.NewQuoteService(quoteRepo)

	log, err := logger.NewZerolog()
	if err != nil {
		panic(err)
	}

	controller := controllers.NewRestController(addr, quoteApi, log)

	stop := make(chan os.Signal, 1)

	go func() {
		signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-stop

		if err := controller.Shutdown(context.Background()); err != nil {
			log.Error().Err(err).Send()
		}
	}()

	log.Info().Msg("press ctrl+c to exit")

	if err := controller.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error().Err(err).Send()
	}
}
