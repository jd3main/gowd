package gowd

import (
	"io"
	"io/ioutil"
	"net/url"
	"strings"
)

func (client *Client) MwApiGetUrl(params Parameters) string {
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

	return client.ApiEndpointUrl() + "?" + parameters.Encode()
}

// MwApiGet directly calls the Mediawiki API through http GET method.
func (client *Client) MwApiGet(params Parameters) (reader io.ReadCloser, fullUrl string, err error) {
	fullUrl = client.MwApiGetUrl(params)
	client.logQuery(fullUrl, nil)
	resp, err := client.Get(fullUrl)
	return resp.Body, fullUrl, err
}

// MwApiPost calls the Mediawiki API through http POST method.
// This function handles the token required by editting and loginning.
// Posted parameters will be included in fullUrl like parameters of get method.
// When NoEditMode is set, the returned reader will be an empty ReaderCloser instead of nil
func (client *Client) MwApiPost(params Parameters, tokenType string) (reader io.ReadCloser, fullUrl string, err error) {

	token, err := client.GetToken(tokenType)
	if err != nil {
		return nil, "", err
	}

	postParams := url.Values{}
	postParams.Set("token", token)
	for key, values := range params {
		for _, value := range values {
			postParams.Set(key, value)
		}
	}

	urlStr := client.ApiEndpointUrl()

	if client.AcceptFormat != "" {
		postParams.Set("format", client.AcceptFormat)
	} else {
		postParams.Set("format", "json")
	}

	client.logQuery(urlStr, postParams)
	if client.NoEditMode {
		return ioutil.NopCloser(strings.NewReader("")), "", nil
	} else {
		resp, err := client.PostForm(urlStr, postParams)
		fullUrl = client.ApiEndpointUrl() + "?" + postParams.Encode()
		return resp.Body, fullUrl, err
	}
}

func (client *Client) ApiEndpointUrl() string {
	return client.ServerAddr + MwApiPath
}
