package controller

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type WeatherController struct {
}

func NewWeatherController() *WeatherController {
	return &WeatherController{}
}

func (controller *WeatherController) Version(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version": "v0.0.1",
	})
	slog.Info("API version sent")
}

func (controller *WeatherController) Current(ctx *gin.Context) {
	place := ctx.Query("place")
	lat := ctx.Query("lat") // Get latitude
	lon := ctx.Query("lon")
	count := ctx.DefaultQuery("n", "1")

	if place == "" && (lat == "" || lon == "") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Either place or lat and lon is required"})
		return
	}

	responseData := make(map[string]string)
	if place != "" {
		responseData["place"] = place
		responseData["n"] = count
	} else {
		responseData["lat"] = lat
		responseData["lon"] = lon
		responseData["n"] = count
	}

	webResponse := Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   responseData,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
	slog.Info("Current weather downloaded")
}
