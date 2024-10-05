package main

import (
	"fmt"

	c "github.com/chartcmd/chart/constants"
	"github.com/chartcmd/chart/pkg/chart"
	"github.com/spf13/cobra"
)

/**
stream should be default
*/

var rootCmd = &cobra.Command{
	Use:   "chart <ticker> <interval> []flags",
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
	if len(args) == 0 {
		cmd.Help()
	} else if len(args) == 1 {
		// TODO
		// chart 4 charts at once of either crypto or stocks
		ticker := args[0]
		if ticker == "stocks" || ticker == "crypto" {
			if isStill {
				fmt.Printf("4 charts of %s", args[0])
			} else {
				fmt.Printf("streaming 4 charts of %s", args[0])
			}
		} else {
			err := chart.DrawChart(args[0], c.DefaultTimeFrame, isStill)
			if err != nil {
				return fmt.Errorf("error: %s", err)
			}
		}
	} else if len(args) == 2 {

		err := chart.DrawChart(args[0], args[1], isStill)
		if err != nil {
			return fmt.Errorf("error: %s", err)
		}

	} else {
		return fmt.Errorf("invalid number of arguments: %d, run \"chart help\" for more info", len(args))
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(configCmd)

	rootCmd.Flags().BoolVarP(&isStill, "still", "s", false, "Freeze data")

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage:")
		fmt.Println("  chart <ticker> <interval> [flags]")
		fmt.Println("  chart list [stocks|crypto] [flags]")
		fmt.Println("  chart config [set|add|pop] [key] [value]")
		fmt.Println("  chart config list [flags]")
		fmt.Println("\nFlags:")
		fmt.Println("  --json       Output in JSON format (for list commands)")
		fmt.Println("  -s, --still  Freeze data")
		fmt.Println("\nAvailable Commands:")
		fmt.Println("  list        List available tickers")
		fmt.Println("  config      Adjust config variables")
		fmt.Println("  help        Help about any command")
		fmt.Println("\nUse \"chart [command] --help\" for more information about a command.")
	})
}
