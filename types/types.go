package types

import "time"

type Candle struct {
	Time  time.Time
	Low   float64
	High  float64
	Open  float64
	Close float64
}

type Config struct {
	ChartColors struct {
		Up   string `json:"up"`
		Down string `json:"down"`
	} `json:"chart_colors"`
	ListCmdTickers struct {
		Equities []string `json:"equities"`
		Crypto   []string `json:"crypto"`
	} `json:"list_cmd_tickers"`
}
