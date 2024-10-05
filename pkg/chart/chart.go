package chart

import (
	"fmt"
	"log"
	"math"
	"time"

	c "github.com/chartcmd/chart/constants"
	utils "github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/pkg/utils/build_chart"
	"github.com/chartcmd/chart/pkg/utils/fetch/stocks"
)

/**
Refactor:
Move all files except chart.go to utils/build_chart
Create a ChartConfig type that stores the info in constants, as well frequent params passed into funcs
Init this DrawChart -> chartConfig := createChartConfig(ticker, interval)
*/

type InputData struct {
	X int
	Y int
}

func DrawChart(ticker, interval string, isStill bool) error {
	intervalIdx := utils.IndexOf(c.Intervals, interval)

	if isStill {
		err := drawChart(ticker, intervalIdx)
		if err != nil {
			return err
		}
	} else {
		err := drawChartStream(ticker, intervalIdx)
		if err != nil {
			return err
		}
	}

	return nil
}

func drawChartStream(ticker string, intervalIdx int) error {
	numIntervals := len(c.Intervals)
	interval := c.Intervals[intervalIdx]
	granularity := c.IntervalToGranularity[interval]

	var watchlist []string
	if utils.StrSliceContains(c.CryptoList, ticker, false) {
		watchlist = c.CryptoWL
	} else if stocks.IsValidTicker(ticker) {
		watchlist = c.EquitiesWL
	}
	tickerIdx := utils.IndexOf(watchlist, ticker)
	numTickers := len(watchlist)

	isCursorOnInterval := 1

	candles, err := getCandles(ticker, interval)
	if err != nil {
		return fmt.Errorf("error: getting initial candles: %v", err)
	}
	if len(candles) > int(c.CoinbaseCandleMax) {
		return fmt.Errorf("error: use a smaller resolution to stream")
	}

	inputChan := make(chan InputData)

	go func() {
		for {
			x, y := utils.GetArrowKeyInput()
			inputChan <- InputData{X: x, Y: y}
		}
	}()

	latestPrice, err := GetLatest(ticker)
	if err != nil {
		return fmt.Errorf("error fetching latest price: %s", err)
	}

	lastCandle := candles[len(candles)-1]
	curCandle := initCurCandle(latestPrice, lastCandle)
	candles = append(candles[1:], curCandle)

	nextCandleTime := lastCandle.Time.Add(time.Duration(granularity) * time.Second)
	timeUntilNextCandle := time.Until(nextCandleTime)
	if timeUntilNextCandle < 0 {
		nextCandleTime = lastCandle.Time.Add((time.Duration(granularity) * time.Second) * 2)
		timeUntilNextCandle = time.Until(nextCandleTime)
	}

	refreshDuration := time.Duration(c.StreamRefreshRateMS) * time.Millisecond
	refreshTicker := time.NewTicker(refreshDuration)
	defer refreshTicker.Stop()

	nextCandleTimer := time.NewTimer(timeUntilNextCandle)
	defer nextCandleTimer.Stop()

	for {
		select {
		case <-refreshTicker.C:
			latestPrice, err := GetLatest(ticker)
			if err != nil {
				log.Printf("Error getting latest price: %v", err)
				continue
			}

			curCandle = updateCurCandle(curCandle, latestPrice)
			candles[len(candles)-1] = curCandle

			chart := build_chart.BuildChart(candles)
			pctChange := ((candles[len(candles)-1].Close - candles[0].Open) / candles[0].Open) * 100
			display(ticker, latestPrice, chart, pctChange)
			displayBar(intervalIdx, c.Intervals, isCursorOnInterval == 1)
			fmt.Println()
			displayBar(tickerIdx, watchlist, isCursorOnInterval == 0)

		case <-nextCandleTimer.C:
			newCandles, err := getCandles(ticker, interval)
			if err != nil {
				return fmt.Errorf("error: refetching candles: %v", err)
			}

			curCandle = initCurCandle(newCandles[len(newCandles)-1].Close, newCandles[len(newCandles)-1])
			newCandles = append(newCandles[1:], curCandle)
			candles = newCandles

			nextCandleTimer.Reset(time.Duration(granularity) * time.Second)

		case input := <-inputChan:
			if input.X == 1 {
				if isCursorOnInterval == 1 {
					intervalIdx = int(math.Min(float64(numIntervals-1), float64(intervalIdx+1)))
				} else {
					tickerIdx = int(math.Min(float64(numTickers-1), float64(tickerIdx+1)))
				}
			} else if input.X == -1 {
				if isCursorOnInterval == 1 {
					intervalIdx = int(math.Max(0, float64(intervalIdx-1)))
				} else {
					tickerIdx = int(math.Max(0, float64(tickerIdx-1)))
				}
			} else if input.Y == 1 {
				isCursorOnInterval = int(math.Min(1, float64(isCursorOnInterval)) + 1)
			} else if input.Y == -1 {
				isCursorOnInterval = int(math.Max(0, float64(isCursorOnInterval)) - 1)
			}

			if input.X != 0 {
				interval = c.Intervals[intervalIdx]
				ticker = watchlist[tickerIdx]
				newCandles, err := getCandles(ticker, interval)
				if err != nil {
					log.Printf("Error getting candles for new interval: %v", err)
					continue
				}
				curCandle = initCurCandle(newCandles[len(newCandles)-1].Close, newCandles[len(newCandles)-1])
				newCandles = append(newCandles[1:], curCandle)
				candles = newCandles
				granularity = c.IntervalToGranularity[interval]

				nextCandleTime = candles[len(candles)-1].Time.Add(time.Duration(granularity) * time.Second)
				nextCandleTimer.Reset(time.Until(nextCandleTime))
			}
		}
	}
}

func drawChart(ticker string, intervalIdx int) error {
	interval := c.Intervals[intervalIdx]
	candles, err := getCandles(ticker, interval)
	if err != nil {
		return err
	}

	latestPrice, err := GetLatest(ticker)
	if err == nil {
		curCandle := initCurCandle(latestPrice, candles[len(candles)-1])
		candles = append(candles, curCandle)[1:]
	} else {
		fmt.Printf("error fetching latest price: %s", err)
	}

	pctChange := ((candles[len(candles)-1].Close - candles[0].Open) / candles[0].Open) * 100

	chart := build_chart.BuildChart(candles)
	display(ticker, latestPrice, chart, pctChange)
	// displayIntervalBar(intervalIdx)
	return nil
}
