package crypto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type CoinbaseLatestResponse struct {
	Data struct {
		Amount string `json:"amount"`
	} `json:"data"`
}

func GetCoinbaseLatest(ticker string) (float64, error) {
	url := fmt.Sprintf(
		CoinbaseLatestEndpointF,
		CoinbaseBaseUrl2,
		fmt.Sprintf(CoinbaseLatestEndpointUrl, ticker),
	)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %v", err)
	}

	var coinbaseResp CoinbaseLatestResponse
	err = json.Unmarshal(body, &coinbaseResp)
	if err != nil {
		return 0, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	price, err := strconv.ParseFloat(coinbaseResp.Data.Amount, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing price: %v", err)
	}

	return price, nil
}
