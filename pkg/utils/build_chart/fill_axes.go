package build_chart

import (
	"time"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/types"
)

func fillYAxis(chartView *[][]string, leftPadding int, priceLabels []float64, max float64, min float64) {
	for _, label := range priceLabels {
		rowIdx := getChartBodyYIdx(label, max, min)
		labelStr := float64ToStr(label)

		for i, letter := range labelStr {
			(*chartView)[rowIdx+int(c.ChartTopPadding)][i] = c.WhiteColor + string(letter) + c.ResetColor
		}

	}

	for i := range c.ChartBodyRows + c.ChartTopPadding + c.ChartBottomPadding {
		(*chartView)[i][leftPadding-1] = c.WhiteColor + c.YAxis + c.ResetColor
	}
}

func fillXAxis(chartView *[][]string, candles []types.Candle, leftPadding int) {
	var timestamps []time.Time
	for _, candle := range candles {
		timestamps = append(timestamps, candle.Time)
	}
	indices, timestampLabels := getTimestampLabels(timestamps)
	for i, idx := range indices {
		if i < len(timestamps) {
			ts := timestampLabels[i]
			offset := int(len(ts) / 2)
			for j, letter := range ts {
				(*chartView)[len(*chartView)-1][leftPadding+idx+j-offset] = c.WhiteColor + string(letter) + c.ResetColor
			}
		}
	}

	for i := range c.ChartBodyCols {
		(*chartView)[len(*chartView)-2][leftPadding+int(i)] = c.WhiteColor + c.XAxis + c.ResetColor
	}
}
