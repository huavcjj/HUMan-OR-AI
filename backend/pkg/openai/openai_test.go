package openai

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGenerateGameTopicAIAnswer(t *testing.T) {
	// 相対パスで.envをロード
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Fatal("OPENAI_API_KEY not set")
	}

	oai := NewOpenAI(apiKey)
	odai, answer, err := oai.GenerateGameTopicAIAnswer(context.Background())
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	if odai == "" {
		t.Error("odai is empty")
	}
	if answer == "" {
		t.Error("answer is empty")
	}
	t.Logf("Odai: %s\nAnswer: %s\n", odai, answer)
}

//go test -v -run TestGenerateGameTopicAIAnswer -count=1

func TestFormatAIAnswer(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Fatal("OPENAI_API_KEY not set")
	}

	oai := NewOpenAI(apiKey)

	for i := 0; i < 10; i++ { // 10回繰り返す
		t.Run(fmt.Sprintf("Test-%d", i+1), func(t *testing.T) {
			// お題と回答を生成
			odai, answer, err := oai.GenerateGameTopicAIAnswer(context.Background())
			if err != nil {
				t.Fatalf("failed to generate: %v", err)
			}

			// AIの回答を整形
			formattedAnswer := oai.FormatAIAnswer(answer)

			// ログに結果を出力
			t.Logf("Test #%d - お題: %s\n", i+1, odai)
			t.Logf("Test #%d - 整形前: %s\n", i+1, answer)
			t.Logf("Test #%d - 整形後: %s\n", i+1, formattedAnswer)

			// 結果が空でないことを確認
			if formattedAnswer == "" {
				t.Errorf("Formatted answer is empty on test #%d", i+1)
			}
		})
	}
}

//go test -v -run TestFormatAIAnswer -count=1
