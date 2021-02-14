package wbutils

import (
	"net/url"
	"strconv"

	"github.com/jd3main/gowd"
)

// SearchEntities constructs parameters for searching entities by name
// If limit is set to a negative value, then the limit will be set to max, which is currently 50 determind by Wikidata.
// If limit ist set to nil, then the default value (currently 7) is used.
// The parameter cont indicates how many entities will be skipped.
func SearchEntities(name string, languageCode string, limit *int, cont *int64) gowd.Parameters {
	parameters := url.Values{}
	parameters.Set("action", "wbsearchentities")
	parameters.Set("search", name)
	if limit != nil && *limit >= 0 {
		parameters.Set("limit", strconv.Itoa(*limit))
	}
	if cont != nil && *cont >= 0 {
		parameters.Set("continue", strconv.FormatInt(*cont, 10))
	}
	parameters.Set("language", languageCode)
	return parameters
}
