//go:build ignore

// Barcode Toolkit — Decode QR code(s) from an image URL
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run qr_decode.go
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
	result, err := bc.QrDecode(context.Background(), map[string]interface{}{
		"url": "https://upload.wikimedia.org/wikipedia/commons/thumb/d/d0/QR_code_for_mobile_English_Wikipedia.svg/320px-QR_code_for_mobile_English_Wikipedia.svg.png",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
