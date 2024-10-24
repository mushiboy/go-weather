package models

type Weather struct {
	Main string `json:"main"`
	Desc string `json:"description"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Rain struct {
	OneHour float64 `json:"1h"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type CurrentWeather struct {
	Weather []Weather `json:"weather"`
	Basic   Main      `json:"main"`
	Wind    Wind      `json:"wind"`
	Rain    Rain      `json:"rain"`
	Cloud   Clouds    `json:"clouds"`
	Sys     Sys       `json:"sys"`
}
