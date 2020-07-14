package iexcloud

import (
	"net"
	"net/http"
	"net/url"
	"strings"
)

// general properties
const (
	defaultBaseURL = "https://cloud.iexapis.com/v1"
)

// Client manage the comunication to iexcloud
type Client struct {
	// HTTP client to call the api
	client *http.Client

	// baseURL if the make the url client base for the api
	baseURL *url.URL

	// conn handler the connetios of the project
	conn net.Conn
}

// SetBaseURL sets the ase URL for API request to a custom ednpoints. urlStr
// shoudl always b specified with a trailing slash
func (c *Client) SetBaseURL(urlStr string) error {
	if !strings.HasSuffix(urlStr, "/") {
		urlStr += "/"
	}
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	c.baseURL = baseURL
	return nil
}
