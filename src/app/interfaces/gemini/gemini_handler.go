package gemini

import (
	"context"
	"github.com/google/generative-ai-go/genai"
)

type GeminiHandler interface {
	Request(ctx context.Context, input string) (*genai.GenerateContentResponse, error)
}
