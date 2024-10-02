package crypto

var (
	CoinbaseBaseUrl1            = "https://api.exchange.coinbase.com"
	CoinbaseBaseUrl2            = "https://api.coinbase.com/v2"
	CoinbaseCandleEndpointUrl   = "/products/%s/candles"
	CoinbaseCandleEndpointF     = "%s%s?start=%s&end=%s&granularity=%d"
	CoinbaseLatestEndpointUrl   = "/prices/%s/spot"
	CoinbaseLatestEndpointF     = "%s%s"
	CoinbaseProductsEndpointUrl = "/products"
	CoinbaseProductsF           = "%s%s"
)
