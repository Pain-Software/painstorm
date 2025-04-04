package repository

import (
	"painstorm/model"
	"time"
)

type WeatherRepository interface {
	AddMeasurement(measurement model.Measurement)
	GetCityByName(cityName string) (model.City, error)
	GetCityByCoords(lat float64, lng float64) (model.City, error)
	GetMissingDates(city model.City, from int64, to int64) []time.Time
	GetCitiesWithIntensity(from int64, to int64, intensity float64) []model.City
}
