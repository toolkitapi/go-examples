//go:build ignore

// Auth Toolkit — Hash a password (bcrypt / argon2 / scrypt)
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run hash_password.go
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
	result, err := auth.HashPassword(context.Background(), map[string]interface{}{
		"password":  "my-secret-password",
		"algorithm": "bcrypt",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
