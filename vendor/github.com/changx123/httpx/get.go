package httpx

import "net/http"

//简单的Get
func Get(url string) (*Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return &Response{}, err
	}
	defer resp.Body.Close()
	var response Response
	err = response.response(resp)
	return &response, err
}

//httpx对象Get
func (httpx *Httpx) Get(url string) (*Response, error) {
	response, err := httpx.Request(GET, url, []byte(""))
	if err != nil {
		return response, err
	}
	return response, nil
}
