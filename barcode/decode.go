//go:build ignore

// Barcode Toolkit — Decode barcode(s) from an image URL
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run decode.go
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
	result, err := bc.BarcodeDecode(context.Background(), map[string]interface{}{
		"url": "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e9/UPC-A-036000291452.svg/220px-UPC-A-036000291452.svg.png",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
