package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
	"os"
	"painstorm/config"
	"painstorm/controller"
	"painstorm/repository"
	"painstorm/router"
	"painstorm/service"
)

func main() {
	slog.Info("Painstorm backend has started.")

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Enable all origins
	cfg := config.Get()

	// Repository
	connURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"))
	conn, err := pgxpool.New(context.Background(), connURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	weatherRepository := repository.NewWeatherRepositoryImpl(conn)

	// Service
	weatherService := service.NewWeatherServiceImpl(weatherRepository)

	// Controller
	weatherController := controller.NewWeatherController(weatherService)

	// Router
	routes := router.NewRouter(weatherController)

	// Server
	server := &http.Server{
		Handler: routes,
		Addr:    cfg.App.Host + ":" + cfg.App.Port,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
