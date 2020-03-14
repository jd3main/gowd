package gowd

import (
	"net/url"
)

// SearchEntities constructs parameters for searching entities by name
func SearchEntities(name string, languageCode string) Parameters {
	parameters := url.Values{}
	parameters.Set("action", "wbsearchentities")
	parameters.Set("search", name)
	parameters.Set("language", languageCode)
	return parameters
}
