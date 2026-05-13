//go:build ignore

// Geo Toolkit — Look up geolocation for an IP address
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run ip_lookup.go
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
	result, err := geo.IPLookup(context.Background(), map[string]string{
		"ip": "8.8.8.8",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
