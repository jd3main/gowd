package gowd

import (
	"log"
	"net/url"
	"time"
)

// Private method for client to log queries to QueryLogger
// When QueryLogger is nil, this function do nothing.
func (client *Client) logQuery(url string, data url.Values) {
	if client.QueryLogger == nil {
		return
	}
	timeStamp := "[" + time.Now().Format(time.RFC3339) + "]"
	logStr := timeStamp + "\t" + url + "\t" + data.Encode() + "\n"
	_, err := client.QueryLogger.Write([]byte(logStr))
	if err != nil {
		log.Println("[Warning] Failed to write query log")
	}
}
