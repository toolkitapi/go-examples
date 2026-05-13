//go:build ignore

// Geo Toolkit — Calculate the distance between two or more geographic points
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run distance.go
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
	result, err := geo.Distance(context.Background(), map[string]interface{}{
		"points": []map[string]interface{}{
			{"lat": 51.5074, "lon": -0.1278}, // London
			{"lat": 48.8566, "lon": 2.3522},  // Paris
		},
		"unit": "km",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
