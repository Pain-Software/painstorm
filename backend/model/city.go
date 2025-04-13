package model

type City struct {
	ID        int64   `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	FindName  string  `json:"find_name,omitempty" db:"findName"`
	Country   string  `json:"country,omitempty" db:"country"`
	Longitude float64 `json:"longitude,omitempty" db:"longitude"`
	Latitude  float64 `json:"latitude,omitempty" db:"latitude"`
}

type CityWithTempDiff struct {
	City
	TempDiff *float64 `json:"tempDiff,omitempty"`
}
