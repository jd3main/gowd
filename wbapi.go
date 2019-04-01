package gowd

import (
	"io"
	"net/url"
)

func (client *Client) callWb(action string, params map[string]string) (io.ReadCloser, string, error) {
	parameters := url.Values{}
	for key,value := range params {
		parameters.Set(key,value);
	}

	url := wdapi_url+"?"+parameters.Encode()
	resp, err := client.Get(url)
	
	return resp.Body, url, err
}
