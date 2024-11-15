package build_chart

import (
	"time"

	c "github.com/chartcmd/chart/constants"
	utils "github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/types"
)

func fillYAxis(chartView *[][]string, priceLabels []float64, max float64, min float64) {
	colIdx := c.ChartBodyCols + c.ChartBodyRightPadding + c.ChartBodyLeftPadding + c.ChartYAxisRightPadding
	for _, label := range priceLabels {
		rowIdx := getChartBodyYIdx(label, max, min)
		rowIdx = rowIdx + int(c.ChartTopPadding)
		labelStr := float64ToStr(label)

		for i, letter := range labelStr {
			(*chartView)[rowIdx][colIdx+uint32(i)] = utils.Fill(string(letter), c.WhiteColor)
		}

	}

	for i := range c.ChartBodyRows + c.ChartTopPadding + c.ChartBottomPadding {
		(*chartView)[i][colIdx-c.ChartYAxisRightPadding] = utils.Fill(c.YAxis, c.WhiteColor)
	}
}

func fillXAxis(chartView *[][]string, candles []types.Candle, leftOffset int) {
	var timestamps []time.Time
	for _, candle := range candles {
		timestamps = append(timestamps, candle.Time)
	}
	indices, timestampLabels := getTimestampLabels(timestamps)
	rowIdx := len(*chartView) - 1
	for i, idx := range indices {
		// TODO do we need this if statement
		if i < len(timestamps) {
			ts := timestampLabels[i]
			offset := int(len(ts) / 2)
			// TODO do -1 to compensate for removing first candle
			colIdx := idx - offset + int(c.ChartBodyLeftPadding) + leftOffset
			for j, letter := range ts {
				(*chartView)[rowIdx][colIdx+j] = utils.Fill(string(letter), c.WhiteColor)
			}
		}
	}

	for i := range c.ChartBodyCols + c.ChartBodyRightPadding + c.ChartXAxisLeftPadding + 1 {
		(*chartView)[rowIdx-1][i+c.ChartXAxisLeftPadding] = utils.Fill(c.XAxis, c.WhiteColor)
	}
}
