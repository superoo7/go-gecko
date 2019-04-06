package types

// AllCurrencies map all currencies (USD, BTC) to float64
type AllCurrencies map[string]float64

// LocalizationItem map all locale (en, zh) into respective string
type LocalizationItem map[string]string

// DescriptionItem map all description (in locale) into respective string
type DescriptionItem map[string]string

// LinksItem map all links
type LinksItem map[string][]string

// MarketDataItem map all market data item
type MarketDataItem struct {
	CurrentPrice        AllCurrencies     `json:"current_price"`
	ROI                 AllCurrencies     `json:"roi"`
	ATH                 AllCurrencies     `json:"ath"`
	ATHChangePercentage AllCurrencies     `json:"ath_change_percentage"`
	ATHDate             map[string]string `json:"ath_date"`
	MarketCap           AllCurrencies     `json:"market_cap"`
	MarketCapRank       uint16            `json:"market_cap_rank"`
	TotalVolume         AllCurrencies     `json:"total_volume"`
	High24              AllCurrencies     `json:"high_24h"`
	Low24               AllCurrencies     `json:"low_24h"`
}

// ImageItem struct for all sizes of image
type ImageItem struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

// Type Specific types

// CoinsListItem item in CoinList
type CoinsListItem struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// EventCountryItem item in EventsCountries
type EventCountryItem struct {
	Country string `json:"country"`
	Code    string `json:"code"`
}

// ExchangeRatesItem item in ExchangeRate
type ExchangeRatesItem map[string]ExchangeRatesItemStruct

// ExchangeRatesItemStruct struct in ExchangeRateItem
type ExchangeRatesItemStruct struct {
	Name  string  `json:"name"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

// Global for data of /global
type Global struct {
	ActiveCryptocurrencies          uint16        `json:"active_cryptocurrencies"`
	UpcomingICOs                    uint16        `json:"upcoming_icos"`
	EndedICOs                       uint16        `json:"ended_icos"`
	Markets                         uint16        `json:"markets"`
	MarketCapChangePercentage24hUSD float32       `json:"market_cap_change_percentage_24h_usd"`
	TotalMarketCap                  AllCurrencies `json:"total_market_cap"`
	TotalVolume                     AllCurrencies `json:"total_volume"`
	MarketCapPercentage             AllCurrencies `json:"market_cap_percentage"`
	UpdatedAt                       int64         `json:"updated_at"`
}
