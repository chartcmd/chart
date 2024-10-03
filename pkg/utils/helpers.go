package utils

import (
	"math"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
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

func GetArrowKeyInput() (int, int) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				return -1, 0
			case termbox.KeyArrowRight:
				return 1, 0
			case termbox.KeyArrowUp:
				return 0, 1
			case termbox.KeyArrowDown:
				return 0, -1
			case termbox.KeyEsc:
				os.Exit(1)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func IndexOf(slice []string, element string) int {
	for i, elem := range slice {
		if elem == element {
			return i
		}
	}
	return -1
}
