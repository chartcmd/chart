package types

import "time"

type Candle struct {
	Time    time.Time
	Low     float64
	High    float64
	Open    float64
	Close   float64
	IsGreen bool
}

type Config struct {
	UpColor           string   `json:"up_color"`
	DownColor         string   `json:"down_color"`
	DefaultTimeFrame  string   `json:"default_tf"`
	EquitiesWatchlist []string `json:"equity_wl"`
	CryptoWatchlist   []string `json:"crypto_wl"`
}
