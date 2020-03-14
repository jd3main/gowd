package gowd

import (
	"net/url"
	"strings"
)

// GetEntities constructs parameters for getting contents of entities with given id
func GetEntities(ids []string) Parameters {
	parameters := url.Values{}
	parameters.Set("action", "wbgetentities")
	parameters.Set("ids", strings.Join(ids, "|"))
	return parameters
}

// GetClaims constructs parameters for getting claims of items with given id
func GetClaims(ids []string) Parameters {
	parameters := url.Values{}
	parameters.Set("action", "wbgetclaims")
	parameters.Set("ids", strings.Join(ids, "|"))
	return parameters
}
