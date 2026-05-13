//go:build ignore

// Webhook Toolkit — Hit a mock endpoint
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run hit_mock.go
// Note: Replace "your-mock-id-here" with a real mock ID from create_mock.go
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

	wh := toolkitapi.NewWebhook(apiKey)
	result, err := wh.HitMock(context.Background(), "your-mock-id-here")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
