package controller

import (
	"log/slog"
	"net/http"
	"painstorm/data"
	"painstorm/model"
	"painstorm/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

// Retrieve 		godoc
// @Summary			Retrieve weather measurements
// @Description		Retrieves weather measurements for a given city or coordinates within a specified time interval
// @Tags			weather
// @Accept			json
// @Produce			application/json
// @Param			place query string false "City name"
// @Param			latitude query number false "Latitude of the location"
// @Param			longitude query number false "Longitude of the location"
// @Param			from query string true "Start date in YYYY-MM-DD format"
// @Param			to query string true "End date in YYYY-MM-DD format"
// @Success 		200 {array} model.Measurement
// @Router			/retrieve [get]
func (controller *WeatherController) Retrieve(ctx *gin.Context) {
	// Retrieving params
	var req data.RetrieveRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Place == "" && (req.Latitude == 0 || req.Longitude == 0) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Either place or lat and lon is required"})
		return
	}

	city := model.City{
		Name:      req.Place,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
	// Downloading needed data
	city, err = controller.service.CheckData(city, req.From, req.To)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Querying for measurement data
	measurements, err := controller.service.RetrieveMeasurements(city, req.From, req.To)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, measurements)
	slog.Info("Weather data retrieved")
}

// Current 		godoc
// @Summary			Download and update recent weather data
// @Description		Downloads weather data for a location for the past N days (default 1 day)
// @Tags			weather
// @Accept			json
// @Produce			application/json
// @Param			place query string false "City name"
// @Param			latitude query number false "Latitude of the location"
// @Param			longitude query number false "Longitude of the location"
// @Param			count query int false "Number of days in the past to fetch data for (default is 1)"
// @Success			200 {array} model.Measurement
// @Router			/current [get]
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
	city, err = controller.service.CheckData(city, from, to)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Querying for measurement data
	measurements, err := controller.service.RetrieveMeasurements(city, from, to)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, measurements)
	slog.Info("Current weather updated")
}

// Generate 		godoc
// @Summary			Generate random weather data
// @Description		Generates and stores N random weather records (default is 1)
// @Tags			weather
// @Accept			json
// @Produce			application/json
// @Param			count query int false "Number of random records to generate (default is 1)"
// @Success			200 {object} interface{} "null response"
// @Router			/generate [get]
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

	ctx.JSON(http.StatusOK, nil)
	slog.Info("Generated " + strconv.Itoa(req.Count) + " weather records")
}

// RainIntensity 		godoc
// @Summary			Get cities by rain intensity
// @Description		Retrieves cities that experienced rain of the specified intensity within a time interval
// @Tags			weather
// @Accept			json
// @Produce			application/json
// @Param			from query string true "Start date in YYYY-MM-DD format"
// @Param			to query string true "End date in YYYY-MM-DD format"
// @Param			intensity query number true "Rain intensity threshold (e.g., mm/hour)"
// @Success			200 {array} model.City
// @Router			/rain-intensity [get]
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

	ctx.JSON(http.StatusOK, cities)
	slog.Info("Found cities with rain intensity")
}

// RainIntensity 		godoc
// @Summary			Get cities with max temperature difference LIMIT 3
// @Description		Retrieves cities that experienced the highest temperature difference on a given date
// @Tags			weather
// @Produce			application/json
// @Param			date query string true "Start date in YYYY-MM-DD format"
// @Success			200 {array} model.CityWithTempDiff
// @Router			/rain-intensity [get]
func (controller *WeatherController) TempDiff(ctx *gin.Context) {
	// Retrieving params
	var req data.TempDiffRequest
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Date == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Date is required"})
		return
	}

	// Parse the date string to a time.Time object
	// Assuming the date is in the format "YYYY-MM-DD"
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	cities := controller.service.TempDiff(parsedDate.Format("2006-01-02"))

	ctx.JSON(http.StatusOK, cities)
	slog.Info("Found cities with temp diff")
}
