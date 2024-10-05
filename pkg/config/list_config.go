package config

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/chartcmd/chart/pkg/utils"
)

func ListConfig(isJson bool) error {
	config := utils.ReadConfig()

	if isJson {
		jsonData, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling config to JSON: %w", err)
		}
		fmt.Println(string(jsonData))
	} else {
		fmt.Println("\nConfig variables:")
		fmt.Printf("up_color=%s\n", config.UpColor)
		fmt.Printf("down_color=%s\n", config.DownColor)
		fmt.Printf("default_tf=%s\n", config.DefaultTimeFrame)
		fmt.Printf("equity_wl=%s\n", strings.Join(config.EquitiesWatchlist, ","))
		fmt.Printf("crypto_wl=%s\n", strings.Join(config.CryptoWatchlist, ","))
	}
	return nil
}
