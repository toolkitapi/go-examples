//go:build ignore

// Barcode Toolkit — Generate a QR code
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run qr_generate.go
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
	result, err := bc.QrGenerate(context.Background(), map[string]interface{}{
		"data":        "https://toolkitapi.io",
		"size":        300,
		"format":      "png",
		"error_level": "M",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
