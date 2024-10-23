package models

type City struct {
	Name string `json:"name"`
}

type Temp struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type FeelsLike struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type ForecastListItem struct {
	Temp      Temp      `json:"temp"`
	FeelsLike FeelsLike `json:"feels_like"`
	Pressure  float64   `json:"pressure"`
	Humidity  float64   `json:"humidity"`
	Weather   Weather   `json:"weather"`
	WSpeed    float64   `json:"speed"`
	WDeg      float64   `json:"deg"`
	Cloud     float64   `json:"clouds"`
	Rain      float64   `json:"rain"`
}

type Forecast struct {
	City    City   `json:"city"`
	Country string `json:"country"`
}
