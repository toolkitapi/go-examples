//go:build ignore

// Geo Toolkit — Convert a timestamp between timezones
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run timezone_convert.go
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

	geo := toolkitapi.NewGeo(apiKey)
	result, err := geo.TimezoneConvert(context.Background(), map[string]string{
		"timestamp": "2026-06-15T14:30:00",
		"from":      "America/New_York",
		"to":        "Asia/Tokyo",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
