package utils

import (
	"os"
	"os/exec"
	"strings"
)

func GetClosestNumDivBy(num, threshold int) int {
	prev := num
	for {
		cur := prev + num
		if cur > threshold {
			return prev
		}
		prev = cur
	}
}

func Fill(text, color string) string {
	return color + text + "\033[0m"
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func StrSliceContains(slice []string, val string, caseSensitive bool) bool {
	for _, str := range slice {
		if val == str {
			return true
		}
		if !caseSensitive && strings.EqualFold(val, str) {
			return true
		}
	}
	return false
}
