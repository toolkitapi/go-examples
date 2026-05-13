//go:build ignore

// Text Analysis Toolkit — Summarize a block of text
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run summarize.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	toolkitapi "github.com/toolkitapi/go-sdk"
)

const longText = `Artificial intelligence (AI) is intelligence demonstrated by machines, ` +
	`as opposed to the natural intelligence displayed by animals including humans. ` +
	`AI research has been defined as the field of study of intelligent agents, which ` +
	`refers to any system that perceives its environment and takes actions that maximize ` +
	`its chance of achieving its goals. The term artificial intelligence was coined by ` +
	`John McCarthy in 1956. Since then, AI has evolved from rule-based systems to modern ` +
	`machine learning and deep learning approaches that power applications like image ` +
	`recognition, natural language processing, and autonomous vehicles.`

func main() {
	apiKey := os.Getenv("TOOLKITAPI_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: TOOLKITAPI_KEY is not set")
		os.Exit(1)
	}

	ta := toolkitapi.NewTextanalysis(apiKey)
	result, err := ta.Summarize(context.Background(), map[string]interface{}{
		"text":      longText,
		"sentences": 2,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
