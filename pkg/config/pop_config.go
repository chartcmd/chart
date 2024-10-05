package config

import (
	"fmt"
	"strings"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/utils"
	"github.com/chartcmd/chart/pkg/utils/fetch/stocks"
)

func PopConfig(key, value string) error {
	config := utils.ReadConfig()

	switch key {
	case "equities_wl":
		if stocks.IsValidTicker(value) {
			config.EquitiesWatchlist = utils.RemoveString(config.EquitiesWatchlist, strings.ToUpper(value))
		} else {
			return fmt.Errorf("invalid equities ticker: %s", value)
		}
	case "crypto_wl":
		if utils.StrSliceContains(c.CryptoList, value, false) {
			config.CryptoWatchlist = utils.RemoveString(config.CryptoWatchlist, strings.ToUpper(value))
		} else {
			return fmt.Errorf("invalid crypto ticker: %s", value)
		}
	case "up_color", "down_color", "default_tf":
		return fmt.Errorf("use chart config set %s %s", key, value)
	default:
		return fmt.Errorf("unknown config key: %s", key)
	}

	err := utils.WriteConfig(config)
	if err != nil {
		return err
	}
	return nil
}
