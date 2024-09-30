package chart

import (
	"fmt"
	"strings"
	"time"

	"github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/pkg/utils/build_chart"
	"github.com/chartcmd/chart/pkg/utils/fetch/crypto"
	"github.com/chartcmd/chart/types"
)

func parseCandleSticks(candles [][]float64) []types.Candle {
	var result []types.Candle
	var numCandles = len(candles)
	for i := range candles {
		candle := candles[numCandles-1-i]
		result = append(result, types.Candle{
			Time:    time.Unix(int64(candle[0]), 0).Add(constants.TimeDiffUTC * time.Hour),
			Low:     candle[1],
			High:    candle[2],
			Open:    candle[3],
			Close:   candle[4],
			IsGreen: candle[3] < candle[4],
		})
	}
	return result
}

func DrawChart(ticker string, interval string) error {
	granularity := constants.IntervalToGranularity[interval]
	end := time.Now()
	start := end.Add(-time.Duration(granularity*constants.NumCandles) * time.Second)

	data, err := crypto.GetCoinbaseCandlestick(ticker+"-USD", start, end, granularity)
	if err != nil {
		return err
	}

	candles := parseCandleSticks(data)
	latestPrice := candles[len(candles)-1].Close
	chart := build_chart.BuildChart(candles)

	display(ticker, latestPrice, chart)
	return nil
}

/**
func DrawChart(ticker string,  string, stream bool) error {
	if stream:
		drawChartStream(ticker, granularity)

	else:
		drawChart(ticker, granularity)
}

func getCryptoCandles(granularity) {
	if c.NumCandles <= c.CoinbaseCandleMax {}
		end := time.Now()
		start := end.Add(-time.Duration(granularity*constants.NumCandles) * time.Second)
		data, err := crypto.GetCoinbase(ticker+"-USD", start, end, granularity)
		if err != nil {
			return nil, err
		}
		return parseCandleStick(data), nil
	} else {
		candles []types.Candle
		newGranularity = c.IntervalToGranularity["1d"]
		for i := range int(granularity/newGranularity) {
			end := time.Now().Add(i*(-newGranularity) * time.Second)
			start :=  end.Add(-time.Duration(newGranularity*c.CoinbaseCandleMax*i) * time.Second)

			data, err := crypto.GetCoinbase(ticker+"-USD", start, end, granularity)
			if err != nil {
				return nil, err
			}
			candles = insert(candles, 0, parseCandleStick(data))
		}
		return parseCandleStick(data), nil
	}
}

func drawChartStream((ticker string, interval string) {

	intervals = c.Intervals
	interValIdx = indexOf(interval, intervals)

	for ticks in ticker.(granlarity*time.Second)
		granularity := constants.IntervalToGranularity[interval]
		candles := getCandles(granularity)

		curCandle := initCurCandle(candles[-1].Close)
		candles = append(candles, curCandle)[1:]
		for ticks in ticker.c.StreamRefreshRate * time.MS
			latestPrice := crypto.GetCoinbaseLatest(ticker)
			curCandle = updateCurCandle(latestPrice)
			candles[-1] = curCandle

			chart := build_chart.BuildChart(candles)
			display(ticker, latestPrice, chart, pctChange)
			input = input()
			if left arrow => index = Max(0, index-1)
			if right arrow => index = Min(len(intervals)-1, idx+1)
				interval = intervals[idx]
				break
}

func drawChart(ticker string, interval string) {
	granularity := constants.IntervalToGranularity[interval]
	candles := getCandles(granularity)
	latestPrice := crypto.GetCoinbaseLatest(ticker)
	curCandle = initCurCandle(latestPrice)
	candles = append(candles, curCandle)[1:]

	chart := build_chart.BuildChart(candles)
	display(ticker, latestPrice, chart, pctChange)
}

*/

func display(ticker string, latestPrice float64, chart string) {
	utils.ClearScreen()
	ticker = strings.ToUpper(ticker)
	// TODO: print UTC time on top right hh:mm:ss
	if latestPrice < 0.1 {
		fmt.Printf("\n%*s%s: $%.8f\n", 4, "", constants.UpColorBold+ticker, latestPrice)
	} else {
		fmt.Printf("\n%*s%s: $%.2f\n", 4, "", constants.UpColorBold+ticker, latestPrice)
	}

	fmt.Println(chart + constants.ResetColor)
	fmt.Println()

	// TODO
	// should be 2 lines below bottom of chart
	// display timeframe bar and highlight selected interval (underlin and bolde it or sum, gray the other ones)

}
