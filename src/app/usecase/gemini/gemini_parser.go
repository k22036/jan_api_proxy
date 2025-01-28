package gemini

import (
	"fmt"
	"github.com/google/generative-ai-go/genai"
)

func ParseInput(input []string) string {
	template := `プロンプト:
		以下の商品リストに共通して含まれる商品名を抽出してください。
		完全に一致する商品名のみを出力してください。
		//出力する商品名は、以下の商品リスト全てに含まれる商品名である必要があります。
		//また、出力する商品名は、以下の商品リストに含まれる商品名の中で最も短い商品名である必要があります。
		不必要な文字列を出力しないように注意してください。

		商品リスト:
	`
	data := ""
	end := min(20, len(input))
	for _, in := range input[:end] {
		data += fmt.Sprintf("- %s\n", in)
	}
	return template + data + "\n出力:\n"
}

func ParseResponse(resp *genai.GenerateContentResponse) string {
	result := ""
	for _, cand := range resp.Candidates {
		if cand.Content == nil {
			continue
		}
		for _, part := range cand.Content.Parts {
			result += fmt.Sprintf("%s", part)
		}
	}
	for i := 0; i < len(result); i++ {
		if result[i] == '\n' {
			result = result[:i]
			break
		}
	}

	return result
}
