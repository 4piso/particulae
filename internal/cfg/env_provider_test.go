package cfg_test

import (
	"fmt"
	"testing"

	"github.com/joshvoll/particulae/internal/cfg"
)

func TestEnvVariables(t *testing.T) {
	envConfig := &cfg.EnvProvider{
		Namespace: "PROC",
	}
	config, err := cfg.New(envConfig)
	if err != nil {
		t.Fatalf("ERROR: %v ", err)
	}
	fmt.Println(config.String("TOKEN"))
}
