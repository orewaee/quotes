package main

import (
	"flag"
	"github.com/orewaee/quotes/internal/controllers"
	"github.com/orewaee/quotes/internal/disk"
	"github.com/orewaee/quotes/internal/services"
)

func main() {
	addr := ""

	flag.StringVar(&addr, "addr", ":8080", "--addr=:8081")
	flag.Parse()

	quoteRepo := disk.NewQuoteRepo()
	quoteApi := services.NewQuoteService(quoteRepo)

	controller := controllers.NewRestController(addr, quoteApi)

	if err := controller.Run(); err != nil {
		panic(err)
	}
}
