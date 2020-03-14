package gowd

import (
	"io"
	"net/url"
)

// CallMwApi calls Mediawiki API of wikidata.org
func (client *Client) CallMwApi(params Parameters) (io.ReadCloser, string, error) {
	parameters := url.Values{}
	for key, values := range params {
		for _, value := range values {
			parameters.Set(key, value)
		}
	}
	if client.AcceptFormat != "" {
		parameters.Set("format", client.AcceptFormat)
	} else {
		parameters.Set("format", "json")
	}

	url := MwApiUrl + "?" + parameters.Encode()
	resp, err := client.Get(url)

	return resp.Body, url, err
}
