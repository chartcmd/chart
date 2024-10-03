package chart

import (
	"math"
	"time"

	"github.com/chartcmd/chart/types"
)

func updateCurCandle(candle types.Candle, latestPrice float64) types.Candle {
	return types.Candle{
		Time:    time.Now(),
		Low:     math.Min(latestPrice, candle.Low),
		High:    math.Max(latestPrice, candle.High),
		Open:    candle.Open,
		Close:   latestPrice,
		IsGreen: latestPrice > candle.Open,
	}
}

func initCurCandle(latestPrice float64, lastCandle types.Candle) types.Candle {
	isGreen := latestPrice > lastCandle.Close
	var low, high float64
	if isGreen {
		low = lastCandle.Close
		high = latestPrice
	} else {
		low = latestPrice
		high = lastCandle.Close
	}

	return types.Candle{
		Time:    time.Now(),
		Low:     low,
		High:    high,
		Open:    lastCandle.Close,
		Close:   latestPrice,
		IsGreen: isGreen,
	}
}

func parseCandleSticks(candles [][]float64) []types.Candle {
	var result []types.Candle
	var numCandles = len(candles)
	for i := range candles {
		candle := candles[numCandles-1-i]
		result = append(result, types.Candle{
			Time:    time.Unix(int64(candle[0]), 0),
			Low:     candle[1],
			High:    candle[2],
			Open:    candle[3],
			Close:   candle[4],
			IsGreen: candle[3] < candle[4],
		})
	}
	return result
}
