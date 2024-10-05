package main

import (
	"fmt"
	"os"
)

var (
	isJSON  bool
	isStill bool
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
