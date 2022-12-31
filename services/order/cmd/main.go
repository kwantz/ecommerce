package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kwantz/ecommerce/services/order/internal/controller"
)

func main() {
	pingController := controller.NewPingController()

	router := chi.NewRouter()

	setupRouting(router, Controller{
		Ping: pingController,
	})

	log.Println("server starting at :8080")
	http.ListenAndServe(":8080", router)
}

func setupRouting(router *chi.Mux, controller Controller) {
	router.Get("/ping", controller.Ping.PingHandler)
}

type Controller struct {
	Ping *controller.PingController
}
