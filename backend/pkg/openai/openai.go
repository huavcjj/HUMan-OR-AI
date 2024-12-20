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

func (o *OpenAI) GenerateAIAnswer(ctx context.Context, topic string) (string, error) {

	//プレイヤーのお題内容を受け取り、お題に対してGPTで生成した回答を返す。
	//未実装

	return "GPTで生成した回答", nil
}

func (o *OpenAI) FormatAIAnswer(aiAnswer string) string {
	//AIの回答を受け取り、整形して返す。
	//未実装

	return "整形した回答"
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

	topic = RemoveChars_topic(topic)

	// ②生成されたお題に対する面白い回答を生成
	answerPrompt := fmt.Sprintf("お題: %s\nこのお題に対して、面白い回答を一つ30文字以内で考えてください。", topic)

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

	//不自然な始まりを除去
	answer = RemovePrefixes_answer(answer)

	//charsToRemoveに含まれる文字を回答(answer)から取り除く（例 「宇宙で一番「困ること」は何？」　→　宇宙で一番困ることは何？）
	answer = RemoveChars_answer(answer)

	//:よりの文字を削除（例 ニャンニャンカフェ：ねこがバリスタ、ワンちゃんがウェイター　→　ねこがバリスタ、ワンちゃんがウェイター）
	//:があったらもう一度答えを生成するのでもいいかも
	answer = TrimBeforeColon_answer(answer)

	return topic, answer, nil
}
