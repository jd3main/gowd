package gowd

import (
	"io"
	"net/url"
	"strings"
)

func (client *Client) GetEntities(ids []string) (io.ReadCloser, string, error) {
	parameters := url.Values{}
	parameters.Set("action","wbgetentities");
	parameters.Set("format",client.AcceptFormat);
	parameters.Set("ids", strings.Join(ids,"|"));

	url := wdapi_url+"?"+parameters.Encode()
	resp, err := client.Get(url)
	
	return resp.Body, url, err
}

func (client *Client) GetClaims(ids []string) (io.ReadCloser, string, error) {
	parameters := url.Values{}
	parameters.Set("action","wbgetclaims");
	parameters.Set("format",client.AcceptFormat);
	parameters.Set("ids", strings.Join(ids,"|"));

	url := wdapi_url+"?"+parameters.Encode()
	resp, err := client.Get(url)
	
	return resp.Body, url, err
}
