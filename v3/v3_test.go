package coingecko

import (
	"io/ioutil"
	"os"
	"testing"

	gock "gopkg.in/h2non/gock.v1"
)

func init() {
	defer gock.Off()
}

var c = NewClient(nil)
var mockURL = "https://api.coingecko.com/api/v3"

func TestPing(t *testing.T) {
	err := setupGock("json/ping.json", "/ping")
	ping, err := c.Ping()
	if err != nil {
		t.FailNow()
	}
	if ping.GeckoSays != "(V3) To the Moon!" {
		t.FailNow()
	}
}

func TestSimpleSinglePrice(t *testing.T) {
	err := setupGock("json/simple_single_price.json", "/simple/price")
	if err != nil {
		t.FailNow()
	}
	simplePrice, err := c.SimpleSinglePrice("bitcoin", "usd")
	if err != nil {
		t.FailNow()
	}
	if simplePrice.ID != "bitcoin" || simplePrice.Currency != "usd" || simplePrice.MarketPrice != float32(5013.61) {
		t.FailNow()
	}
}

func TestSimplePrice(t *testing.T) {
	err := setupGock("json/simple_price.json", "/simple/price")
	if err != nil {
		t.FailNow()
	}
	ids := []string{"bitcoin", "ethereum"}
	vc := []string{"usd", "myr"}
	sp, err := c.SimplePrice(ids, vc)
	if err != nil {
		t.FailNow()
	}
	bitcoin := (*sp)["bitcoin"]
	eth := (*sp)["ethereum"]

	if bitcoin["usd"] != 5005.73 || bitcoin["myr"] != 20474 {
		t.FailNow()
	}
	if eth["usd"] != 163.58 || eth["myr"] != 669.07 {
		t.FailNow()
	}
}

func TestSimpleTokenPrice(t *testing.T) {
	err := setupGock("json/simple_token_price.json", "/simple/token_price")
	if err != nil {
		t.FailNow()
	}

	platformID := "ethereum"
	contractAddresses := []string{
		"0x5218E472cFCFE0b64A064F055B43b4cdC9EfD3A6",
		"0x6b175474e89094c44da98b954eedeac495271d0f",
		"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
	}
	currencies := []string{"usd"}

	tokens, err := c.SimpleTokenPrice(platformID, contractAddresses, currencies, nil)
	if err != nil {
		t.FailNow()
	}

	if tokens == nil {
		t.FailNow()
	}

	if len(tokens) != 3 {
		t.FailNow()
	}
}

func TestSimpleSupportedVSCurrencies(t *testing.T) {
	err := setupGock("json/simple_supported_vs_currencies.json", "/simple/supported_vs_currencies")
	s, err := c.SimpleSupportedVSCurrencies()
	if err != nil {
		t.FailNow()
	}
	if len(*s) != 54 {
		t.FailNow()
	}
}

func TestCoinsList(t *testing.T) {
	err := setupGock("json/coins_list.json", "/coins/list")
	list, err := c.CoinsList()
	if err != nil {
		t.FailNow()
	}
	item := (*list)[0]
	if item.ID != "01coin" {
		t.FailNow()
	}
}

// Util: Setup Gock
func setupGock(filename string, url string) error {
	testJSON, err := os.Open(filename)
	defer testJSON.Close()
	if err != nil {
		return err
	}
	testByte, err := ioutil.ReadAll(testJSON)
	if err != nil {
		return err
	}
	gock.New(mockURL).
		Get(url).
		Reply(200).
		JSON(testByte)

	return nil
}
