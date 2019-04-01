package gowd

import (
    "io"
    "net/url"
)

const WDQS_EndPoint = "https://query.wikidata.org/sparql"

// SparqlQuery queries data from Wikidata Query Service with SPARQL code
// parameter ql is a string containing your SPARQL code
func (client *Client) SparqlQuery(ql string) (content io.ReadCloser, err error) {
    if ql == "" {
        return nil,nil
    }

    req_url := WDQS_EndPoint +  
                "?format=" + client.AcceptFormat +
                "&query=" + url.PathEscape(ql)
    resp, err := client.Get(req_url)
    content = resp.Body
    return
}