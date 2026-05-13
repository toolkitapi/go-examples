//go:build ignore

// Scrape Toolkit — Scrape a web page
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run scrape.go
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

	sc := toolkitapi.NewScrape(apiKey)
	result, err := sc.Scrape(context.Background(), map[string]interface{}{
		"url":  "https://toolkitapi.io",
		"type": "html",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
