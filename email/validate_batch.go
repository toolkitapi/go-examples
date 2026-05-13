//go:build ignore

// Email Toolkit — Validate a batch of email addresses
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run validate_batch.go
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
	result, err := email.ValidateBatch(context.Background(), map[string]interface{}{
		"emails": []string{
			"user@example.com",
			"invalid-email",
			"test@gmail.com",
			"noreply@toolkitapi.io",
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
