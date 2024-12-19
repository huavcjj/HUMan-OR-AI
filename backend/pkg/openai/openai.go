package openai

import (
	"context"
	"fmt"
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
	// ①大喜利のお題を生成
	topicPrompt := "あなたは創造的で面白い大喜利のお題を考えるAIです。人をクスッと笑わせるような新しい大喜利のお題を一つ考えてください。ただし、お題の内容だけ、返してください。"

	topicResp, err := o.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo, // 必要に応じてモデル選択(gpt-4でも可)
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "あなたは優れた大喜利のお題を考え出すことができます。お題の内容だけを返します",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: topicPrompt,
			},
		},
		Temperature: 0.8,
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to get topic: %w", err)
	}
	if len(topicResp.Choices) == 0 {
		return "", "", fmt.Errorf("no topic returned from GPT")
	}
	topic := topicResp.Choices[0].Message.Content

	// ②生成されたお題に対する面白い回答を生成
	answerPrompt := fmt.Sprintf("お題: %s\nこのお題に対して、面白い回答を一つ考えてください。", topic)

	answerResp, err := o.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "あなたはユーモアに富んだ大喜利回答者です。",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: answerPrompt,
			},
		},
		Temperature: 0.7,
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to get answer: %w", err)
	}
	if len(answerResp.Choices) == 0 {
		return "", "", fmt.Errorf("no answer returned from GPT")
	}
	answer := answerResp.Choices[0].Message.Content

	return topic, answer, nil
}
