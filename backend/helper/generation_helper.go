package helper

import (
	"math/rand"
	"painstorm/model"
	"time"
)

var weatherConditions = []model.MeasurementWeather{
	{200, "Thunderstorm", "Thunderstorm with light rain", "11d"},
	{201, "Thunderstorm", "Thunderstorm with rain", "11d"},
	{300, "Drizzle", "Light intensity drizzle", "09d"},
	{500, "Rain", "Light rain", "10d"},
	{600, "Snow", "Light snow", "13d"},
	{701, "Mist", "Mist", "50d"},
	{800, "Clear", "Clear sky", "01d"},
	{801, "Clouds", "Few clouds: 11-25%", "02d"},
	{802, "Clouds", "Scattered clouds: 25-50%", "03d"},
	{803, "Clouds", "Broken clouds: 51-84%", "04d"},
	{804, "Clouds", "Overcast clouds: 85-100%", "04d"},
}

// GenerateRandomMeasurements generates random measurements for testing purposes
func GenerateRandomMeasurements(n int) []model.Measurement {
	measurements := make([]model.Measurement, n)

	for i := 0; i < n; i++ {
		measurements[i] = model.Measurement{
			ID:            int64(i + 1),
			Timestamp:     time.Now().Add(time.Duration(-rand.Intn(1000)) * time.Hour),
			CityID:        int64(rand.Intn(100) + 800),
			MinTemp:       rand.Float64()*15 + 5,
			MaxTemp:       rand.Float64()*10 + 20,
			Temperature:   rand.Float64()*25 + 5,
			Humidity:      rand.Intn(101),
			Pressure:      rand.Float64()*50 + 950,
			SeaLevel:      rand.Float64()*5 + 100,
			GroundLevel:   rand.Float64()*10 + 50,
			WindSpeed:     rand.Float64() * 15,
			WindDegrees:   rand.Float64() * 360,
			RainIntensity: rand.Float64() * 20,
			Weather:       getRandomWeather(rand.Intn(3)),
		}
	}

	return measurements
}

// getRandomWeather generates random weather conditions
func getRandomWeather(num int) []model.MeasurementWeather {
	var randomWeatherList []model.MeasurementWeather

	for i := 0; i < num; i++ {
		randomIndex := rand.Intn(len(weatherConditions))
		randomWeatherList = append(randomWeatherList, weatherConditions[randomIndex])
	}

	return randomWeatherList
}
