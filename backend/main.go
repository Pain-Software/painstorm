package main

import (
	"log/slog"
	"net/http"
	"painstorm/config"
	"painstorm/controller"
	"painstorm/router"
)

func main() {
	slog.Info("Painstorm backend has started.")

	// Enable all origins
	cfg := config.Get()

	// Controller
	weatherController := controller.NewWeatherController()

	// Router
	routes := router.NewRouter(weatherController)

	// Server
	server := &http.Server{
		Handler: routes,
		Addr:    cfg.App.Host + ":" + cfg.App.Port,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
