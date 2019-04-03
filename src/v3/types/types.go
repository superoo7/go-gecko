package types

// Ping https://api.coingecko.com/api/v3/ping
type Ping struct {
	GeckoSays string `json:"gecko_says"`
}

// SimpleSinglePrice https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd
type SimpleSinglePrice struct {
	ID          string
	Currency    string
	MarketPrice float32
}

// SimpleSupportedVSCurrencies https://api.coingecko.com/api/v3/simple/supported_vs_currencies
type SimpleSupportedVSCurrencies []string

// CoinList https://api.coingecko.com/api/v3/coins/list
type CoinList []CoinsListItem

// EventsTypes https://api.coingecko.com/api/v3/events/types
type EventsTypes struct {
	Data  []string `json:"data"`
	Count uint16   `json:"count"`
}

// ExchangeRatesResponse https://api.coingecko.com/api/v3/exchange_rates
type ExchangeRatesResponse struct {
	Rates ExchangeRatesItem `json:"rates"`
}

// GlobalResponse https://api.coingecko.com/api/v3/global
type GlobalResponse struct {
	Data Global `json:"data"`
}
