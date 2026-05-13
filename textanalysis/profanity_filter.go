//go:build ignore

// Text Analysis Toolkit — Filter profanity from a text
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run profanity_filter.go
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
	result, err := ta.Profanity(context.Background(), map[string]interface{}{
		"text": "This is a clean sentence with no bad words.",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
