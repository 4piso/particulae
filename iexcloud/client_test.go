package iexcloud

import (
	"context"
	"testing"

	"github.com/joshvoll/particulae/internal/cfg"
)

var (
	token  = ""
	client = NewClient(getToken())
)

func getToken() string {
	envConfig := &cfg.EnvProvider{
		Namespace: "PROC",
	}
	config, err := cfg.New(envConfig)
	if err != nil {
		panic(err)
	}
	tok, _ := config.String("TOKEN")
	return tok
}

func TestMakeRequest(t *testing.T) {
	ctx := context.Background()
	token := getToken()
	endpoint := "/"
	params := map[string]string{
		"token": token,
	}
	client := NewClient(getToken())
	_, err := client.Get(ctx, endpoint, params, nil)
	if err != nil {
		t.Fatalf("Error %v ", err)
	}
}

func TestQuoteRequest(t *testing.T) {
	ctx := context.Background()
	quote, err := client.Quote(ctx, "aapl", true)
	if err != nil {
		t.Fatalf("Error: %v ", err)
	}
	if quote.Symbol != "AAPL" {
		t.Error("should return AAPL")
	}
}

func TestEarningsRequest(t *testing.T) {
	ctx := context.Background()
	earnings, err := client.Earnings(ctx, "aapl")
	if err != nil {
		t.Fatalf("error earnings: %v ", err)
	}
	if earnings.Symbol != "AAPL" {
		t.Errorf("should be AAPL, got=%v ", earnings.Symbol)
	}
}
