//go:build ignore

// Barcode Toolkit — Bulk generate QR codes
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run qr_bulk.go
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

	bc := toolkitapi.NewBarcode(apiKey)
	result, err := bc.QrBulk(context.Background(), map[string]interface{}{
		"items": []map[string]interface{}{
			{"data": "https://toolkitapi.io", "label": "Homepage"},
			{"data": "https://toolkitapi.io/docs", "label": "Docs"},
		},
		"size":   200,
		"format": "png",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
