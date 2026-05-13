//go:build ignore

// DevTools Toolkit — Evaluate a math expression
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run math_eval.go
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

	dt := toolkitapi.NewDevtools(apiKey)
	result, err := dt.Math(context.Background(), map[string]interface{}{
		"expression": "sqrt(144) + (2^10 / 4)",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
