//go:build ignore

// Auth Toolkit — Verify a password against a stored hash
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run verify_password.go
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

	const password = "my-secret-password"
	auth := toolkitapi.NewAuth(apiKey)

	// Hash first, then verify
	hashResp, err := auth.HashPassword(context.Background(), map[string]interface{}{
		"password": password,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error hashing password: %v\n", err)
		os.Exit(1)
	}
	hash := hashResp.(map[string]interface{})["hash"].(string)
	fmt.Println("Password hash:", hash)

	result, err := auth.Verify(context.Background(), map[string]interface{}{
		"password": password,
		"hash":     hash,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
