//go:build ignore

// Auth Toolkit — Generate a JWT token
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run jwt_generate.go
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

	auth := toolkitapi.NewAuth(apiKey)
	result, err := auth.JWTGenerate(context.Background(), map[string]interface{}{
		"payload": map[string]interface{}{
			"sub":   "user-123",
			"role":  "admin",
			"email": "user@example.com",
		},
		"secret":     "my-signing-secret",
		"algorithm":  "HS256",
		"expires_in": 3600,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
