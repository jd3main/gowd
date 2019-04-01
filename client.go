package gowd

import (
    "net/http"
)

const (
    wdapi_url       = "https://www.wikidata.org/w/api.php"
    entitydata_url  = "https://www.wikidata.org/wiki/Special:EntityData/"
)



// Client with single connection to access Wikidata's API
type Client struct {
    // embedded http.Client 
    http.Client

    // Format specified the prefered format to accept from wikidata
    AcceptFormat    string

    // langFallback is a list of prefered langauges
    // languages with lower index has higher priorirty
    LangFallback    []string
}
