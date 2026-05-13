//go:build ignore

// Webhook Toolkit — List captured requests for a bin
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run list_requests.go
// Note: Replace "your-bin-id-here" with a real bin ID from create_bin.go
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
	result, err := wh.List(context.Background(), "your-bin-id-here", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
