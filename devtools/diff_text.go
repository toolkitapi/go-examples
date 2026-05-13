//go:build ignore

// DevTools Toolkit — Diff two text strings
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run diff_text.go
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

	dt := toolkitapi.NewDevtools(apiKey)
	result, err := dt.Diff(context.Background(), map[string]interface{}{
		"a": "The quick brown fox jumps over the lazy dog.",
		"b": "The quick brown fox leaps over the lazy cat.",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
