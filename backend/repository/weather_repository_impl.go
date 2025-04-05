package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"painstorm/helper"
	"painstorm/model"
	"time"
)

type WeatherRepositoryImpl struct {
	DbConn *pgxpool.Pool
}

func NewWeatherRepositoryImpl(conn *pgxpool.Pool) WeatherRepository {
	return &WeatherRepositoryImpl{
		DbConn: conn,
	}
}

// Inserts

// AddMeasurement inserts a measurement to DB
func (w *WeatherRepositoryImpl) AddMeasurement(m model.Measurement) {
	// Insert into measurements
	query := "INSERT INTO Measurement (timestamp, id_city, min_temperature, max_temperature, temperature, humidity, pressure, sea_level, ground_level, wind_speed, wind_degrees, rain_intensity) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id"
	err := w.DbConn.QueryRow(context.Background(), query, m.Timestamp, m.CityID, m.MinTemp, m.MaxTemp, m.Temperature, m.Humidity, m.Pressure, m.SeaLevel, m.GroundLevel, m.WindSpeed, m.WindDegrees, m.RainIntensity).Scan(&m.ID)
	if err != nil {
		slog.Error("Error adding measurement: ", err)
		return
	}

	// Insert connection to weather
	for _, weather := range m.Weather {
		err = w.DbConn.QueryRow(context.Background(), "INSERT INTO weather_in_measurement (id_weather, id_measurement) VALUES ($1, $2) RETURNING id_measurement", weather.ID, m.ID).Scan(&m.ID)
		if err != nil {
			slog.Error("Error adding measurement: ", err)
		}
	}
}

// Queries

// GetCityByName retrieves a city using name
func (w *WeatherRepositoryImpl) GetCityByName(cityName string) (model.City, error) {
	var city model.City
	err := w.DbConn.QueryRow(context.Background(), "SELECT id, name, latitude, longitude FROM City WHERE name=$1", cityName).Scan(&city.ID, &city.Name, &city.Latitude, &city.Longitude)
	return city, err
}

// GetCityByCoords retrieves a city using coordinates
func (w *WeatherRepositoryImpl) GetCityByCoords(lat float64, lng float64) (model.City, error) {
	var city model.City
	err := w.DbConn.QueryRow(context.Background(), "SELECT id, name, latitude, longitude FROM City WHERE latitude=$1 AND longitude=$2", lat, lng).Scan(&city.ID, &city.Name, &city.Latitude, &city.Longitude)
	return city, err
}

// GetMeasurements retrieves measurements for a city in specified interval
func (w *WeatherRepositoryImpl) GetMeasurements(city model.City, from int64, to int64) ([]model.Measurement, error) {
	measurements := make([]model.Measurement, 0)

	query := "SELECT id, timestamp, id_city, min_temperature, max_temperature, temperature, humidity, pressure, sea_level, ground_level, wind_speed, wind_degrees, rain_intensity" +
		" FROM Measurement WHERE id_city = $1 AND timestamp BETWEEN $2 AND $3"
	rows, err := w.DbConn.Query(context.Background(), query, city.ID, helper.NormalizeToNoonUTC(time.Unix(from, 0)), helper.NormalizeToNoonUTC(time.Unix(to, 0)))
	if err != nil {
		slog.Error("Error getting measurements: ", err)
		return make([]model.Measurement, 0), err
	}
	defer rows.Close()

	for rows.Next() {
		var m model.Measurement
		err = rows.Scan(&m.ID, &m.Timestamp, &m.CityID, &m.MinTemp, &m.MaxTemp, &m.Temperature, &m.Humidity, &m.Pressure, &m.SeaLevel, &m.GroundLevel, &m.WindSpeed, &m.WindDegrees, &m.RainIntensity)
		if err != nil {
			slog.Error("Error getting measurements: ", err)
			continue
		}

		weatherList, err := w.GetWeatherByMeasurementID(m.ID)
		if err != nil {
			slog.Error("Error getting weather for measurement ID: ", m.ID, ":", err)
		} else {
			m.Weather = weatherList
		}

		measurements = append(measurements, m)
	}
	return measurements, nil
}

// GetWeatherByMeasurementID retrieves all weather for measurement
func (w *WeatherRepositoryImpl) GetWeatherByMeasurementID(measurementID int64) ([]model.MeasurementWeather, error) {
	query := `
		SELECT w.id, w.main, w.description, w.icon
		FROM WEATHER w
		JOIN WEATHER_IN_MEASUREMENT wim ON wim.id_weather = w.id
		WHERE wim.id_measurement = $1`

	rows, err := w.DbConn.Query(context.Background(), query, measurementID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	weatherList := make([]model.MeasurementWeather, 0)
	for rows.Next() {
		var weather model.MeasurementWeather
		if err := rows.Scan(&weather.ID, &weather.Main, &weather.Description, &weather.Icon); err != nil {
			continue
		}
		weatherList = append(weatherList, weather)
	}
	return weatherList, nil
}

// GetMissingDates retrieves missing dates for a city from an interval
func (w *WeatherRepositoryImpl) GetMissingDates(city model.City, from int64, to int64) []time.Time {
	dates := make([]time.Time, 0)

	query := "WITH date_series AS (SELECT generate_series($1::timestamp, $2::timestamp, '1 day') AS missing_date) SELECT ds.missing_date FROM date_series ds LEFT JOIN Measurement m ON ds.missing_date::date = DATE(m.timestamp) AND m.id_city = $3 WHERE m.timestamp IS NULL;"
	rows, err := w.DbConn.Query(context.Background(), query, helper.NormalizeToNoonUTC(time.Unix(from, 0)), helper.NormalizeToNoonUTC(time.Unix(to, 0)), city.ID)
	if err != nil {
		slog.Error("Error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var date time.Time
		err = rows.Scan(&date)
		if err != nil {
			slog.Error("Error scanning row: %v", err)
			continue
		}
		dates = append(dates, date)
	}
	return dates
}

// GetCitiesWithIntensity retrieves cities with specified intensity in interval
func (w *WeatherRepositoryImpl) GetCitiesWithIntensity(from int64, to int64, intensity float64) []model.City {
	cities := make([]model.City, 0)
	rows, err := w.DbConn.Query(context.Background(), "SELECT id, name, findName, country, longitude, latitude FROM CITY WHERE id in (SELECT id_city FROM MEASUREMENT WHERE TIMESTAMP BETWEEN $1 AND $2 AND rain_intensity = $3)", time.Unix(from, 0), time.Unix(to, 0), intensity)
	if err != nil {
		slog.Error("Error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var c model.City

		err = rows.Scan(&c.ID, &c.Name, &c.FindName, &c.Country, &c.Longitude, &c.Latitude)
		if err != nil {
			slog.Error("Error scanning row: %v", err)
			continue
		}

		cities = append(cities, c)
	}
	return cities
}
