package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/chartcmd/chart/pkg/chart"
)

var (
	isJSON   bool
	isStream bool
)

var validTickers = []string{"btc", "eth", "aapl", "googl", "msft"} // Add more as needed
// var validIntervals = []string{"1m", "5m", "15m", "1h", "4h", "1d", "1w"}

var rootCmd = &cobra.Command{
	Use:   "chart [ticker] [interval|dur] [duration]",
	Short: "Chart application",
	Long:  `A CLI application for charting financial data.`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return handleChartCommand(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		handleChartCommand(cmd, args)
	},
}

func handleChartCommand(cmd *cobra.Command, args []string) error {
	if len(args) == 2 {
		ticker := strings.ToLower(args[0])
		interval := args[1]
		// if !contains(validTickers, ticker) {
		// 	return fmt.Errorf("invalid ticker: %s", ticker)
		// }
		// if !contains(validIntervals, interval) {
		// 	return fmt.Errorf("invalid interval: %s", interval)
		// }

		err := chart.DrawChart(ticker, interval)
		if err != nil {
			return fmt.Errorf("error: %s", err)
		}
	} else if len(args) == 0 {
		cmd.Help()
	} else if len(args) == 3 && args[1] == "dur" {
		ticker := strings.ToLower(args[0])
		duration := args[2]
		// draw chart
		fmt.Printf("Generating chart for %s over the last %s\n", ticker, duration)
	} else {
		return fmt.Errorf("invalid number of arguments: %d", len(args))
	}
	return nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available charts",
	Args:  cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		data := map[string]interface{}{
			"charts": validTickers,
		}

		if isJSON {
			jsonData, _ := json.MarshalIndent(data, "", "  ")
			fmt.Println(string(jsonData))
		} else if isStream {
			fmt.Println("Streaming data...")
			for i := 0; i < 5; i++ {
				fmt.Printf("Data point %d: %v\n", i+1, validTickers)
			}
		} else if len(args) == 0 || (len(args) == 1 && args[0] == "help") {
			cmd.Help()
		} else {
			fmt.Println("Available charts:", validTickers)
		}
	},
}

var listStocksCmd = &cobra.Command{
	Use:   "stocks",
	Short: "List available stock charts",
	Run: func(cmd *cobra.Command, args []string) {
		stocks := []string{"aapl", "googl", "msft"}
		data := map[string]interface{}{
			"stocks": stocks,
		}

		if isJSON {
			jsonData, _ := json.MarshalIndent(data, "", "  ")
			fmt.Println(string(jsonData))
		} else {
			fmt.Println("Available stocks:", stocks)
		}
	},
}

var listCryptoCmd = &cobra.Command{
	Use:   "crypto",
	Short: "List available crypto charts",
	Run: func(cmd *cobra.Command, args []string) {
		crypto := []string{"btc", "eth"}
		data := map[string]interface{}{
			"crypto": crypto,
		}

		if isJSON {
			jsonData, _ := json.MarshalIndent(data, "", "  ")
			fmt.Println(string(jsonData))
		} else {
			fmt.Println("Available crypto:", crypto)
		}
	},
}

// func contains(slice []string, item string) bool {
// 	for _, a := range slice {
// 		if a == item {
// 			return true
// 		}
// 	}
// 	return false
// }

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listStocksCmd)
	listCmd.AddCommand(listCryptoCmd)

	listCmd.PersistentFlags().BoolVar(&isJSON, "json", false, "Output in JSON format")
	listCmd.Flags().BoolVar(&isStream, "stream", false, "Stream data")

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage:")
		fmt.Println("  chart <ticker> <interval>")
		fmt.Println("  chart <ticker> dur <time>")
		fmt.Println("  chart list [stocks|crypto] [flags]")
		fmt.Println("\nFlags:")
		fmt.Println("  --json       Output in JSON format (for list commands)")
		fmt.Println("  --stream     Stream data (for main list command only)")
		fmt.Println("\nAvailable Commands:")
		fmt.Println("  list        List available tickers")
		fmt.Println("  help        Help about any command")
		fmt.Println("\nUse \"chart [command] --help\" for more information about a command.")
	})

	listCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage:")
		fmt.Println("  chart list [stocks|crypto] [flags]")
		fmt.Println("\nFlags:")
		fmt.Println("  --json       Output in JSON format")
		fmt.Println("  --stream     Stream data (for main list command only)")
		fmt.Println("\nAvailable Commands:")
		fmt.Println("  stocks      List available stock charts")
		fmt.Println("  crypto      List available crypto charts")
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
