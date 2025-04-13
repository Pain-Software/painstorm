package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"painstorm/config"
	"painstorm/controller"
	_ "painstorm/docs"
	"painstorm/repository"
	"painstorm/router"
	"painstorm/service"

	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	slog.Info("Painstorm backend has started.")

	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	credential := options.Credential{
		Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	}

	clientOpts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(credential)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	db := client.Database("painstorm")
	measurementCollection := db.Collection("measurements")

	client.Ping(ctx, nil)

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
	weatherRepository := repository.NewWeatherRepositoryImpl(conn, measurementCollection)

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
