package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiHandler struct{}

func NewGeminiHandler() *GeminiHandler {
	ctx := context.Background()
	credentials := os.Getenv("GEMINI_API_KEY")

	client, err := genai.NewClient(ctx, option.WithAPIKey(credentials))
	if err != nil {
		panic(err)
	}
	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(client)

	return &GeminiHandler{}
}

func (g *GeminiHandler) Request(ctx context.Context, input string) (*genai.GenerateContentResponse, error) {
	client, err := newClient(ctx)
	if err != nil {
		return nil, err
	}
	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(client)

	model := client.GenerativeModel("gemini-2.0-flash-exp")
	output, err := model.GenerateContent(ctx, genai.Text(input))
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}
	return output, nil
}

func newClient(ctx context.Context) (*genai.Client, error) {
	credentials := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(credentials))
	if err != nil {
		return nil, fmt.Errorf("failed to create gemini client: %w", err)
	}
	return client, nil
}
