package build_chart

import (
	"fmt"
	"time"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/types"
)

func fillYAxis(chartView *[][]string, priceLabels []float64, max float64, min float64) {
	for _, label := range priceLabels {
		rowIdx := getChartBodyYIdx(label, max, min)
		labelStr := float64ToStr(label)

		for i, letter := range labelStr {
			(*chartView)[rowIdx+int(c.ChartTopPadding)][c.ChartBodyCols+uint32(i)+c.ChartBodyRightPadding+c.ChartBodyLeftPadding+2] = c.WhiteColor + string(letter) + c.ResetColor
		}

	}

	for i := range c.ChartBodyRows + c.ChartTopPadding + c.ChartBottomPadding {
		(*chartView)[i][c.ChartBodyCols+c.ChartBodyRightPadding+c.ChartBodyLeftPadding] = c.WhiteColor + c.YAxis + c.ResetColor
	}
}

func fillXAxis(chartView *[][]string, candles []types.Candle) {
	var timestamps []time.Time
	for _, candle := range candles {
		timestamps = append(timestamps, candle.Time)
	}
	indices, timestampLabels := getTimestampLabels(timestamps)
	for i, idx := range indices {
		fmt.Println(idx)
		if i < len(timestamps) {
			ts := timestampLabels[i]
			offset := int(len(ts) / 2)
			for j, letter := range ts {
				(*chartView)[len(*chartView)-1][idx+j-offset+int(c.ChartBodyLeftPadding)] = c.WhiteColor + string(letter) + c.ResetColor
			}
		}
	}

	for i := range c.ChartBodyCols + c.ChartBodyRightPadding + c.ChartXAxisLeftPadding + 1 {
		(*chartView)[len(*chartView)-2][i+c.ChartXAxisLeftPadding] = c.WhiteColor + c.XAxis + c.ResetColor
	}
}
