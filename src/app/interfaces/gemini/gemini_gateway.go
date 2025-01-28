package gemini

import (
	"context"
	"github.com/google/generative-ai-go/genai"
)

type GeminiGateway struct {
	GeminiHandler GeminiHandler
}

func (gateway *GeminiGateway) Request(ctx context.Context, input string) (*genai.GenerateContentResponse, error) {
	return gateway.GeminiHandler.Request(ctx, input)
}
