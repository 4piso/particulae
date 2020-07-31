package iexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

	// token is the token from the client
	Token string
}

// NewClient is the constructor function
func NewClient(token string) *Client {
	return &Client{
		httpClient: createHTTPClient(),
		Token:      token,
	}
}

// Quote is going to return the quote base on the services url
// GET /stock/{symbol}/quote
func (c *Client) Quote(ctx context.Context, symbol string, displayPercent bool) (*Quote, error) {
	endpoint := "/stock/" + symbol + "/quote"
	quote := new(Quote)
	params := make(map[string]string)
	if displayPercent {
		params = map[string]string{
			"displayPercent": "true",
		}
	}
	res, err := c.Get(ctx, endpoint, params, nil)
	if err != nil {
		log.Fatalf("Error sending the request : %v ", err)
		return quote, nil
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&quote); err != nil {
		log.Fatalf("Error decoding the response: %v ", err)
		return quote, err
	}
	return quote, nil
}

// Earnings return the eanings base on the service url
// GET /stock/{symbol}/earnings/
func (c *Client) Earnings(ctx context.Context, symbol string) (*Earnings, error) {
	endpoint := "/stock/" + symbol + "/earnings"
	earning := new(Earnings)
	res, err := c.Get(ctx, endpoint, nil, nil)
	if err != nil {
		log.Fatalf("Error doing the request: %v ", err)
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&earning); err != nil {
		log.Fatalf("Errror decodeing json: %v ", err)
		return earning, nil
	}
	return earning, nil
}

// EarningsToday is going to return the earnig from this day
// GET /stock/market/today-earnings
func (c *Client) EarningsToday(ctx context.Context) (*EarningsToday, error) {
	endpoint := "/stock/market/today-earnings"
	earningsToday := new(EarningsToday)
	res, err := c.Get(ctx, endpoint, nil, nil)
	if err != nil {
		log.Fatalf("earning today request failed: %v ", err)
		return earningsToday, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&earningsToday); err != nil {
		log.Fatalf("cannot parse the response: %v ", err)
		return earningsToday, err
	}
	return earningsToday, nil
}

// HELPERS METHODS
// ADDING FOR USE BY OTHER METHOD

// GetToken return the token
func (c *Client) GetToken() string {
	return c.Token
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
	q.Set("token", c.GetToken())
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
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
