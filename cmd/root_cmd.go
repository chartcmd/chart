package main

import (
	"fmt"
	"strings"

	"github.com/chartcmd/chart/pkg/chart"
	"github.com/spf13/cobra"
)

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

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(configCmd)

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
}
