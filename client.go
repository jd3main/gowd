package gowd

import (
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
)

// Client is a wrapper of http.Client with some extra configurations for accessing Wikidata's APIs
type Client struct {
	// embedded http.Client
	http.Client

	// AcceptFormat specified the prefered format to accept from Wikidata
	AcceptFormat string

	// Languages is a list of prefered langauges
	Languages []string

	// ServerAddr is an URL specifying the mediawiki site to connect.
	ServerAddr string

	// If NoEditMode is set to true, queries will be output to QueryLogger only
	NoEditMode bool

	// All Queries will be output to QueryLogger.
	// Set this to nil to turn off output logging.
	QueryLogger io.Writer
}

// Create a default client for accessing Wikidata with default settings.
// Please make sure your program been tested before using this client.
// For safety, Actionless of this client is set to true.
func NewClient() *Client {
	client := &Client{
		AcceptFormat: "json",
		ServerAddr:   WdUrl,
		NoEditMode:   true,
	}
	client.Jar, _ = cookiejar.New(nil)

	logFile, err := os.OpenFile("gowd_query.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Print("[Warning] QueryLogger is not opened.")
	} else {
		client.QueryLogger = logFile
	}

	return client
}

// The default test client almost the same as DefaultClient.
// Except that ServerAddr set to the test site (test.wikidata.org) and the actionless mode is off.
// This is recommanded for developing bots.
func NewTestClient() *Client {
	client := &Client{
		AcceptFormat: "json",
		ServerAddr:   TestWdUrl,
		NoEditMode:   false,
	}
	client.Jar, _ = cookiejar.New(nil)

	return client
}
