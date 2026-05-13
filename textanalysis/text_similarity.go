//go:build ignore

// Text Analysis Toolkit — Compute similarity between two texts
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run text_similarity.go
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
	result, err := ta.TextSimilarity(context.Background(), map[string]interface{}{
		"a": "The cat sat on the mat",
		"b": "The cat is sitting on the mat",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
