package data

type GenerateRequest struct {
	Count int `form:"count"`
}

type RainRequest struct {
	From      int64   `form:"from"`
	To        int64   `form:"to"`
	Intensity float64 `form:"intensity"`
}

type TempDiffRequest struct {
	Date string `form:"date"`
}

type StableWeatherRequest struct {
	Place       string  `form:"place"`
	Latitude    float64 `form:"lat"`
	Longitude   float64 `form:"lon"`
	From        string  `form:"from"`
	To          string  `form:"to"`
	WeatherType string  `form:"weatherType"`
}

type CurrentRequest struct {
	Place     string  `form:"place"`
	Latitude  float64 `form:"lat"`
	Longitude float64 `form:"lon"`
	Count     int     `form:"n"`
}

type RetrieveRequest struct {
	Place     string  `form:"place"`
	Latitude  float64 `form:"lat"`
	Longitude float64 `form:"lon"`
	From      int64   `form:"from"`
	To        int64   `form:"to"`
}
