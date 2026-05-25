//go:build ignore

// PDF Toolkit — Convert document formats
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run convert_document.go
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

	pdf := toolkitapi.NewPdf(apiKey)
	data, err := pdf.ConvertDocument(context.Background(), map[string]string{
		"url":         "https://raw.githubusercontent.com/mozilla/pdf.js/master/examples/learning/helloworld.pdf",
		"from_format": "pdf",
		"to_format":   "docx",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.Marshal(map[string]interface{}{"status": "ok", "bytes": len(data)})
	fmt.Println(string(out))
}
