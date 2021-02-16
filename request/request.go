package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

// helper
// doReq HTTP client
func doReq(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil {
		// workaround for possible net.http memory leak, ref:
		// http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#close_http_resp_body
		defer resp.Body.Close()
	}
	if err != nil {
		if resp != nil {
			// workaround for possible issue when we want to reuse the
			// underlying connection, the resp.Body should be read in every
			// possible path
			io.Copy(ioutil.Discard, resp.Body)
		}
		return nil, err
	}
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
