//go:build ignore

// Convert Toolkit — Convert between spreadsheet formats (CSV → XLSX)
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run spreadsheet.go
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
	result, err := conv.ConvertSpreadsheetSpreadsheet(context.Background(), map[string]interface{}{
		"data":        "name,age,city\nAlice,30,London\nBob,25,Paris",
		"from_format": "csv",
		"to_format":   "xlsx",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
