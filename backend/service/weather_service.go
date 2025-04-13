package service

import (
	"painstorm/model"
)

type WeatherService interface {
	CheckData(city model.City, from int64, to int64) (model.City, error)
	GenerateData(count int)
	RetrieveMeasurements(city model.City, from int64, to int64) ([]model.Measurement, error)
	RainIntensity(from int64, to int64, intensity float64) []model.City
	TempDiff(date string) []model.CityWithTempDiff
	StableWeather(city model.City, from string, to string, weatherType string) ([][]model.Measurement, error)
}
