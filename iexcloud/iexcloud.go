package iexcloud

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://cloud.iexapis.com/v1"
)

// Client manages communications with iexcloud
type Client struct {
	// client will handler connections
	httpClient *http.Client

	// baseURl for the api request. Default to the iexcloud API
	baseURL string

	// token definition for the credentials
	Token string

	// version of the api
	version string
}

// Error represent the an IEX API error
type Error struct {
	StatusCode int
	Message    string
}

// Error implementes the error interface
func (e *Error) Error() string {
	return fmt.Sprintf("%d %s : %s ", e.StatusCode, http.StatusText(e.StatusCode), e.Message)
}

// ClientOption is a func that operates the *Client
type ClientOption func(*Client) error

// NewClient contructor function
// Apply options for middlewear
func NewClient(ctx context.Context, token string, options ...ClientOption) (*Client, error) {
	client := &Client{
		Token:      token,
		httpClient: &http.Client{},
	}
	for _, option := range options {
		err := option(client)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

// WithHTTPClient set the http.Client for a new iexcloud Client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = httpClient
		return nil
	}
}

// WithBaseURL set the baseURL for a new iexcooud Client
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

// addToken add the tokens to the endpoint
func (c *Client) addToken(endpoint string) (string, error) {
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return "", nil
	}
	v := u.Query()
	v.Add("token", c.Token)
	u.RawQuery = v.Encode()
	return u.String(), nil
}
