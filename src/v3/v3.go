package coingecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	types "github.com/superoo7/go-gecko/src/v3/types"
)

var baseURL = "https://api.coingecko.com/api/v3"

// Ping /ping endpoint
func Ping() (*types.Ping, error) {
	url := fmt.Sprintf("%s/ping", baseURL)
	resp, err := makeReq(url)
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
	resp, err := makeReq(url)
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

// helper
// doReq HTTP client
func doReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

// makeReq HTTP request helper
func makeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := doReq(req)
	if err != nil {
		return nil, err
	}
	return resp, err
}
