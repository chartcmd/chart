package chart

import (
	"fmt"
	"strings"
	"time"

	"github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/utils/build_chart"
	"github.com/chartcmd/chart/pkg/utils/fetch/crypto"
	"github.com/chartcmd/chart/types"
)

func parseCandleStick(candles [][]float64) []types.Candle {
	var result []types.Candle
	var numCandles = len(candles)
	for i := range candles {
		candle := candles[numCandles-1-i]
		result = append(result, types.Candle{
			Time:  time.Unix(int64(candle[0]), 0).Add(7 * time.Hour),
			Low:   candle[1],
			High:  candle[2],
			Open:  candle[3],
			Close: candle[4],
		})
	}
	return result
}

func DrawChart(ticker string, interval string) error {
	granularity := constants.IntervalToGranularity[interval]
	end := time.Now()
	start := end.Add(-time.Duration(granularity*constants.NumCandles) * time.Second)

	data, err := crypto.GetCoinbase(ticker+"-USD", start, end, granularity)
	if err != nil {
		return err
	}

	candles := parseCandleStick(data)
	latestPrice := candles[0].Close
	chart, leftPadding := build_chart.BuildChart(candles, int(end.Sub(start).Minutes()))

	display(ticker, latestPrice, chart, leftPadding)
	return nil
}

func display(ticker string, latestPrice float64, chart string, leftPadding int) {
	constants.ClearScreen()
	ticker = strings.ToUpper(ticker)
	fmt.Printf("\n%*s%s: $%.2f\n\n", leftPadding-1, "", constants.UpColor+ticker, latestPrice)
	fmt.Println(chart + constants.ResetColor)
	fmt.Println()
}
