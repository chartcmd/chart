package chart

import (
	"fmt"
	"math"
	"strings"
	"time"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/pkg/utils/fetch/crypto"
	"github.com/chartcmd/chart/pkg/utils/fetch/stocks"
	"github.com/chartcmd/chart/types"
)

func getCandles(ticker, interval string) ([]types.Candle, error) {
	granularity := c.IntervalToGranularity[interval]
	if utils.StrSliceContains(c.CryptoList, ticker, false) {
		return getCryptoCandles(ticker, granularity)
	} else if stocks.IsValidTicker(ticker) {
		if strings.ToUpper(interval) != "1D" {
			return nil, fmt.Errorf("error: only 1d interval available for stocks")
		}
		return getStockCandles(ticker, interval)
	}

	return nil, fmt.Errorf("unknown error in getCandles")
}

func getCryptoCandles(ticker string, granularity uint32) ([]types.Candle, error) {
	end := time.Now()
	totalDuration := time.Duration(granularity*c.NumCandles) * time.Second
	start := end.Add(-totalDuration)

	if c.NumCandles <= c.CoinbaseCandleMax {
		data, err := crypto.GetCoinbaseCandlestick(ticker+"-USD", start, end, granularity)
		if err != nil {
			return nil, err
		}
		return parseCoinbaseCandleSticks(data), nil
	} else {
		var candles []types.Candle
		remainingCandles := c.NumCandles
		currentStart := start

		for remainingCandles > 0 {
			batchSize := uint32(math.Min(float64(remainingCandles), float64(c.CoinbaseCandleMax)))
			batchEnd := currentStart.Add(time.Duration(granularity*batchSize) * time.Second)

			if batchEnd.After(end) {
				batchEnd = end
			}

			data, err := crypto.GetCoinbaseCandlestick(ticker+"-USD", currentStart, batchEnd, granularity)
			if err != nil {
				return nil, err
			}
			candles = append(candles, parseCoinbaseCandleSticks(data)...)

			remainingCandles -= batchSize
			currentStart = batchEnd
		}

		return candles, nil
	}
}

func getStockCandles(ticker, interval string) ([]types.Candle, error) {
	return stocks.GetYFCandleStick(ticker, interval)
}
