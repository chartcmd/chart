package stocks

import (
	"fmt"
)

func IsValidTicker(ticker string) bool {
	_, err := GetYFLatest(ticker)
	if err != nil {
		fmt.Printf("error error: %s", err)
		return false
	}
	return true
}
