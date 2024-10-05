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
	var textColor string
	var bgColor string
	if candle.IsGreen {
		color = c.UpColorBold
		bgColor = c.UpColorBg
	} else {
		color = c.DownColorBold
		bgColor = c.DownColorBg
	}
	if color == c.ColorToAnsi["bold_white"] {
		textColor = c.ColorToAnsi["bold_black"]
	} else {
		textColor = c.ColorToAnsi["bold_white"]
	}

	colIdx := int(c.ChartBodyCols) + 1
	rowIdx := getChartBodyYIdx(candle.Close, max, min) + 1

	for i := range int(c.ChartBodyLeftPadding) + int(c.ChartBodyRightPadding) - 2 {
		(*chartView)[rowIdx][colIdx+i+1] = utils.Fill(c.LatestPrice, color)
	}

	sliceAmount := maxPriceLabel - 2
	var labelStr string = float64ToStr(candle.Close)
	if len(labelStr) >= sliceAmount {
		labelStr = labelStr[:sliceAmount]
	}
	priceColIdx := int(c.ChartBodyLeftPadding) + int(c.ChartBodyRightPadding) + 1
	for i, letter := range labelStr {
		letter_str := utils.Fill(string(letter), textColor)
		(*chartView)[rowIdx][colIdx+i+priceColIdx] = utils.Fill(letter_str, bgColor)
	}

}
