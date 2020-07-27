package cfg

import (
	"fmt"
	"net/url"
	"sync"
)

// Config set the configuration from different source: file, environemnt variable and docker secrets
type Config struct {
	m  map[string]string
	mu sync.RWMutex
}

// Provider is implement by the user to provide the configuration as one map
type Provider interface {
	Provide() (map[string]string, error)
}

// New constructor function from Provider
func New(p Provider) (*Config, error) {
	m, err := p.Provide()
	if err != nil {
		return nil, err
	}
	c := &Config{
		m: m,
	}
	return c, nil
}

// String return the value of give key as string, it will return the error if not found
func (c *Config) String(key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.m[key]
	if !ok {
		return "", fmt.Errorf("Unknow key: %v ", key)
	}
	return value, nil
}

// MustString return a value of given key as string.
// it wil panic is there is na error
func (c *Config) MustString(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.m[key]
	if !ok {
		panic(fmt.Sprintf("Unknow key %s : ", key))
	}
	return value
}

// URL returns the value of the given key as URL, it will return an error
// if the key was not found
func (c *Config) URL(key string) (*url.URL, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.m[key]
	if !ok {
		return nil, fmt.Errorf("unknow key : %s ", key)
	}
	u, err := url.Parse(value)
	if err != nil {
		return u, err
	}
	return u, nil
}
