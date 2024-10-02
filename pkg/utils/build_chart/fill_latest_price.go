package build_chart

import (
	c "github.com/chartcmd/chart/constants"
	utils "github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/types"
)

func fillLatestPriceLine(chartView *[][]string, candles []types.Candle, max float64, min float64, maxPriceLabel int) {
	numCandles := len(candles)
	candle := candles[numCandles-1]
	var color string
	if candle.IsGreen {
		color = c.UpColorBold
	} else {
		color = c.DownColorBold
	}

	colIdx := numCandles + 1
	rowIdx := getChartBodyYIdx(candle.Close, max, min) + 1

	for i := range int(c.ChartBodyLeftPadding) + int(c.ChartBodyRightPadding) - 2 {
		(*chartView)[rowIdx][colIdx+i+1] = utils.Fill(c.LatestPrice, color)
	}

	if candle.IsGreen {
		color = c.UpColorBg
	} else {
		color = c.DownColorBg
	}
	labelStr := float64ToStr(candle.Close)[:maxPriceLabel-2]
	priceColIdx := int(c.ChartBodyLeftPadding) + int(c.ChartBodyRightPadding) + 1
	for i, letter := range labelStr {
		(*chartView)[rowIdx][colIdx+i+priceColIdx] = utils.Fill(string(letter), color)
	}

}
