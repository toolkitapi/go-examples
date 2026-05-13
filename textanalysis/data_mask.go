//go:build ignore

// Text Analysis Toolkit — Detect and mask PII in text
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run data_mask.go
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

	ta := toolkitapi.NewTextanalysis(apiKey)
	result, err := ta.Data(context.Background(), map[string]interface{}{
		"text": "John Smith's SSN is 123-45-6789 and his email is john.smith@example.com. " +
			"His credit card number is 4111-1111-1111-1111.",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
