//go:build ignore

// Convert Toolkit — Convert between document formats (via URL)
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run document.go
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

	conv := toolkitapi.NewConvert(apiKey)
	result, err := conv.ConvertDocumentDocument(context.Background(), map[string]interface{}{
		"url":         "https://raw.githubusercontent.com/mozilla/pdf.js/master/examples/learning/helloworld.pdf",
		"from_format": "pdf",
		"to_format":   "docx",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
