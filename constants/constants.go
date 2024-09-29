package constants

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/types"
	"golang.org/x/term"
)

var (
	CoinbaseBaseUrl               = "https://api.exchange.coinbase.com"
	CoinbaseCandleEndpointUrl     = "/products/%s/candles"
	CoinbaseCandleEndpointFullUrl = "%s%s?start=%s&end=%s&granularity=%d"

	ChartBodyCols         uint32 = 128
	ChartBodyRows         uint32 = 32
	ChartTopPadding       uint32 = 2
	ChartBottomPadding    uint32 = 2
	ChartAddlBottomSpace  uint32 = 2
	ChartBodyRightPadding uint32 = 1
	ChartBodyLeftPadding  uint32 = 4
	ChartXAxisLeftPadding uint32 = 2
	NumYLabels            uint32 = uint32(ChartBodyRows / 4)
	NumXLabels            uint32 = NumYLabels
	NumCandles            uint32 = 128

	CandleBody string = "█"
	WickBody   string = "│"
	YAxis      string = "|"
	XAxis      string = "-"

	IntervalOptions = "[%s] [%s] [%s] [%s] [%s] [%s]"

	IntervalToGranularity = map[string]uint32{
		"1m":  60,
		"5m":  300,
		"15m": 900,
		"1h":  3600,
		"6h":  21600,
		"1d":  86400,
	}

	ColorToAnsi = map[string]string{
		"black":   "\033[30m",
		"red":     "\033[91m",
		"green":   "\033[92m",
		"yellow":  "\033[93m",
		"purple":  "\033[34m",
		"magenta": "\033[95m",
		"cyan":    "\033[96m",
		"white":   "\033[97m",
		"gray":    "\033[90m",
		"pink":    "\033[94m",
	}

	Config          types.Config
	UpColor         string = ColorToAnsi["green"]
	DownColor       string = ColorToAnsi["red"]
	WhiteColor      string = ColorToAnsi["white"]
	ResetColor      string = "\033[0m"
	SelectedColor   string = WhiteColor
	UnselectedColor string = ColorToAnsi["gray"]
	Equities        []string
	Crypto          []string
)

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func init() {
	fileContent, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
	} else {
		err = json.Unmarshal(fileContent, &Config)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			os.Exit(1)
		}
		UpColor = ColorToAnsi[Config.ChartColors.Up]
		DownColor = ColorToAnsi[Config.ChartColors.Down]
		Equities = Config.ListCmdTickers.Equities
		Crypto = Config.ListCmdTickers.Crypto
	}

	width, height, err := term.GetSize(0)
	if err == nil {
		ChartBodyCols = uint32(utils.GetClosestNumDivBy(32, width))
		ChartBodyRows = uint32(utils.GetClosestNumDivBy(8, height-12))
		NumCandles = ChartBodyCols
		NumYLabels = uint32(ChartBodyRows / 4)
		NumXLabels = NumYLabels
	}

}
