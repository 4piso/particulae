package iexcloud

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Global properties
const (
	BaseURL              = "https://cloud.iexapis.com/v1"
	MaxConnectionPerHost = 10
	RequestTimeout       = 2
)

// Client handler the client for iexcloud rest API.
type Client struct {
	// httpClient handle the http client configuration
	httpClient *http.Client
}

// Get is a wrapper of the http.Get method but for making the request
// NOTE: (c Client) without pointer just means you cannot modify the reciever
// build the iexURL using the BaseURL and some endpoint. example https://cloud.iexapis.com/v1/stocks/apple
// Query the params from the URL
// grap the params for url params and add it to the url
func (c Client) Get(ctx context.Context, endpoint string, params map[string]string, body io.Reader) (*http.Response, error) {
	iexURL := BaseURL + endpoint
	req, err := http.NewRequestWithContext(ctx, "GET", iexURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not make the request: %v ", err)
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(v, k)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}

// NewClient constructor function
func NewClient() *Client {
	return &Client{
		httpClient: createHTTPClient(),
	}
}

// createHTTPClient private method that return the http.Client
func createHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: MaxConnectionPerHost,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
}
