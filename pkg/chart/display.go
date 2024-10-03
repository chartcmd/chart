package chart

import (
	"fmt"
	"strings"
	"time"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/utils"
)

func display(ticker string, latestPrice float64, chart string, pctChange float64, intervalIdx int) {
	utils.ClearScreen()
	now := time.Now()
	timeString := now.Format("15:04:05")

	ticker = strings.ToUpper(ticker)
	var priceStr string
	if latestPrice < 0.1 {
		priceStr = fmt.Sprintf("%s: $%.8f", c.WhiteColorBold+ticker, latestPrice)
	} else {
		priceStr = fmt.Sprintf("%s: $%.2f", c.WhiteColorBold+ticker, latestPrice)
	}
	fmt.Printf("\n%*s%s%*s%s\n", 4, "", priceStr, int(c.ChartBodyCols)-20-(len(priceStr)-20), "", timeString)
	if pctChange > 0 {
		fmt.Printf("    %s%.2f%%\n", c.UpColorBold, pctChange)
	} else if pctChange < 0 {
		fmt.Printf("    %s%.2f%%\n", c.DownColorBold, pctChange)
	} else {
		fmt.Printf("    %s%.2f%%\n", c.WhiteColorBold, pctChange)
	}

	fmt.Println(chart + c.ResetColor)
	DisplayIntervalBar(intervalIdx)
}
