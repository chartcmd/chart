package build_chart

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/types"
)

func meshBodyToView(chartView *[][]string, chartBody [][]string, leftPadding int) {
	for i, row := range chartBody {
		for j, char := range row {
			(*chartView)[int(i)+int(c.ChartTopPadding)][leftPadding+j] = char
		}
	}
}

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

func formatTimestamp(ts time.Time) string {
	if ts.Minute() == 0 {
		if ts.Hour() == 0 {
			if ts.Day() == 1 {
				return ts.Format("Jan")
			}
			return strconv.Itoa(ts.Day())
		}
		return fmt.Sprintf("%02d:00", ts.Hour())
	}
	return fmt.Sprintf("%02d:%02d", ts.Hour(), ts.Minute())
}

func getTimestampLabels(timestamps []time.Time) ([]int, []string) {
	if len(timestamps) == 0 {
		return nil, nil
	}

	roundestIndex, _ := findRoundestTimestamp(timestamps)
	windowSize := (len(timestamps) / 8) - 1
	if windowSize < 1 {
		windowSize = 1
	}
	startIndexInWindow := roundestIndex % windowSize

	var labels []string
	var indices []int
	for i := 0; i < 8; i++ {
		index := startIndexInWindow + (i * windowSize)
		if index < len(timestamps) {
			indices = append(indices, index)
			labels = append(labels, formatTimestamp(timestamps[index]))
		}
	}

	return indices, labels
}

func findRoundestTimestamp(timestamps []time.Time) (int, time.Time) {
	roundestIndex := 0
	minScore := math.MaxInt32

	for i, ts := range timestamps {
		score := getRoundnessTimestampScore(ts)
		if score < minScore {
			minScore = score
			roundestIndex = i
		}
	}

	return roundestIndex, timestamps[roundestIndex]
}

// TODO: how else to do this
func getRoundnessTimestampScore(ts time.Time) int {
	year := ts.Year()
	month := ts.Month()
	day := ts.Day()
	hour := ts.Hour()
	minute := ts.Minute()

	isStartOfDay := hour*minute == 0
	isStartOfMonth := day+(hour*minute) == 1

	if year%100 == 0 && month == 1 && isStartOfMonth {
		return 1
	}
	if year%50 == 0 && month == 1 && isStartOfMonth {
		return 2
	}
	if year%10 == 0 && month == 1 && isStartOfMonth {
		return 3
	}
	if year%5 == 0 && month == 1 && isStartOfMonth {
		return 4
	}
	if month == 1 && isStartOfMonth {
		return 5
	}
	if (month == 1 || month == 4 || month == 7 || month == 10) && isStartOfMonth {
		return 6
	}
	if isStartOfMonth {
		return 7
	}
	if day == 15 && isStartOfDay {
		return 8
	}
	if day%3 == 1 && isStartOfDay {
		return 9
	}
	if isStartOfDay {
		return 10
	}
	if hour%12 == 0 && minute == 0 {
		return 11
	}
	if hour%6 == 0 && minute == 0 {
		return 12
	}
	if hour%3 == 0 && minute == 0 {
		return 13
	}
	if minute == 0 {
		return 14
	}
	if minute%30 == 0 {
		return 15
	}
	if minute%15 == 0 {
		return 16
	}

	return 17
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

func lengthOfMaxLabel(priceLabels []float64) int {
	maxLen := 0
	for _, f := range priceLabels {
		s := float64ToStr(f)
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}
	return maxLen
}
