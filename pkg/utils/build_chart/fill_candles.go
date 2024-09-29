package build_chart

import (
	"math"

	c "github.com/chartcmd/chart/constants"
	utils "github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/types"
)

func insertCandleWick(chartBody *[][]string, rowIdx int, colIdx int, dist int, dir int, isGreen bool) {
	color := c.UpColor
	if !isGreen {
		color = c.DownColor
	}
	rowIdx = rowIdx + (dist * dir)
	if rowIdx >= int(c.ChartBodyRows) || rowIdx <= 0 {
		return
	} else if colIdx >= int(c.ChartBodyCols) || colIdx <= 0 {
		return
	}
	(*chartBody)[rowIdx][colIdx] = utils.Fill(c.WickBody, color)
}

func insertCandleBody(chartBody *[][]string, rowIdx int, colIdx int, dist int, isGreen bool) {
	var dir int = -1
	color := c.UpColor
	if !isGreen {
		dir = 1
		color = c.DownColor
	}
	rowIdx = rowIdx + (dist * dir)
	if rowIdx >= int(c.ChartBodyRows) || rowIdx <= 0 {
		return
	} else if colIdx >= int(c.ChartBodyCols) || colIdx <= 0 {
		return
	}
	(*chartBody)[rowIdx][colIdx] = utils.Fill(c.CandleBody, color)

}

func fillCandles(chartBody *[][]string, candles []types.Candle, max float64, min float64, pricePerYUnit float64) {
	for i, candle := range candles {
		colIdx := i + int(c.ChartBodyLeftPadding)

		sizeOfUpWick := math.Abs(candle.High-math.Max(candle.Open, candle.Close)) / pricePerYUnit
		if sizeOfUpWick > 0.5 {
			rowIdx := getChartBodyYIdx(math.Max(candle.Open, candle.Close), max, min)
			for j := range int(sizeOfUpWick) + 1 {
				insertCandleWick(chartBody, rowIdx, colIdx, j, -1, candle.IsGreen)
			}
		}

		sizeOfDownWick := math.Abs(candle.Low-math.Min(candle.Open, candle.Close)) / pricePerYUnit
		if sizeOfDownWick > 0.5 {
			rowIdx := getChartBodyYIdx(math.Min(candle.Open, candle.Close), max, min)
			for j := range int(sizeOfDownWick) + 1 {
				insertCandleWick(chartBody, rowIdx, colIdx, j, 1, candle.IsGreen)
			}
		}

		rowIdx := getChartBodyYIdx(candle.Open, max, min)
		sizeOfCandle := int(math.Abs(candle.Open-candle.Close)/pricePerYUnit + 1)
		for j := range sizeOfCandle {
			insertCandleBody(chartBody, rowIdx, colIdx, j, candle.IsGreen)
		}

	}
}
