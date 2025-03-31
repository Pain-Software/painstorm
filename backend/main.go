package main

import (
	"log/slog"
	"net"
	"net/http"
	"painstorm/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("Painstorm backend has started.")
	router := gin.Default()

	// Enable all origins
	router.Use(cors.Default())
	api := router.Group("/api")
	cfg := config.Get()

	api.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "v0.0.1",
		})
	})

	router.Run(net.JoinHostPort(cfg.App.Host, cfg.App.Port))
}
