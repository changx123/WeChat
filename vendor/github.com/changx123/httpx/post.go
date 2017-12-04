package httpx

import (
	"net/http"
	"bytes"
)

//简单的POST
func Post(url string, contentType string, b []byte) (*Response, error) {
	resp, err := http.Post(url, contentType, bytes.NewReader(b))
	if err != nil {
		return &Response{}, err
	}
	defer resp.Body.Close()
	var response Response
	err = response.response(resp)
	return &response, err
}

//httpx对象Post
func (httpx *Httpx) Post(url string, body []byte) (*Response, error) {
	response, err := httpx.Request(POST, url, body)
	if err != nil {
		return response, err
	}
	return response, nil
}
