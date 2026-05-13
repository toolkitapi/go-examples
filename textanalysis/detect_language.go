//go:build ignore

// Text Analysis Toolkit — Detect the language of a text
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run detect_language.go
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

	ta := toolkitapi.NewTextanalysis(apiKey)
	result, err := ta.Detect(context.Background(), map[string]interface{}{
		"text": "Bonjour le monde, comment allez-vous aujourd'hui?",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
