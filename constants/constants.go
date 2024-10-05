package constants

import (
	"fmt"

	"github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/pkg/utils/fetch/crypto"
	"github.com/chartcmd/chart/types"
	"golang.org/x/term"
)

var (
	// TODO move to chart and make a type
	ChartBodyCols          uint32 = 128
	ChartBodyRows          uint32 = 32
	ChartTopPadding        uint32 = 1
	ChartBottomPadding     uint32 = 2
	ChartAddlBottomSpace   uint32 = 2
	ChartBodyRightPadding  uint32 = 1
	ChartBodyLeftPadding   uint32 = 4
	ChartXAxisLeftPadding  uint32 = 2
	ChartYAxisRightPadding uint32 = 2
	NumYLabels             uint32 = uint32(ChartBodyRows / 4)
	NumXLabels             uint32 = NumYLabels
	NumCandles             uint32 = 128
	IntervalBarPadding     uint32 = 2

	CandleBody  string = "┃"
	WickBody    string = "│"
	WickTop     string = "╽"
	WickBottom  string = "╿"
	YAxis       string = "|"
	XAxis       string = "-"
	LatestPrice string = "▁"

	CoinbaseCandleMax uint32 = 300

	StreamRefreshRateMS int = 255

	CryptoList []string = crypto.GetCryptoList()
	StockList  []string = []string{}

	Intervals []string = []string{
		"1m",
		"5m",
		"15m",
		"1h",
		"6h",
		"1d",
	}

	IntervalToGranularity = map[string]uint32{
		"1m":  60,
		"5m":  300,
		"15m": 900,
		"1h":  3600,
		"6h":  21600,
		"1d":  86400,
	}

	ColorToAnsi = map[string]string{
		// colors
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

		// bold + colored
		"bold_black":   "\033[1;30m",
		"bold_red":     "\033[1;91m",
		"bold_green":   "\033[1;92m",
		"bold_yellow":  "\033[1;93m",
		"bold_purple":  "\033[1;34m",
		"bold_magenta": "\033[1;95m",
		"bold_cyan":    "\033[1;96m",
		"bold_white":   "\033[1;97m",
		"bold_gray":    "\033[1;90m",
		"bold_pink":    "\033[1;94m",

		// bg colors
		"bg_black":   "\033[40m",
		"bg_red":     "\033[101m",
		"bg_green":   "\033[102m",
		"bg_yellow":  "\033[103m",
		"bg_purple":  "\033[44m",
		"bg_blue":    "\033[44m",
		"bg_magenta": "\033[105m",
		"bg_cyan":    "\033[106m",
		"bg_white":   "\033[107m",
		"bg_gray":    "\033[100m",
		"bg_none":    "",
	}

	Config              types.Config
	UpColor             string = ColorToAnsi["green"]
	UpColorBold         string = ColorToAnsi["bold_green"]
	DownColor           string = ColorToAnsi["red"]
	DownColorBold       string = ColorToAnsi["bold_red"]
	UpColorBg           string = ColorToAnsi["bg_green"]
	DownColorBg         string = ColorToAnsi["bg_red"]
	CyanBg              string = ColorToAnsi["bg_cyan"]
	GrayBg              string = ColorToAnsi["bi_gray"]
	WhiteColor          string = ColorToAnsi["white"]
	WhiteColorBold      string = ColorToAnsi["bold_white"]
	SelectedColorBold   string = ColorToAnsi["bold_cyan"]
	UnselectedColorBold string = ColorToAnsi["bold_gray"]
	ResetColor          string = "\033[0m"
	SelectedColor       string = WhiteColor
	UnselectedColor     string = ColorToAnsi["gray"]
	EquitiesWL          []string
	CryptoWL            []string
	DefaultTimeFrame    string = "1D"
)

func init() {

	Config := utils.ReadConfig()

	UpColor = ColorToAnsi[Config.UpColor]
	DownColor = ColorToAnsi[Config.DownColor]
	UpColorBold = ColorToAnsi[fmt.Sprintf("bold_%s", Config.UpColor)]
	DownColorBold = ColorToAnsi[fmt.Sprintf("bold_%s", Config.DownColor)]
	UpColorBg = ColorToAnsi[fmt.Sprintf("bg_%s", Config.UpColor)]
	DownColorBg = ColorToAnsi[fmt.Sprintf("bg_%s", Config.DownColor)]
	EquitiesWL = Config.EquitiesWatchlist
	CryptoWL = Config.CryptoWatchlist
	DefaultTimeFrame = Config.DefaultTimeFrame

	// TODO make this refresh in display
	width, height, err := term.GetSize(0)
	if err == nil {
		ChartBodyCols = uint32(utils.GetClosestNumDivBy(2, width-15))
		// TODO subtract ChartBodyRows by 2 more for the timeframe viewer
		ChartBodyRows = uint32(utils.GetClosestNumDivBy(4, height-7)) - 3
		NumCandles = ChartBodyCols
		NumYLabels = uint32(ChartBodyRows / 6)
		NumXLabels = uint32(ChartBodyCols/12) + 1
	}

}
