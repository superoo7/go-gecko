package coingecko

import (
	"testing"
)

func TestPing(t *testing.T) {
	ping, err := Ping()

	if err != nil {
		t.FailNow()
	}

	if ping.GeckoSays != "(V3) To the Moon!" {
		t.FailNow()
	}
}

func TestSimpleSinglePrice(t *testing.T) {
	price, err := SimpleSinglePrice("bitcoin", "usd")

	if err != nil {
		t.FailNow()
	}
}
