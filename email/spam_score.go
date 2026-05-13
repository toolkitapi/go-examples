//go:build ignore

// Email Toolkit — Score an email for spam likelihood
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run spam_score.go
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

	email := toolkitapi.NewEmail(apiKey)
	result, err := email.Spam(context.Background(), map[string]interface{}{
		"subject": "URGENT: You've WON a FREE prize! Click HERE NOW!!!",
		"body":    "Congratulations! You have been selected as our lucky winner. Click the link below to claim your FREE reward. Limited time offer! Act NOW before it expires!",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
