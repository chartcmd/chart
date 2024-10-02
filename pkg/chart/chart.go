package chart

import (
	"fmt"
	"log"
	"time"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/utils/build_chart"
	"github.com/chartcmd/chart/pkg/utils/fetch/crypto"
)

func DrawChart(ticker, interval string, stream bool) error {
	if stream {
		return drawChartStream(ticker, interval)

	}
	return drawChart(ticker, interval)
}

func drawChartStream(ticker, interval string) error {
	granularity := c.IntervalToGranularity[interval]
	if granularity > 3600 {
		return fmt.Errorf("error: please use a smaller timeframe to stream [1m, 5m, 15m, 1h]")
	}

	candles, err := getCandles(ticker, interval)
	if err != nil {
		return fmt.Errorf("error: getting initial candles: %v", err)
	}

	latestPrice, err := crypto.GetCoinbaseLatest(ticker + "-USD")
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
			latestPrice, err := crypto.GetCoinbaseLatest(ticker + "-USD")
			if err != nil {
				log.Printf("Error getting latest price: %v", err)
				continue
			}

			curCandle = updateCurCandle(curCandle, latestPrice)
			candles[len(candles)-1] = curCandle

			chart := build_chart.BuildChart(candles)
			pctChange := ((candles[len(candles)-1].Close - candles[0].Open) / candles[0].Open) * 100
			display(ticker, latestPrice, chart, pctChange)

		case <-nextCandleTimer.C:
			newCandles, err := getCandles(ticker, interval)
			if err != nil {
				return fmt.Errorf("error: refetching candles: %v", err)
			}

			curCandle = initCurCandle(newCandles[len(newCandles)-1].Close, newCandles[len(newCandles)-1])
			newCandles = append(newCandles[1:], curCandle)
			candles = newCandles

			nextCandleTimer.Reset(time.Duration(granularity) * time.Second)

			// chart := build_chart.BuildChart(candles)
			// pctChange := ((candles[len(candles)-1].Close - candles[0].Open) / candles[0].Open) * 100
			// display(ticker, latestPrice, chart, pctChange)
		}
	}
}

func drawChart(ticker, interval string) error {
	candles, err := getCandles(ticker, interval)
	if err != nil {
		return err
	}

	latestPrice, err := crypto.GetCoinbaseLatest(ticker + "-USD")
	if err == nil {
		curCandle := initCurCandle(latestPrice, candles[len(candles)-1])
		candles = append(candles, curCandle)[1:]
	} else {
		fmt.Printf("error fetching latest price: %s", err)
	}

	pctChange := ((candles[len(candles)-1].Close - candles[0].Open) / candles[0].Open) * 100

	chart := build_chart.BuildChart(candles)
	display(ticker, latestPrice, chart, pctChange)
	return nil
}
