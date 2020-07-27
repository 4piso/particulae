package iexcloud

import (
	"context"
	"fmt"
	"testing"
)

const (
	token = "sk_63b7a5ae9a6a43eebbf78e64b3f61a5a"
)

func TestMakeRequest(t *testing.T) {
	ctx := context.Background()
	endpoint := "/"
	params := map[string]string{
		"token": token,
	}
	client := NewClient()
	res, err := client.Get(ctx, endpoint, params, nil)
	if err != nil {
		t.Fatalf("Error %v ", err)
	}
	fmt.Println(res)

}
