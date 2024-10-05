package config

import (
	"fmt"

	c "github.com/chartcmd/chart/constants"
)

func OptionsConfig() {
	fmt.Println("\nConfig options:")
	fmt.Println("up_color:   [red green yellow purple magenta cyan white gray black]")
	fmt.Println("down_color: [red green yellow purple magenta cyan white gray black]")
	fmt.Printf("default_tf: %s\n", c.Intervals)
	fmt.Println("equity_wl:  any shorthand equity ticker (AAPL, TSLA, etc) available on Yahoo Finance")
	fmt.Println("crypto_wl:  any shorthand crypto ticker (BTC, ETH, etc) available on Coinbase")
}
