package build_chart

import (
	"math"
	"strconv"
	"strings"
	"time"

	c "github.com/chartcmd/chart/constants"
	ts "github.com/chartcmd/chart/pkg/utils/build_chart/timestamps"
	"github.com/chartcmd/chart/types"
)

func getChartBodyYIdx(price float64, max float64, min float64) int {
	idx := float64(c.ChartBodyRows) - (price-min)/(max-min)*float64(c.ChartBodyRows)
	idx = math.Round(idx)
	idx = math.Max(0, math.Min(float64(c.ChartBodyRows), idx))

	return int(idx)
}

func matrixToString(chart [][]string) string {
	var result strings.Builder
	for _, row := range chart {
		result.WriteString(strings.Join(row, "") + "\n")
	}
	return result.String()
}

func getMaxPrice(candles []types.Candle) float64 {
	maxPrice := candles[0].High
	for _, candle := range candles {
		high := candle.High
		if high > maxPrice {
			maxPrice = high
		}
	}
	return maxPrice
}

func getMinPrice(candles []types.Candle) float64 {
	minPrice := candles[0].Low
	for _, candle := range candles {
		low := candle.Low
		if low < minPrice {
			minPrice = low
		}
	}
	return minPrice
}

func getPricePerYUnit(max float64, min float64) float64 {
	return (max - min) / float64(c.ChartBodyRows)
}

func getTimestampLabels(timestamps []time.Time) ([]int, []string) {
	if len(timestamps) == 0 {
		return nil, nil
	}
	granularity := int64(timestamps[1].Sub(timestamps[0]).Seconds())
	if granularity == int64(c.IntervalToGranularity["1m"]) {
		return ts.Get15mTimestampLabels(timestamps)
	} else if granularity == int64(c.IntervalToGranularity["5m"]) {
		return ts.Get1hTimestampLabels(timestamps)
	} else if granularity == int64(c.IntervalToGranularity["15m"]) {
		return ts.Get4hTimestampLabels(timestamps)
	} else if granularity == int64(c.IntervalToGranularity["1h"]) {
		return ts.Get1dTimestampLabels(timestamps)
	} else if granularity == int64(c.IntervalToGranularity["6h"]) {
		return ts.Get1wTimestampLabels(timestamps)
	} else if granularity == int64(c.IntervalToGranularity["1d"]) {
		return ts.Get1mTimestampLabels(timestamps)
	}
	return []int{}, []string{}
}

func getRoundPriceLabels(candles []types.Candle) []float64 {
	max := getMaxPrice(candles)
	min := getMinPrice(candles)
	targetCount := c.NumYLabels

	rangeVal := max - min
	step := rangeVal / float64(targetCount-1)
	magnitude := math.Pow(10, math.Floor(math.Log10(step)))
	niceSteps := []float64{1, 2, 2.5, 5, 10}

	for _, niceStep := range niceSteps {
		if magnitude*niceStep >= step {
			step = magnitude * niceStep
			break
		}
	}

	start := math.Floor(min/step) * step
	end := math.Ceil(max/step) * step

	var labels []float64
	for current := start; current < end+step/2; current += step {
		labels = append(labels, math.Round(current*1e10)/1e10)
	}

	if labels[0] > min {
		labels = append([]float64{math.Round(min*1e10) / 1e10}, labels...)
	}
	if labels[len(labels)-1] < max {
		labels = append(labels, math.Round(max*1e10)/1e10)
	}

	return labels
}

func float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func maxLengthOfLabel(priceLabels []float64) int {
	maxLen := 0
	for _, f := range priceLabels {
		s := float64ToStr(f)
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}
	return maxLen
}
