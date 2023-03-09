package bybit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Bybit struct {
	client  *http.Client
	baseURL string
}

const (
	MAIN_NET = "https://api.bybit.com"
	TEST_NET = "https://api-testnet.bybit.com"
)

func New(options ...BybitOption) *Bybit {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	bybit := &Bybit{
		client:  client,
		baseURL: MAIN_NET,
	}
	for _, opt := range options {
		opt(bybit)
	}

	return bybit
}

func (b *Bybit) PublicRequest(
	method string,
	apiURL string,
	params map[string]interface{},
	result interface{},
) (fullURL string, resp []byte, err error) {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var p []string
	for _, k := range keys {
		p = append(p, fmt.Sprintf("%v=%v", k, params[k]))
	}

	param := strings.Join(p, "&")
	fullURL = b.baseURL + apiURL
	if param != "" {
		fullURL += "?" + param
	}
	var binBody = bytes.NewReader(make([]byte, 0))

	// get a http request
	var request *http.Request
	request, err = http.NewRequest(method, fullURL, binBody)
	if err != nil {
		return
	}

	var response *http.Response
	response, err = b.client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, result)
	return
}
