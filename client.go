package gowd

import (
	"net/http"
)

// Client is a wrapper of http.Client with some extra configurations for accessing Wikidata's APIs
type Client struct {
	// embedded http.Client
	http.Client

	// AcceptFormat specified the prefered format to accept from Wikidata
	AcceptFormat string

	// Languages is a list of prefered langauges
	Languages []string
}
