//go:build ignore

// PDF Toolkit — Split pages out of a PDF
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run split.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	toolkitapi "github.com/toolkitapi/go-sdk"
)

const pdfURL = "https://raw.githubusercontent.com/mozilla/pdf.js/master/examples/learning/helloworld.pdf"

func main() {
	apiKey := os.Getenv("TOOLKITAPI_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: TOOLKITAPI_KEY is not set")
		os.Exit(1)
	}

	pdf := toolkitapi.NewPdf(apiKey)
	result, err := pdf.Split(context.Background(), map[string]string{
		"url":   pdfURL,
		"pages": "1",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
