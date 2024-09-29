package build_chart

import (
	"math"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/types"
)

// fill_candles.go
func insertCandleWick(chartBody *[][]string, rowIdx int, colIdx int, dist int, dir int, isPositiveCandle bool) {
	color := c.UpColor
	if !isPositiveCandle {
		color = c.DownColor
	}
	rowIdx = rowIdx + (dist * dir)
	if rowIdx >= int(c.ChartBodyRows) || rowIdx <= 0 {
		return
	} else if colIdx >= int(c.ChartBodyCols) || colIdx <= 0 {
		return
	}
	(*chartBody)[rowIdx][colIdx] = color + c.WickBody + c.ResetColor
}

// fill_candles.go
func insertCandleBody(chartBody *[][]string, rowIdx int, colIdx int, dist int, isPositiveCandle bool) {
	var dir int = -1
	color := c.UpColor
	if !isPositiveCandle {
		dir = 1
		color = c.DownColor
	}
	rowIdx = rowIdx + (dist * dir)
	if rowIdx >= int(c.ChartBodyRows) || rowIdx <= 0 {
		return
	} else if colIdx >= int(c.ChartBodyCols) || colIdx <= 0 {
		return
	}
	(*chartBody)[rowIdx][colIdx] = color + c.CandleBody + c.ResetColor

}

// fill_candles.gp
func fillCandles(chartBody *[][]string, candles []types.Candle, max float64, min float64, pricePerYUnit float64) {
	for i, candle := range candles {
		isPositiveCandle := candle.Open < candle.Close

		sizeOfUpWick := math.Abs(candle.High-math.Max(candle.Open, candle.Close)) / pricePerYUnit
		rowIdx := getChartBodyYIdx(math.Max(candle.Open, candle.Close), max, min)
		if sizeOfUpWick > 0.5 {
			sizeOfUpWick := int(sizeOfUpWick) + 1
			for j := range sizeOfUpWick {
				insertCandleWick(chartBody, rowIdx, i, j, -1, isPositiveCandle)
			}
		}

		sizeOfDownWick := math.Abs(candle.Low-math.Min(candle.Open, candle.Close)) / pricePerYUnit
		rowIdx = getChartBodyYIdx(math.Min(candle.Open, candle.Close), max, min)
		if sizeOfDownWick > 0.5 {
			sizeOfDownWick := int(sizeOfDownWick) + 1
			for j := range sizeOfDownWick {
				insertCandleWick(chartBody, rowIdx, i, j, 1, isPositiveCandle)
			}
		}

		rowIdx = getChartBodyYIdx(candle.Open, max, min)
		sizeOfCandle := int(math.Abs(candle.Open-candle.Close)/pricePerYUnit + 1)
		for j := range sizeOfCandle {
			insertCandleBody(chartBody, rowIdx, i, j, isPositiveCandle)
		}

	}
}
