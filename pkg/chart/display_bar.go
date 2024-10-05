package chart

import (
	"fmt"

	c "github.com/chartcmd/chart/constants"
	utils "github.com/chartcmd/chart/pkg/utils"
)

func displayBar(selectedIdx int, options []string, rowIsSelected bool) {
	baseStr := ""
	watchlistBarLen := 1
	for i, option := range options {
		watchlistBarLen += len(option)
		selectedStr := ""
		option = fmt.Sprintf("[%s]", option)
		if i == selectedIdx {
			if rowIsSelected {
				selectedStr = utils.Fill("<", c.WhiteColorBold)
			}
			option = utils.Fill(option, c.SelectedColorBold)
		} else {
			option = utils.Fill(option, c.UnselectedColorBold)
		}
		baseStr += fmt.Sprintf("%*s%s%s%*s", c.IntervalBarPadding, "", option, selectedStr, c.IntervalBarPadding, "")
		watchlistBarLen += int(c.IntervalBarPadding) * 2
	}
	leftPadding := (int(c.ChartBodyCols) - watchlistBarLen) / 2
	fmt.Printf("%*s%s\n", leftPadding, "", baseStr)
}
