package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure settings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CMD")
	},
}
