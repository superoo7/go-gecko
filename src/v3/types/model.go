package types

// OrderType

// OrderType in CoinGecko
type OrderType struct {
	MarketCapDesc string
	MarketCapAsc  string
	GeckoDesc     string
	GeckoAsc      string
	VolumeAsc     string
	VolumeDesc    string
}

// OrderTypeObject for certain order
var OrderTypeObject = &OrderType{
	MarketCapDesc: "market_cap_desc",
	MarketCapAsc:  "market_cap_asc",
	GeckoDesc:     "gecko_desc",
	GeckoAsc:      "gecko_asc",
	VolumeAsc:     "volume_asc",
	VolumeDesc:    "volume_desc",
}

// PriceChangePercentage

// PriceChangePercentage in different amount of time
type PriceChangePercentage struct {
	PCP1h   string
	PCP24h  string
	PCP7d   string
	PCP14d  string
	PCP30d  string
	PCP200d string
	PCP1y   string
}

// PriceChangePercentageObject for different amount of time
var PriceChangePercentageObject = &PriceChangePercentage{
	PCP1h:   "1h",
	PCP24h:  "24h",
	PCP7d:   "7d",
	PCP14d:  "14d",
	PCP30d:  "30d",
	PCP200d: "200d",
	PCP1y:   "1y",
}

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

// SHARED
// coinBaseStruct [private]
type coinBaseStruct struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// CoinsListItem item in CoinList
type CoinsListItem struct {
	coinBaseStruct
}

// CoinsMarketItem item in CoinMarket
type CoinsMarketItem struct {
	coinBaseStruct
	Image                        string  `json:"image"`
	CurrentPrice                 float64 `json:"current_price"`
	MarketCap                    float64 `json:"market_cap"`
	MarketCapRank                int16   `json:"market_cap_rank"`
	TotalVolume                  float64 `json:"total_volume"`
	High24                       float64 `json:"high_24h"`
	Low24                        float64 `json:"low_24h"`
	PriceChange24h               float64 `json:"price_change_24h"`
	PriceChangePercentage24h     float64 `json:"price_change_percentage_24h"`
	MarketCapChange24h           float64 `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float64 `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            string  `json:"circulating_supply"`
	TotalSupply                  float64 `json:"total_supply"`
	ATH                          float64 `json:"ath"`
	ATHChangePercentage          float64 `json:"ath_change_percentage"`
	ATHDate                      string  `json:"ath_date"`
	ROI                          *struct {
		Times      float64 `json:"times"`
		Currency   string  `json:"currency"`
		Percentage float64 `json:"percentage"`
	} `json:"roi"`
	LastUpdated   string `json:"last_updated"`
	SparklineIn7d *struct {
		Price []float64 `json:"price"`
	} `json:"sparkline_in_7d"`
	PriceChangePercentage1hInCurrency   *float64 `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency  *float64 `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7dInCurrency   *float64 `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14dInCurrency  *float64 `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage30dInCurrency  *float64 `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage200dInCurrency *float64 `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1yInCurrency   *float64 `json:"price_change_percentage_1y_in_currency"`
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
