package service

import (
	"log/slog"
	"painstorm/helper"
	"painstorm/model"
	"painstorm/repository"
	"slices"
	"time"
)

type WeatherServiceImpl struct {
	Repository repository.WeatherRepository
}

func NewWeatherServiceImpl(weatherRepository repository.WeatherRepository) WeatherService {
	return &WeatherServiceImpl{
		Repository: weatherRepository,
	}
}

// CheckData ensures requested data is stored in DB
func (service *WeatherServiceImpl) CheckData(city model.City, from int64, to int64) (model.City, error) {
	// Get city from DB using name or coords
	var foundCity model.City
	var err error
	if city.Name != "" {
		foundCity, err = service.Repository.GetCityByName(city.Name)
	} else {
		foundCity, err = service.Repository.GetCityByCoords(city.Latitude, city.Longitude)
	}
	if err != nil {
		slog.Error("City query failed: ", err)
		return foundCity, err
	}

	// If no dates missing, return
	missingDates := service.Repository.GetMissingDates(foundCity, from, to)
	if len(missingDates) <= 0 {
		return foundCity, nil
	}

	// Otherwise update db using API
	// Takes measurements at 12pm from location and saves it (also time really sucks in golang)
	response := helper.GetApiDataByCoord(foundCity.Longitude, foundCity.Latitude, from, to)
	for _, measurement := range response.List {
		if time.Unix(measurement.Dt, 0).Hour() != 12 {
			continue
		}
		measurement.Dt = helper.NormalizeToNoonUTC(time.Unix(measurement.Dt, 0)).Unix()
		if slices.Contains(missingDates, time.Unix(measurement.Dt, 0).UTC()) {
			ms := helper.MapWeatherDataToMeasurement(measurement, foundCity.ID)
			service.Repository.AddMeasurement(ms)
			// TODO implement saving to mongo here using measurement
		}
	}
	return foundCity, nil
}

// GenerateData generates random data for testing purposes
func (service *WeatherServiceImpl) GenerateData(count int) {
	// Generate measurements
	measurements := helper.GenerateRandomMeasurements(count)

	// Save measurements to db
	for _, measurement := range measurements {
		service.Repository.AddMeasurement(measurement)
	}
}

func (service *WeatherServiceImpl) RetrieveMeasurements(city model.City, from int64, to int64) ([]model.Measurement, error) {
	return service.Repository.GetMeasurements(city, from, to)
}

// RainIntensity finds cities with specified rain intensity in specified date interval
func (service *WeatherServiceImpl) RainIntensity(from int64, to int64, intensity float64) []model.City {
	// Query the DB
	return service.Repository.GetCitiesWithIntensity(from, to, intensity)
}

// TempDiff retrieves cities with the highest temp difference on specified date
func (service *WeatherServiceImpl) TempDiff(date int64) []model.City {
	//TODO implement me (Dont forget to use helper.NormalizeToNoonUTC() for dates)
	panic("implement me")
}

// StableWeather retrieves dates with constant weather in interval
func (service *WeatherServiceImpl) StableWeather(city model.City, from int64, to int64, weatherType string) []time.Time {
	//TODO implement me (Dont forget to use helper.NormalizeToNoonUTC() for dates)
	panic("implement me")
}
