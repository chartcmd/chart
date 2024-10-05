package stocks

import (
	"fmt"
)

func GetYFLatest(ticker string) (float64, error) {
	candles, err := GetYFCandleStick(ticker, "15m")
	if err != nil {
		return 0.0, fmt.Errorf("error fetching quote for %s: %v", ticker, err)
	}
	return candles[len(candles)-1].Close, nil
}
