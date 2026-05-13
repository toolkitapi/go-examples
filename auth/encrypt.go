//go:build ignore

// Auth Toolkit — Encrypt plaintext data (AES-256-GCM)
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run encrypt.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	toolkitapi "github.com/toolkitapi/go-sdk"
)

func main() {
	apiKey := os.Getenv("TOOLKITAPI_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: TOOLKITAPI_KEY is not set")
		os.Exit(1)
	}

	auth := toolkitapi.NewAuth(apiKey)
	// 32-byte AES-256 key as a 64-character hex string
	hexKey := strings.Repeat("0", 64)
	result, err := auth.Encrypt(context.Background(), map[string]interface{}{
		"plaintext": "Sensitive data: SSN 123-45-6789",
		"key":       hexKey,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
