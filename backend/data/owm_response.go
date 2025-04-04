package data

import (
	"encoding/json"
	"log/slog"
)

func FormatWeatherResponse(data WeatherResponse) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		slog.Error("Failed to convert response to JSON: %v", err)
	}
	return string(jsonData)
}

type WeatherResponse struct {
	Message  string        `json:"message"`
	Code     string        `json:"cod"`
	CityID   int           `json:"city_id"`
	CalcTime float64       `json:"calctime"`
	Cnt      int           `json:"cnt"`
	List     []WeatherData `json:"list"`
}

type WeatherData struct {
	Dt      int64         `json:"dt"`
	Main    MainData      `json:"main"`
	Wind    WindData      `json:"wind"`
	Clouds  CloudData     `json:"clouds"`
	Weather []WeatherDesc `json:"weather"`
	Rain    *RainData     `json:"rain,omitempty"`
}

type MainData struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
}

type WindData struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type CloudData struct {
	All int `json:"all"`
}

type WeatherDesc struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type RainData struct {
	OneHour float64 `json:"1h"`
}
