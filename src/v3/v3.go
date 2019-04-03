package coingecko

import (
	"encoding/json"
	"fmt"
	"strings"

	helper "github.com/superoo7/go-gecko/src/helper"
	types "github.com/superoo7/go-gecko/src/v3/types"
)

var baseURL = "https://api.coingecko.com/api/v3"

// Ping /ping endpoint
func Ping() (*types.Ping, error) {
	url := fmt.Sprintf("%s/ping", baseURL)
	resp, err := helper.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.Ping
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SimpleSinglePrice /simple/price (id, vs_currency)
func SimpleSinglePrice(cID string, vC string) (*types.SimpleSinglePrice, error) {
	cID = strings.ToLower(cID)
	vC = strings.ToLower(vC)
	params := []string{fmt.Sprintf("ids=%s", cID), fmt.Sprintf("vs_currencies=%s", vC)}
	url := fmt.Sprintf("%s/simple/price?%s", baseURL, strings.Join(params, "&"))
	resp, err := helper.MakeReq(url)
	if err != nil {
		return nil, err
	}
	t := make(map[string]map[string]float32)
	err = json.Unmarshal(resp, &t)
	if err != nil {
		return nil, err
	}
	c := t[cID]
	data := &types.SimpleSinglePrice{ID: cID, Currency: vC, MarketPrice: c[vC]}
	return data, nil
}

// SimpleSupportedVSCurrencies /simple/supported_vs_currencies
func SimpleSupportedVSCurrencies() (*types.SimpleSupportedVSCurrencies, error) {
	url := fmt.Sprintf("%s/simple/supported_vs_currencies", baseURL)
	resp, err := helper.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.SimpleSupportedVSCurrencies
	err = json.Unmarshal(resp, &data)
	return data, nil
}

// Global https://api.coingecko.com/api/v3/global
func Global() (*types.Global, error) {
	url := fmt.Sprintf("%s/global", baseURL)
	resp, err := helper.MakeReq(url)
	if err != nil {
		return nil, err
	}
	var data *types.GlobalResponse
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data.Data, nil
}
