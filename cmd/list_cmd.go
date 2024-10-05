package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var validTickers = []string{"btc", "eth", "aapl", "googl", "msft"}

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
		} else if isStill {
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

var listHelpCmd = &cobra.Command{
	Use:   "help",
	Short: "Help about any command ",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	listCmd.AddCommand(listStocksCmd)
	listCmd.AddCommand(listCryptoCmd)
	listCmd.AddCommand(listHelpCmd)

	listCmd.PersistentFlags().BoolVar(&isJSON, "json", false, "Output in JSON format")
	listCmd.Flags().BoolVarP(&isStill, "still", "s", false, "Freeze data")

	listCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage:")
		fmt.Println("  chart list [stocks|crypto] [flags]")
		fmt.Println("  chart config help")
		fmt.Println("\nFlags:")
		fmt.Println("  --json       Output in JSON format")
		fmt.Println("  -s, --still  Freeze data")
		fmt.Println("\nAvailable Commands:")
		fmt.Println("  stocks      List available stock charts")
		fmt.Println("  crypto      List available crypto charts")
		fmt.Println("  help        Help about any command")
		fmt.Println("\nUse \"chart [command] --help\" for more information about a command.")
	})
}
