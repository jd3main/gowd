package gowd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

// GetToken send a `action=query&meta=tokens` request to get token for editing/login
func (client *Client) GetToken(tokenType string) (string, error) {
	queryUrl := client.ServerAddr + MwApiPath +
		"?format=json&action=query&meta=tokens"
	if tokenType != "" {
		queryUrl += "&type=" + tokenType
	}

	resp, err := client.Get(queryUrl)
	if err != nil {
		return "", err
	}

	jsonBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Query struct {
			Tokens map[string]string
		}
	}
	json.Unmarshal(jsonBody, &result)

	for key, token := range result.Query.Tokens {
		if strings.Contains(key, "token") {
			if token == "+\\" {
				return "", errors.New("Got empty token")
			}
			return token, nil
		}
	}
	return "", errors.New("Failed to get token")
}
