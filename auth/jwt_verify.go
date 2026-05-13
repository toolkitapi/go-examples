//go:build ignore

// Auth Toolkit — Verify and decode a JWT token
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run jwt_verify.go
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

	const secret = "my-signing-secret"
	auth := toolkitapi.NewAuth(apiKey)

	// Generate a token, then immediately verify it
	genResp, err := auth.JWTGenerate(context.Background(), map[string]interface{}{
		"payload":    map[string]interface{}{"sub": "user-123", "role": "admin"},
		"secret":     secret,
		"algorithm":  "HS256",
		"expires_in": 3600,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating token: %v\n", err)
		os.Exit(1)
	}
	token := genResp.(map[string]interface{})["token"].(string)
	fmt.Println("Generated token:", token[:40]+"...")

	result, err := auth.JWTVerify(context.Background(), map[string]interface{}{
		"token":  token,
		"secret": secret,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
