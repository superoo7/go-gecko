# CoinGecko API Client for GoLang

[![Build Status](https://travis-ci.com/superoo7/go-gecko.svg?branch=master)](https://travis-ci.com/superoo7/go-gecko)

Simple API Client for CoinGecko written in Golang

<p align="center">
    <img src="gogecko.png" alt="gogecko" height="200" />
</p>

gopher resources from [free-gophers-pack](https://github.com/MariaLetta/free-gophers-pack)

## Available endpoint

[Refer to CoinGecko official API](https://www.coingecko.com/api)

|            Endpoint             | Status | Testing |            Function            |
| :-----------------------------: | :----: | :-----: | :----------------------------: |
|              /ping              |  [/]   |   [/]   |              Ping              |
|          /simple/price          |  [/]   |   [/]   | SimpleSinglePrice, SimplePrice |
| /simple/supported_vs_currencies |  [/]   |   [/]   |  SimpleSupportedVSCurrencies   |
|           /coins/list           |  [/]   |   [/]   |           CoinsList            |
|          /coins/market          |  [/]   |   [/]   |          CoinsMarket           |
|           /coins/{id}           |  [/]   |         |            CoinsID             |
|       /coins/{id}/history       |  [/]   |         |         CoinsIDHistory         |
|        /events/countries        |  [/]   |         |        EventsCountries         |
|          /events/types          |  [/]   |         |           EventsType           |
|         /exchange_rates         |  [/]   |         |          ExchangeRate          |
|             /global             |  [/]   |         |             Global             |

## Usage

Installation with go get.

```
go get -u github.com/superoo7/go-gecko
```

For usage, checkout [Example folder for v3](/_example/v3)

## License

MIT
