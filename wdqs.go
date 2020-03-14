package gowd

import (
	"io"
	"net/url"
)

// SparqlQuery queries data from Wikidata Query Service with SPARQL code
// parameter ql is a string containing your SPARQL code
func (client *Client) SparqlQuery(ql string) (content io.ReadCloser, err error) {
	if ql == "" {
		return nil, nil
	}

	reqURL := WdqsEndPoint +
		"?format=" + client.AcceptFormat +
		"&query=" + url.PathEscape(ql)
	resp, err := client.Get(reqURL)
	content = resp.Body
	return
}
