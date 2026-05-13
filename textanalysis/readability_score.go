//go:build ignore

// Text Analysis Toolkit — Calculate the readability score of a text
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run readability_score.go
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
	result, err := ta.Readability(context.Background(), map[string]interface{}{
		"text": "The quick brown fox jumps over the lazy dog. " +
			"Pack my box with five dozen liquor jugs. " +
			"How vexingly quick daft zebras jump!",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
