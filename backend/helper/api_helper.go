package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"painstorm/data"
)

const baseURL = "https://history.openweathermap.org/data/2.5/history/city"

// GetApiDataByCoord fetches historical data from the OWM API using coordinates and a date interval
func GetApiDataByCoord(lon float64, lat float64, start int64, end int64) data.WeatherResponse {
	apiKey := os.Getenv("OWM_API_KEY")
	if apiKey == "" {
		log.Fatalln("API key is missing")
	}

	// Construct the request URL
	url := fmt.Sprintf("%s?lat=%.4f&lon=%.4f&type=hour&start=%d&end=%d&appid=%s&units=metric",
		baseURL, lat, lon, start, end, apiKey)

	// Send the request
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// Parse the request response
	var weatherData data.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		slog.Error("Failed to parse JSON: %v", err)
	}
	return weatherData
}
