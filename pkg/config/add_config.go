package config

import (
	"fmt"
	"strings"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/pkg/utils/fetch/stocks"
)

func AddConfig(key string, values []string) error {
	config := utils.ReadConfig()

	switch key {
	case "equities_wl":
		for _, ticker := range values {
			if stocks.IsValidTicker(ticker) {
				config.EquitiesWatchlist = append(config.EquitiesWatchlist, strings.ToUpper(ticker))
			} else {
				return fmt.Errorf("invalid equities ticker: %s", ticker)
			}
		}
	case "crypto_wl":
		for _, ticker := range values {
			if utils.StrSliceContains(c.CryptoList, ticker, false) {
				config.CryptoWatchlist = append(config.CryptoWatchlist, strings.ToUpper(ticker))
			} else {
				return fmt.Errorf("invalid crypto ticker: %s", ticker)
			}
		}
	case "up_color", "down_color", "default_tf":
		return fmt.Errorf("use chart config set %s %s", key, values)
	default:
		return fmt.Errorf("unknown config key: %s", key)
	}

	err := utils.WriteConfig(config)
	if err != nil {
		return err
	}
	return nil
}
