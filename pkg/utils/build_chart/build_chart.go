package build_chart

import (
	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/types"
)

func BuildChart(candles []types.Candle) string {
	leftOffset := int(c.ChartBodyCols) - len(candles)
	// leftOffset := 0
	priceLabels := getRoundPriceLabels(candles)
	rightPadding := maxLengthOfLabel(priceLabels) + 2

	topPriceLabel := priceLabels[len(priceLabels)-1]
	bottomPriceLabel := priceLabels[0]
	pricePerYUnit := getPricePerYUnit(topPriceLabel, bottomPriceLabel)

	chartBody := initChartBody()
	fillCandles(&chartBody, candles, topPriceLabel, bottomPriceLabel, pricePerYUnit, leftOffset)

	chartView := initChartView(chartBody, rightPadding)
	fillYAxis(&chartView, priceLabels, topPriceLabel, bottomPriceLabel)
	fillXAxis(&chartView, candles, leftOffset)

	return matrixToString(chartView)
}

func initChartView(chartBody [][]string, rightPadding int) [][]string {
	chart := make([][]string, c.ChartBodyRows+c.ChartTopPadding+c.ChartBottomPadding+c.ChartAddlBottomSpace)
	for i := range chart {
		chart[i] = make([]string, c.ChartBodyCols+uint32(rightPadding)+c.ChartBodyRightPadding+c.ChartBodyLeftPadding)
		for j := range chart[i] {
			chart[i][j] = " "
		}
	}

	for i, row := range chartBody {
		for j, char := range row {
			chart[int(i)+int(c.ChartTopPadding)][j] = char
		}
	}
	return chart
}

func initChartBody() [][]string {
	chart := make([][]string, c.ChartBodyRows)
	for i := range chart {
		chart[i] = make([]string, c.ChartBodyCols)
		for j := range chart[i] {
			chart[i][j] = " "
		}
	}
	return chart
}
