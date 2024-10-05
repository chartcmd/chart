package main

import (
	"fmt"

	"github.com/chartcmd/chart/pkg/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure settings",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set config variables",
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.SetConfig(args[0], args[1])
	},
}

var configAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add config to variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.AddConfig(args[0], args[1:])
	},
}

var configPopCmd = &cobra.Command{
	Use:   "pop",
	Short: "Remove config from variable",
	RunE: func(cmd *cobra.Command, args []string) error {
		return config.PopConfig(args[0], args[1:])
	},
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List config variables",
	Run: func(cmd *cobra.Command, args []string) {
		config.ListConfig(isJSON)
	},
}

var configHelpCmd = &cobra.Command{
	Use:   "help",
	Short: "Help about any command ",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	configListCmd.PersistentFlags().BoolVar(&isJSON, "json", false, "Output in JSON format")
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configAddCmd)
	configCmd.AddCommand(configPopCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configHelpCmd)

	configCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage:")
		fmt.Println("  chart config [set|add|pop] [key] [value]")
		fmt.Println("  chart config list [flags]")
		fmt.Println("  chart config help")
		fmt.Println("\nFlags:")
		fmt.Println("  --json    Output in JSON format")
		fmt.Println("\nAvailable Commands:")
		fmt.Println("  set       Set config variabless")
		fmt.Println("  add       Add config to variable")
		fmt.Println("  pop       Pop config from variable")
		fmt.Println("  list      List config variables")
		fmt.Println("  help      Help about any command")
		fmt.Println("\nUse \"chart [command] --help\" for more information about a command.")
	})
}
