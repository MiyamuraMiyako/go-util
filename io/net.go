package io

import (
	"io/ioutil"
	"net/http"
	"strings"
)

//WebClient is a simple api to send http request and receive response.
type WebClient struct {
	Headers map[string]string
	client  *http.Client
	request *http.Request
}

//NewWebClient new WebClient,default has application/json and utf-8 headers.
func NewWebClient() *WebClient {
	wc := &WebClient{}
	wc.client = &http.Client{}
	wc.Headers = make(map[string]string)
	wc.Headers["Content-Type"] = "application/json"
	wc.Headers["Encoding"] = "UTF-8"
	return wc
}

//Get Send get request to remote
func (wc *WebClient) Get(url string) (string, error) {
	wc.request, _ = http.NewRequest("GET", url, nil)
	rsp, err := wc.client.Do(wc.request)
	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

//Post Send post request to remote
func (wc *WebClient) Post(url, body string) (string, error) {
	wc.request, _ = http.NewRequest("POST", url, strings.NewReader(body))
	rsp, err := wc.client.Do(wc.request)
	if err != nil {
		return "", err
	}

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
