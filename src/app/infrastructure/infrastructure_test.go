package infrastructure_test

import (
	"app/infrastructure"
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	t.Run("gemini_api_key", func(t *testing.T) {
		geminiApiKey := os.Getenv("GEMINI_API_KEY")
		if geminiApiKey == "" {
			t.Error("GEMINI_API_KEY is not set")
		}
	})
	t.Run("yahoo_app_id", func(t *testing.T) {
		yahooAppId := os.Getenv("YAHOO_APP_ID")
		if yahooAppId == "" {
			t.Error("YAHOO_APP_ID is not set")
		}
	})
}

func TestGetProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		jan := "4902105022924"
		h := infrastructure.NewYahooShoppingHandler()
		res, err := h.GetProduct(jan)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(res)
	})
}

func TestRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		input := "I'm a software engineer."
		ctx := context.Background()
		h := infrastructure.NewGeminiHandler(ctx)
		res, err := h.Request(ctx, input)
		if err != nil {
			t.Error(err)
		}
		printResponse(res)
	})
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content == nil {
			continue
		}
		for _, part := range cand.Content.Parts {
			fmt.Println(part)
		}
	}
}
