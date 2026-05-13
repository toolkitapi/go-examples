//go:build ignore

// Scrape Toolkit — Check mobile-friendliness of a URL
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run seo_mobile_friendly.go
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
	result, err := sc.SeoMobileFriendly(context.Background(), map[string]string{
		"url": "https://toolkitapi.io",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
