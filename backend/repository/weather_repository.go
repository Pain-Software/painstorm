package repository

import (
	"painstorm/model"
	"time"
)

type WeatherRepository interface {
	AddMeasurement(measurement model.Measurement)
	GetCityByName(cityName string) (model.City, error)
	GetCityByCoords(lat float64, lng float64) (model.City, error)
	GetMeasurements(city model.City, from int64, to int64) ([]model.Measurement, error)
	GetWeatherByMeasurementID(measurementID int64) ([]model.MeasurementWeather, error)
	GetMissingDates(city model.City, from int64, to int64) []time.Time
	GetCitiesWithIntensity(from int64, to int64, intensity float64) []model.City
	GetCitiesWithMaxTempDiff(date string) []model.CityWithTempDiff
	GetStableData(city model.City, from string, to string, weatherType string) ([][]model.Measurement, error)
}
