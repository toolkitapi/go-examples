//go:build ignore

// Text Analysis Toolkit — Detect and mask PII in text
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run data_mask.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("TOOLKITAPI_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: TOOLKITAPI_KEY is not set")
		os.Exit(1)
	}

	body, _ := json.Marshal(map[string]interface{}{
		"text": "John Smith's SSN is 123-45-6789 and his email is john.smith@example.com. " +
			"His credit card number is 4111-1111-1111-1111.",
	})
	req, _ := http.NewRequest("POST",
		"https://textanalysis.toolkitapi.io/v1/text/mask",
		bytes.NewReader(body))
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Fprintf(os.Stderr, "Error: HTTP %d: %s\n", resp.StatusCode, string(respBody))
		os.Exit(1)
	}

	var result interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		fmt.Fprintf(os.Stderr, "Error: decoding response: %v\n", err)
		os.Exit(1)
	}
	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
