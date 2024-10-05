package config

import (
	"fmt"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/utils"
)

func SetConfig(key, value string) error {
	config := utils.ReadConfig()

	switch key {
	case "up_color":
		config.UpColor = value
	case "down_color":
		config.DownColor = value
	case "default_tf":
		if utils.StrSliceContains(c.Intervals, value, false) {
			config.DefaultTimeFrame = value
		} else {
			return fmt.Errorf("invalid interval: %s\nuse %s", value, c.Intervals)
		}
	case "equities_wl", "crypto_wl":
		return fmt.Errorf("use chart config add equities_wl %s", value)
	default:
		return fmt.Errorf("unknown config key: %s", key)
	}

	err := utils.WriteConfig(config)
	if err != nil {
		return err
	}
	return nil
}
