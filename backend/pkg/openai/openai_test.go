package openai

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestGenerateAIAnswer(t *testing.T) {
	// 相対パスで.envをロード
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Fatal("OPENAI_API_KEY not set")
	}
	//ここでお題を変えられる
	topic := "学校が急に休みになった！なぜ？？"
	oai := NewOpenAI(apiKey)
	answer, err := oai.GenerateAIAnswer(context.Background(), topic)
	if err != nil {
		t.Fatalf("failed to generate: %v", err)
	}

	if topic == "" {
		t.Error("odai is empty")
	}
	if answer == "" {
		t.Error("answer is empty")
	}
	t.Logf("Odai: %s\nAnswer: %s\n", topic, answer)
}

//go test -v -run TestGenerateAIAnswer -count=1

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
