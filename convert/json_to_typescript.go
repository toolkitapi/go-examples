//go:build ignore

// Convert Toolkit — Generate TypeScript interfaces from JSON
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run json_to_typescript.go
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
	result, err := conv.JSON(context.Background(), map[string]interface{}{
		"data": map[string]interface{}{
			"user": map[string]interface{}{
				"id":    1,
				"name":  "Alice",
				"email": "alice@example.com",
				"roles": []string{"admin", "user"},
			},
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
