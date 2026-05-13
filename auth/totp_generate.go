//go:build ignore

// Auth Toolkit — Generate a TOTP secret and QR code
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run totp_generate.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	toolkitapi "github.com/toolkitapi/go-sdk"
)

func main() {
	apiKey := os.Getenv("TOOLKITAPI_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: TOOLKITAPI_KEY is not set")
		os.Exit(1)
	}

	auth := toolkitapi.NewAuth(apiKey)
	result, err := auth.TOTPGenerate(context.Background(), map[string]string{
		"issuer":       "MyApp",
		"account_name": "user@example.com",
		"digits":       "6",
		"period":       "30",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
