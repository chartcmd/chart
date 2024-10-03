package chart

import (
	"fmt"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/pkg/utils/fetch/crypto"
	"github.com/chartcmd/chart/pkg/utils/fetch/stocks"
)

func GetLatest(ticker string) (float64, error) {
	if utils.StrSliceContains(c.CryptoList, ticker, false) {
		return crypto.GetCoinbaseLatest(ticker + "-USD")
	} else if stocks.IsValidTicker(ticker) {
		return stocks.GetYFLatest(ticker)
	}
	return 0.0, fmt.Errorf("error: error fetching latest")
}
