package build_chart

import (
	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/types"
)

func BuildChart(candles []types.Candle, timeDuration int) (string, int) {
	priceLabels := getRoundPriceLabels(candles)
	leftPadding := lengthOfMaxLabel(priceLabels) + 2

	topPriceLabel := priceLabels[len(priceLabels)-1]
	bottomPriceLabel := priceLabels[0]
	pricePerYUnit := getPricePerYUnit(topPriceLabel, bottomPriceLabel)

	chartBody := initChartBody()
	fillCandles(&chartBody, candles, topPriceLabel, bottomPriceLabel, pricePerYUnit)

	chartView := initChartView(leftPadding)
	fillYAxis(&chartView, leftPadding, priceLabels, topPriceLabel, bottomPriceLabel)
	fillXAxis(&chartView, candles, leftPadding)

	meshBodyToView(&chartView, chartBody, leftPadding)

	return matrixToString(chartView), leftPadding
}

func initChartView(leftPadding int) [][]string {
	chart := make([][]string, c.ChartBodyRows+c.ChartTopPadding+c.ChartBottomPadding+c.ChartAddlBottomSpace)
	for i := range chart {
		chart[i] = make([]string, c.ChartBodyCols+uint32(leftPadding))
		for j := range chart[i] {
			chart[i][j] = " "
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
