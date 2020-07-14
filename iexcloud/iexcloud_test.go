package iexcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/joshvoll/particulae/internal/cfg"
)

func TestMakeRequest(t *testing.T) {
	ctx := context.Background()
	file := cfg.FileProvider{
		Filename: "config.json",
	}
	config, err := cfg.New(file)
	if err != nil {
		t.Fatalf("Error opening file: %v ", err)
	}
	client, err := NewClient(ctx, config.MustString("token"), WithBaseURL(config.MustString("baseURL")))
	if err != nil {
		t.Fatalf("could open http client: %v ", err)
	}
	endpoint := fmt.Sprintf("/stock/%s/quote", "aa")
	lastURL, err := client.addToken(endpoint)
	if err != nil {
		t.Fatalf("Error adding token to the url: %v ", err)
	}
	fmt.Println(lastURL)

}
