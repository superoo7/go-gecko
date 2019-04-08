package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

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

// MakeReq HTTP request helper
func MakeReq(url string) ([]byte, error) {
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

// Bool2String boolean to string
func Bool2String(b bool) string {
	return strconv.FormatBool(b)
}

// Int2String Integer to string
func Int2String(i int) string {
	return strconv.Itoa(i)
}
