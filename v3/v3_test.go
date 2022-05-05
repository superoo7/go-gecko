package coingecko

import (
	"fmt"
	"io/ioutil"
	"log"
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

func TestCoinsIDOHLC(t *testing.T) {
	const invalidDaysErr = "invalid days value"
	// tests were ran at 3:00 PM PST
	// Actual response length may vary slightly by time of day / availability of current day's data sets
	var tests = []struct {
		Id                 string
		Currency           string
		Days               string
		wantErr            error
		WantResponseLength int
	}{
		{
			Id:                 "bitcoin",
			Currency:           "usd",
			Days:               "1",
			wantErr:            nil,
			WantResponseLength: 48, // 1d = 24 * 0.5
		},
		{
			Id:                 "ethereum",
			Currency:           "usd",
			Days:               "7",
			wantErr:            nil,
			WantResponseLength: 43, // 7d = (24 / 4) * 7 + 1
		},
		{
			Id:                 "ethereum",
			Currency:           "usd",
			Days:               "14",
			wantErr:            nil,
			WantResponseLength: 85, // 14d = (24 / 4) * 14 + 1
		},
		{
			Id:                 "bitcoin",
			Currency:           "usd",
			Days:               "30",
			wantErr:            nil,
			WantResponseLength: 181, // 30d = (24 / 4) * 30 + 1
		},
		{
			Id:                 "bitcoin",
			Currency:           "usd",
			Days:               "90",
			wantErr:            nil,
			WantResponseLength: 24, // 90d = (90 / 4) + 2
		},
		{
			Id:                 "bitcoin",
			Currency:           "usd",
			Days:               "180",
			wantErr:            nil,
			WantResponseLength: 48, // 180d = ((90 / 4) + 2) * 2
		},
		{
			Id:                 "bitcoin",
			Currency:           "usd",
			Days:               "365",
			wantErr:            nil,
			WantResponseLength: 96, // 365d = ((90 / 4) + 2) * 4
		},
		{
			Id:                 "bitcoin",
			Currency:           "usd",
			Days:               "2",
			wantErr:            fmt.Errorf(invalidDaysErr),
			WantResponseLength: 1, // 365d = ((90 / 4) + 2) * 4
		},
	}
	c := NewClient(nil)

	for _, test := range tests {
		chart, err := c.CoinsIDOHLC(test.Id, test.Currency, test.Days)
		if err != nil {
			if test.wantErr == nil {
				t.Errorf("FAIL - ERROR: %v; WANT: nil", err)
			}
			if err.Error() != test.wantErr.Error() {
				t.Errorf("FAIL - ERROR: %v; WANT: %v", err, test.wantErr)
			}
			continue // no chart to check length against
		}

		if len(*chart) != test.WantResponseLength {
			t.Errorf("Invalid response length: %d; WANT: %d", len(*chart), test.WantResponseLength)
		}
		log.Printf("CHART:\n %v", *chart)
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
