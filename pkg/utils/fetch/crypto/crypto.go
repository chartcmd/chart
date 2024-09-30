package crypto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/chartcmd/chart/constants"
)

var (
	data [][]float64
)

func GetCoinbaseCandlestick(ticker string, start, end time.Time, granularity uint32) ([][]float64, error) {
	url := fmt.Sprintf(constants.CoinbaseCandleEndpointFullUrl,
		constants.CoinbaseBaseUrl,
		fmt.Sprintf(constants.CoinbaseCandleEndpointUrl, ticker),
		start.Format(time.RFC3339),
		end.Format(time.RFC3339),
		granularity,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
