//go:build ignore

// Image Toolkit — Remove the background from an image
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run remove_background.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	toolkitapi "github.com/toolkitapi/go-sdk"
)

const imageURL = "https://picsum.photos/seed/toolkitapi/400/300"

func main() {
	apiKey := os.Getenv("TOOLKITAPI_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: TOOLKITAPI_KEY is not set")
		os.Exit(1)
	}

	img := toolkitapi.NewImage(apiKey)
	result, err := img.ImageRemoveBackground(context.Background(), map[string]string{
		"url": imageURL,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
