package iexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// global properties
const (
	baseURL              = "https://cloud.iexapis.com/v1"
	RequestTimeout       = 2
	MaxConnectionPerHost = 10
)

// Client handler comunication to the iexapis
type Client struct {
	// httpClient is the http.Client wrapper
	httpClient *http.Client
}

// NewClient is the constructor function
func NewClient() *Client {
	return &Client{
		httpClient: createHTTPClient(),
	}
}

// Get is a wrapper for the http.Get to get more options
// NOTE: (c Client) without pointer just means you cannot modify the receiver
// (c *Client) has a pointer so you can mutate the public variables on the receiver c
func (c Client) Get(ctx context.Context, endpoint string, params map[string]string, body io.Reader) (*http.Response, error) {
	iexURL := baseURL + endpoint
	req, err := http.NewRequestWithContext(ctx, "GET", iexURL, nil)
	if err != nil {
		return nil, fmt.Errorf("could not build the request: %v ", err)
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}

// Quote is going to return the quote base on the services url
// GET /stock/{symbol}/quote/
func (c *Client) Quote(ctx context.Context, symbol string, displayPercent bool) (*Quote, error) {
	endpoint := "stock/" + symbol + "/quote"
	quote := new(Quote)
	if displayPercent {
		endpoint = endpoint + "?displayPercent=true"
	}
	res, err := c.Get(ctx, endpoint, nil, nil)
	if err != nil {
		return quote, fmt.Errorf("could not make the quote: %v ", err)
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&quote); err != nil {
		return quote, fmt.Errorf("could not decode the quote struct: %v ", err)
	}
	return quote, nil
}

// createHTTPClient is a private method to creat the client
func createHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: MaxConnectionPerHost,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
}
