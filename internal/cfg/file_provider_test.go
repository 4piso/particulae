package cfg_test

import (
	"fmt"
	"testing"

	"github.com/joshvoll/particulae/internal/cfg"
)

func TestOpenFile(t *testing.T) {
	file := &cfg.FileProvider{
		Filename: "config.json",
	}
	config, err := cfg.New(file)
	if err != nil {
		t.Errorf("Error opening file: %v ", err)
	}
	db, _ := config.String("database")
	server, errorURL := config.URL("server")
	if errorURL != nil {
		t.Errorf("Error from url: %v ", errorURL)
	}
	token, _ := config.String("token")
	fmt.Printf("db, server, token :  %v, %v, %v ", db, server, token)
}
