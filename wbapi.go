package gowd

import (
	"io"
	"net/url"
)

func (client *Client) CallWb(params map[string][]string) (io.ReadCloser, string, error) {
	parameters := url.Values{}
	for key,values := range params {
		for _, value := range values {
			parameters.Set(key,value);
		}
	}

	url := wdapi_url+"?"+parameters.Encode()
	resp, err := client.Get(url)
	
	return resp.Body, url, err
}
