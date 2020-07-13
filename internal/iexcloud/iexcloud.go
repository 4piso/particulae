package iexcloud

import (
	"net/http"
	"net/url"
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
}
