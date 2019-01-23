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
