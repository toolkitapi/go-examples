//go:build ignore

// Convert Toolkit — Convert between data formats (JSON → CSV)
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run data.go
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
	result, err := conv.ConvertDataData(context.Background(), map[string]interface{}{
		"data":        []map[string]interface{}{{"name": "Alice", "age": 30}, {"name": "Bob", "age": 25}},
		"from_format": "json",
		"to_format":   "csv",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
