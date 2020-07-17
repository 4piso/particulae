package iexcloud

import (
	"context"
	"fmt"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	ctx := context.Background()
	endpoint := "/stocs/appl"
	params := map[string]string{
		"token": "cadfasf",
		"user":  "josue",
	}
	client := NewClient()
	res, err := client.Get(ctx, endpoint, params, nil)
	if err != nil {
		t.Fatalf("could not make the request: %v ", err)
	}
	fmt.Println(res)
}
