package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiHandler struct {
	client *genai.Client
}

func NewGeminiHandler(ctx context.Context) *GeminiHandler {
	credentials := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(credentials))
	if err != nil {
		panic(err)
	}
	return &GeminiHandler{
		client: client,
	}
}

func (g *GeminiHandler) Request(ctx context.Context, input string) (*genai.GenerateContentResponse, error) {
	defer func(client *genai.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(g.client)

	model := g.client.GenerativeModel("gemini-2.0-flash-exp")
	output, err := model.GenerateContent(ctx, genai.Text(input))
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}
	return output, nil
}
