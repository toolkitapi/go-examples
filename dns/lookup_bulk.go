//go:build ignore

// DNS Toolkit — Bulk DNS lookup for multiple domains
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run lookup_bulk.go
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

	dns := toolkitapi.NewDNS(apiKey)
	result, err := dns.LookupBulk(context.Background(), map[string]string{
		"type": "MX",
	}, []string{"github.com", "python.org", "nodejs.org"})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
