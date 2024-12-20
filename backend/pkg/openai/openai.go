package openai

import (
	"context"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	client *openai.Client
}

func NewOpenAI(apiKey string) *OpenAI {
	client := openai.NewClient(apiKey)
	return &OpenAI{client: client}
}

// 12/20追加作成。GenerateGameTopicAIAnswerからほぼ流用。テスト済み
func (o *OpenAI) GenerateAIAnswer(ctx context.Context, topic string) (string, error) {

	//プレイヤーのお題内容を受け取り、お題に対してGPTで生成した回答を返す。
	answerPrompt := fmt.Sprintf("お題: %s\nこのお題に対して、芸人レベルに面白い、現実でもあり得る回答を一つ30文字以内で考えてください。", topic)

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
		return "", fmt.Errorf("failed to get answer: %w", err)
	}
	if len(answerResp.Choices) == 0 {
		return "", fmt.Errorf("no answer returned from GPT")
	}
	answer := answerResp.Choices[0].Message.Content

	//不自然な始まりを除去
	answer = RemovePrefixes_answer(answer)

	//charsToRemoveに含まれる文字を回答(answer)から取り除く（例 「宇宙で一番「困ること」は何？」　→　宇宙で一番困ることは何？）
	answer = RemoveChars_answer(answer)

	//:よりの文字を削除（例 ニャンニャンカフェ：ねこがバリスタ、ワンちゃんがウェイター　→　ねこがバリスタ、ワンちゃんがウェイター）
	//:があったらもう一度答えを生成するのでもいいかも
	answer = TrimBeforeColon_answer(answer)

	return answer, nil
}

func (o *OpenAI) FormatAIAnswer(aiAnswer string) string {
	//AIの回答を受け取り、整形して返す。

	//回答の最初と最後にある「」を取り除く
	aiAnswer = strings.TrimPrefix(aiAnswer, "「")
	aiAnswer = strings.TrimSuffix(aiAnswer, "」")

	// prefixesに含まれる言葉で始まる場合、それを除去する
	prefixes := []string{"なぜならば、", "なぜならば", "なぜなら、", "なぜなら"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(aiAnswer, prefix) {
			aiAnswer = strings.TrimPrefix(aiAnswer, prefix)
			break
		}
	}

	// charsToRemove_answerに含まれる文字を除去
	charsToRemove_answer := []string{"「", "」", "！", "。", "\"", "“", "”"}
	for _, char := range charsToRemove_answer {
		aiAnswer = strings.ReplaceAll(aiAnswer, char, "")
	}

	// :がある場合、:までの文字を除去
	if idx := strings.LastIndex(aiAnswer, ":"); idx != -1 {
		aiAnswer = aiAnswer[idx+1:]
	}

	return aiAnswer
}

// 大喜利のお題はGPTではなく、プレイヤーが考えることになったため、この関数は不要になった
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

	//回答整形
	answer = o.FormatAIAnswer(answer)

	return topic, answer, nil
}
