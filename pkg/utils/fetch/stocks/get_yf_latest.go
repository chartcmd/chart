package stocks

import (
	"fmt"
	"strings"

	"github.com/piquette/finance-go/quote"
)

func GetYFLatest(ticker string) (float64, error) {
	if !StockMarketIsOpen() {
		candles, err := GetYFCandleStick(ticker, "15m")
		if err != nil {
			return 0.0, fmt.Errorf("error fetching quote for %s: %v", ticker, err)
		}
		return candles[len(candles)-1].Close, nil
	}
	q, err := quote.Get(strings.ToUpper(ticker))
	if err != nil {
		return 0.0, fmt.Errorf("error fetching quote for %s: %v", ticker, err)
	}

	if q == nil {
		return 0.0, fmt.Errorf("no quote data available for %s", ticker)
	}

	return q.RegularMarketPrice, nil
}
