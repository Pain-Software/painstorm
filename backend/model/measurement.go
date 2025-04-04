package model

import "time"

type Measurement struct {
	ID            int64                `json:"id"`
	Timestamp     time.Time            `json:"timestamp"`
	CityID        int64                `json:"city_id"`
	MinTemp       float64              `json:"min_temperature,omitempty"`
	MaxTemp       float64              `json:"max_temperature,omitempty"`
	Temperature   float64              `json:"temperature,omitempty"`
	Humidity      int                  `json:"humidity,omitempty"`
	Pressure      float64              `json:"pressure,omitempty"`
	SeaLevel      float64              `json:"sea_level,omitempty"`
	GroundLevel   float64              `json:"ground_level,omitempty"`
	WindSpeed     float64              `json:"wind_speed,omitempty"`
	WindDegrees   float64              `json:"wind_degrees,omitempty"`
	RainIntensity float64              `json:"rain_intensity,omitempty"`
	Weather       []MeasurementWeather `json:"weather"`
}

type MeasurementWeather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
