//go:build ignore

// Media Toolkit — Get information about a YouTube channel
// Usage: export TOOLKITAPI_KEY=tk_live_...; go run youtube_channel.go
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

	media := toolkitapi.NewMedia(apiKey)
	result, err := media.YoutubeChannel(context.Background(), map[string]string{
		"id": "UCVHFbw7woebKtFFbuylTdig",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	out, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(out))
}
