package chart

import (
	"fmt"

	c "github.com/chartcmd/chart/constants"
	utils "github.com/chartcmd/chart/pkg/utils"
)

func DisplayIntervalBar(selectedIntervalIdx int) {
	baseStr := ""
	for i, interval := range c.Intervals {
		interval = fmt.Sprintf("[%s]", interval)
		if i == selectedIntervalIdx {
			interval = utils.Fill(interval, c.SelectedColorBold)
		} else {
			interval = utils.Fill(interval, c.UnselectedColorBold)
		}
		baseStr += fmt.Sprintf("%*s%s%*s", c.IntervalBarPadding, "", interval, c.IntervalBarPadding, "")
	}
	// leftPadding := (c.ChartBodyCols - len(baseStr)) / 2
	leftPadding := (c.ChartBodyCols - 35) / 2
	fmt.Printf("%*s%s", leftPadding, "", baseStr)
}
