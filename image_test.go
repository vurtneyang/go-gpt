package gogpt

import (
	"context"
	"encoding/json"
	"os"
	"testing"
)

func TestCreateImage(t *testing.T) {
	apiToken := os.Getenv("OPENAI_KEY")
	if apiToken == "" {
		t.Skip("Skipping testing against production OpenAI API. Set OPENAI_TOKEN environment variable to enable it.")
	}

	var err error
	c := NewClient(apiToken)
	ctx := context.Background()
	res, err := c.CreateImage(ctx, ImageRequest{
		Prompt:         "a cat",
		N:              1,
		Size:           "1024x1024",
		ResponseFormat: "",
		User:           "",
	})
	if err != nil {
		t.Fatalf("ListEngines error: %v", err)
	}
	jsonStr, _ := json.Marshal(res)
	t.Logf("%s", string(jsonStr))
}
