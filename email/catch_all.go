//go:build ignore

// Email Toolkit — Check if a domain accepts catch-all email
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run catch_all.go
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

	email := toolkitapi.NewEmail(apiKey)
	result, err := email.Catch(context.Background(), map[string]string{
		"domain": "example.com",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
