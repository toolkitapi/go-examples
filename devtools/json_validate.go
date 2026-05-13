//go:build ignore

// DevTools Toolkit — Validate JSON
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run json_validate.go
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
	result, err := dt.JSONValidate(context.Background(), map[string]interface{}{
		"data": `{"name": "Alice", "age": 30, "active": true}`,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
