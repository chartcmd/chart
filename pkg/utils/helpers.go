package utils

import (
	"math"
	"os"
	"os/exec"
	"strings"
	"time"
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

func GetUTCOffsetHours() int {
	now := time.Now()
	_, offset := now.Zone()

	hours := float64(offset) / 3600.0
	return int(math.Round(hours*100) / 100)
}
