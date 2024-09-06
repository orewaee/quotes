package main

import (
	"github.com/orewaee/quotes/internal/controllers"
	"github.com/orewaee/quotes/internal/disk"
	"github.com/orewaee/quotes/internal/services"
)

func main() {
	quoteRepo := disk.NewQuoteRepo()
	quoteApi := services.NewQuoteService(quoteRepo)

	controller := controllers.NewRestController(":8080", quoteApi)

	if err := controller.Run(); err != nil {
		panic(err)
	}
}
