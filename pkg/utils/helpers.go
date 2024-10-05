package utils

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/chartcmd/chart/types"
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

func RemoveString(slice []string, s string) []string {
	result := make([]string, 0)
	for _, v := range slice {
		if v != s {
			result = append(result, v)
		}
	}
	return result
}

func IndexOf(slice []string, element string) int {
	for i, elem := range slice {
		if strings.EqualFold(elem, element) {
			return i
		}
	}
	return -1
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

func getConfigFilePath() (string, error) {
	var configHome string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %w", err)
	}
	configHome = filepath.Join(homeDir, ".config")

	return filepath.Join(configHome, "chart", "config.json"), nil
}

func ReadConfig() types.Config {
	var config types.Config
	filePath, err := getConfigFilePath()
	if err != nil {
		return types.Config{}
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return types.Config{}
	} else {
		err = json.Unmarshal(data, &config)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			os.Exit(1)
		}
		return config
	}
}

func WriteConfig(config types.Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	tempFile := filePath + ".tmp"
	if err := os.WriteFile(tempFile, data, 0644); err != nil {
		return fmt.Errorf("error writing temporary config file: %w", err)
	}

	if err := os.Rename(tempFile, filePath); err != nil {
		return fmt.Errorf("error replacing config file: %w", err)
	}

	return nil
}

func IsBrightColor(color string) bool {
	brightColors := []string{"\033[107m", "\033[102m", "\033[103m", "\033[106m"}
	for _, brightColor := range brightColors {
		if color == brightColor {
			return true
		}
	}
	return false
}
