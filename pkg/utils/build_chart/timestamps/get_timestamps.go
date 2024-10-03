package timestamps

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chartcmd/chart/pkg/utils"
)

func Get15mTimestampLabels(timestamps []time.Time) ([]int, []string) {
	var indices []int
	var labels []string

	for i, ts := range timestamps {
		if ts.Minute()%15 == 0 {
			indices = append(indices, i)
			if ts.Minute() != 0 {
				labels = append(labels, fmt.Sprintf("%02d:%02d", ts.Hour(), ts.Minute()))
			} else if ts.Hour() != 0 {
				labels = append(labels, fmt.Sprintf("%02d:00", ts.Hour()))
			} else {
				labels = append(labels, ts.Format("Jan 02"))
			}
		}
	}

	return indices, labels
}

func Get1hTimestampLabels(timestamps []time.Time) ([]int, []string) {
	var indices []int
	var labels []string

	for i, ts := range timestamps {
		if ts.Minute() == 0 {
			indices = append(indices, i)
			if ts.Hour() != 0 {
				labels = append(labels, fmt.Sprintf("%02d:00", ts.Hour()))
			} else {
				labels = append(labels, ts.Format("Jan 02"))
			}
		}
	}

	return indices, labels
}

func Get4hTimestampLabels(timestamps []time.Time) ([]int, []string) {
	var indices []int
	var labels []string

	for i, ts := range timestamps {
		if ts.Minute() == 0 && ts.Hour()%4 == 0 {
			indices = append(indices, i)
			if ts.Hour() != 0 {
				labels = append(labels, fmt.Sprintf("%02d:00", ts.Hour()))
			} else {
				labels = append(labels, ts.Format("Jan 02"))
			}
		}
	}

	return indices, labels
}

func Get1dTimestampLabels(timestamps []time.Time) ([]int, []string) {
	var indices []int
	var labels []string

	for i, ts := range timestamps {
		if ts.Hour() == 0 && ts.Minute() == 0 {
			indices = append(indices, i)
			if ts.Day() == 1 {
				labels = append(labels, ts.Format("Jan"))
			} else {
				labels = append(labels, strconv.Itoa(ts.Day()))
			}
		}
	}

	return indices, labels
}

func Get1wTimestampLabels(timestamps []time.Time) ([]int, []string) {
	var indices []int
	var labels []string

	for i, ts := range timestamps {
		if (ts.Day() == 1 || ts.Day()%7 == 0) && (ts.Hour() == (24+utils.GetUTCOffsetHours())%24) {
			indices = append(indices, i)
			if ts.Day() == 1 {
				labels = append(labels, ts.Format("Jan"))
			} else {
				labels = append(labels, strconv.Itoa(ts.Day()))
			}
		}
	}

	return indices, labels
}

func Get1mTimestampLabels(timestamps []time.Time) ([]int, []string) {
	var indices []int
	var labels []string

	for i, ts := range timestamps {
		if ts.Day() == 1 {
			indices = append(indices, i)
			labels = append(labels, ts.Format("Jan"))
		}
	}

	return indices, labels
}
