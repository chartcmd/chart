package stocks

import (
	"fmt"
	"strings"
	"time"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/types"
	yfchart "github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/shopspring/decimal"
)

func isWeekend(t time.Time) bool {
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

func calculateStartDate(end time.Time, totalDuration time.Duration) time.Time {
	start := end
	for totalDuration > 0 {
		start = start.Add(-24 * time.Hour)
		if !isWeekend(start) && !isHoliday(start) {
			totalDuration -= 24 * time.Hour
		}
	}
	return start
}

func GetYFCandleStick(ticker string, interval string) ([]types.Candle, error) {
	end := time.Now()
	totalDuration := time.Duration(c.IntervalToGranularity[interval]*c.NumCandles) * time.Second
	start := calculateStartDate(end, totalDuration)

	yahooInterval, err := intervalToYahooInterval(interval)
	if err != nil {
		return nil, err
	}

	params := &yfchart.Params{
		Symbol:   strings.ToUpper(ticker),
		Start:    datetime.New(&start),
		End:      datetime.New(&end),
		Interval: yahooInterval,
	}
	iter := yfchart.Get(params)

	var candles []types.Candle
	for iter.Next() {
		bar := iter.Bar()
		candles = append(candles, types.Candle{
			Time:    time.Unix(int64(bar.Timestamp), 0).Add(30 * time.Minute),
			Low:     decimalToFloat64(bar.Low),
			High:    decimalToFloat64(bar.High),
			Open:    decimalToFloat64(bar.Open),
			Close:   decimalToFloat64(bar.Close),
			IsGreen: bar.Open.LessThan(bar.Close),
		})
	}

	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("error fetching stock data: %v", err)
	}

	return candles, nil
}

func decimalToFloat64(d decimal.Decimal) float64 {
	f, _ := d.Float64()
	return f
}

func intervalToYahooInterval(interval string) (datetime.Interval, error) {
	switch interval {
	case "1m":
		return datetime.OneMin, nil
	case "5m":
		return datetime.FiveMins, nil
	case "15m":
		return datetime.FifteenMins, nil
	case "1h":
		return datetime.OneHour, nil
	case "1d":
		return datetime.OneDay, nil
	default:
		return "", fmt.Errorf("unsupported interval: %s", interval)
	}
}
