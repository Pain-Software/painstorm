package controller

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"painstorm/data"
	"painstorm/model"
	"painstorm/service"
	"strconv"
	"time"
)

type WeatherController struct {
	service service.WeatherService
}

func NewWeatherController(weatherService service.WeatherService) *WeatherController {
	return &WeatherController{
		service: weatherService,
	}
}

// Version returns the current version of the backend
func (controller *WeatherController) Version(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"version": "v0.1.0",
	})
	slog.Info("API version sent")
}

// Current downloads data for location for n days in the past
func (controller *WeatherController) Current(ctx *gin.Context) {
	// Retrieving params
	var req data.CurrentRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Place == "" && (req.Latitude == 0 || req.Longitude == 0) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Either place or lat and lon is required"})
		return
	}
	if req.Count == 0 {
		req.Count = 1
	}

	dCount := time.Duration(req.Count)
	from := time.Now().Add(dCount * -24 * time.Hour).Unix()
	to := time.Now().Unix()
	city := model.City{
		Name:      req.Place,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	// Downloading needed data
	err = controller.service.CheckData(city, from, to)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond
	webResponse := data.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
	slog.Info("Current weather updated")
}

// Generate generates n random records
func (controller *WeatherController) Generate(ctx *gin.Context) {
	// Retrieving params
	var req data.GenerateRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Count == 0 {
		req.Count = 1
	}

	// Generate and insert data
	controller.service.GenerateData(req.Count)

	// Respond
	webResponse := data.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
	slog.Info("Generated " + strconv.Itoa(req.Count) + " weather records")
}

// RainIntensity retrieves cities with rain of specified intensity in specified time interval
func (controller *WeatherController) RainIntensity(ctx *gin.Context) {
	// Retrieving params
	var req data.RainRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Query DB
	cities := controller.service.RainIntensity(req.From, req.To, req.Intensity)

	// Respond
	webResponse := data.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   cities,
	}
	ctx.JSON(http.StatusOK, webResponse)
	slog.Info("Found cities with rain intensity")
}
