package crypto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

type CoinbaseProduct struct {
	ID            string `json:"id"`
	BaseCurrency  string `json:"base_currency"`
	QuoteCurrency string `json:"quote_currency"`
}

func getCoinbaseTradingPairs() ([]CoinbaseProduct, error) {
	url := fmt.Sprintf(
		CoinbaseProductsF,
		CoinbaseBaseUrl1,
		CoinbaseProductsEndpointUrl,
	)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var products []CoinbaseProduct
	err = json.Unmarshal(body, &products)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return products, nil
}

func GetCryptoList() []string {
	products, err := getCoinbaseTradingPairs()
	if err != nil {
		return nil
	}

	cryptos := []string{}
	for _, product := range products {
		if strings.HasSuffix(product.ID, "-USD") {
			cryptos = append(cryptos, strings.TrimSuffix(product.ID, "-USD"))
		}
	}
	sort.Strings(cryptos)
	return cryptos
}
