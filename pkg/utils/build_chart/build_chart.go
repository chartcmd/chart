package build_chart

import (
	"fmt"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/types"
)

func BuildChart(candles []types.Candle, timeDuration int) string {
	priceLabels := getRoundPriceLabels(candles)
	rightPadding := lengthOfMaxLabel(priceLabels) + 2

	topPriceLabel := priceLabels[len(priceLabels)-1]
	bottomPriceLabel := priceLabels[0]
	pricePerYUnit := getPricePerYUnit(topPriceLabel, bottomPriceLabel)

	chartBody := initChartBody()
	fillCandles(&chartBody, candles, topPriceLabel, bottomPriceLabel, pricePerYUnit)

	chartView := initChartView(rightPadding, chartBody)
	fmt.Println(matrixToString(chartView))
	fillYAxis(&chartView, priceLabels, topPriceLabel, bottomPriceLabel)
	fmt.Println(matrixToString(chartView))
	fillXAxis(&chartView, candles)

	return matrixToString(chartView)
}

func initChartView(rightPadding int, chartBody [][]string) [][]string {
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
