package types

// Ping https://api.coingecko.com/api/v3/ping
type Ping struct {
	GeckoSays string `json:"gecko_says"`
}

// SimplePrice https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd
type SimplePrice struct {
	ID          string
	Currency    string
	MarketPrice string
}
