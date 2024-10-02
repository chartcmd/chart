package chart

import (
	"fmt"
	"math"
	"time"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/pkg/utils/fetch/crypto"
	"github.com/chartcmd/chart/types"
)

func getCryptoCandles(ticker string, granularity uint32) ([]types.Candle, error) {
	end := time.Now()
	totalDuration := time.Duration(granularity*c.NumCandles) * time.Second
	start := end.Add(-totalDuration)

	if c.NumCandles <= c.CoinbaseCandleMax {
		data, err := crypto.GetCoinbaseCandlestick(ticker+"-USD", start, end, granularity)
		if err != nil {
			return nil, err
		}
		return parseCandleSticks(data), nil
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
			candles = append(candles, parseCandleSticks(data)...)

			remainingCandles -= batchSize
			currentStart = batchEnd
		}

		return candles, nil
	}
}

func getCandles(ticker, interval string) ([]types.Candle, error) {
	granularity := c.IntervalToGranularity[interval]
	if utils.StrSliceContains(c.CryptoList, ticker, false) {
		return getCryptoCandles(ticker, granularity)
	}

	return nil, fmt.Errorf("unknown error in getCandles")
}
