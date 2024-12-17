package openai

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	client *openai.Client
}

func NewOpenAI(apiKey string) *OpenAI {
	client := openai.NewClient(apiKey)
	return &OpenAI{client: client}
}

func (o *OpenAI) GenerateGameTopicAIAnswer(ctx context.Context) (string, string, error) {
	//未実装

	return "abc", "def", nil
}
