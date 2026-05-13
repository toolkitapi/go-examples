//go:build ignore

// DevTools Toolkit — Test a regular expression
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run regex_run.go
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
	result, err := dt.RegexTest(context.Background(), map[string]interface{}{
		"pattern": `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
		"text":    "user@example.com",
		"flags":   "i",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
