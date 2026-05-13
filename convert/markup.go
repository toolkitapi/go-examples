//go:build ignore

// Convert Toolkit — Convert between markup formats (Markdown → HTML)
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run markup.go
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
	result, err := conv.ConvertMarkupMarkup(context.Background(), map[string]interface{}{
		"content":     "# Hello\n\nThis is **bold** and _italic_ text.\n\n- Item 1\n- Item 2",
		"from_format": "markdown",
		"to_format":   "html",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
