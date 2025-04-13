package helper

import (
	"painstorm/data"
	"painstorm/model"
	"time"
)

// MapWeatherDataToMeasurement maps WeatherData from OWM API to measurement model
func MapWeatherDataToMeasurement(wd data.WeatherData, cityID int64) model.Measurement {
	weathers := make([]model.MeasurementWeather, 0)
	for _, weather := range wd.Weather {
		weathers = append(weathers, model.MeasurementWeather{
			ID:          int64(weather.ID),
			Main:        weather.Main,
			Description: weather.Description,
			Icon:        weather.Icon,
		})
	}
	measurement := model.Measurement{
		CityID:    cityID,
		Timestamp: time.Unix(wd.Dt, 0),
		Weather:   weathers,
	}

	measurement.MinTemp = wd.Main.TempMin
	measurement.MaxTemp = wd.Main.TempMax
	measurement.Temperature = wd.Main.Temp
	measurement.Humidity = wd.Main.Humidity
	measurement.Pressure = float64(wd.Main.Pressure)
	measurement.Clouds = wd.Clouds.All

	measurement.WindSpeed = wd.Wind.Speed
	measurement.WindDegrees = float64(wd.Wind.Deg)

	if wd.Rain != nil {
		measurement.RainIntensity = wd.Rain.OneHour
	}
	return measurement
}
