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

//GPTに大喜利のお題を送り、回答を生成してもらう
//引数：ゲームID,お題num,お題txt,
//戻り値：生成した回答
func (o *OpenAI) GetGPTResponse(ctx context.Context, gameID string,themaNum int,themaTxt string)(string,error){
	//プロンプト作成
	prompt :=fmt.Sprintf(
		"以下は大喜利ゲームのお題です。\nお題: %s\nこのお題に対して人間に負けないくらい面白い回答を一つ考えてください。",
		gameID, themaNum,themaTxt,
	)

	//ChatCompletion API呼び出し
	resp, err:= o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,//ここにモデル名
			Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: "あなたは創造的でユーモアのある大喜利回答者です。短く面白い回答のみを返してください。",
		
			},
			{
				Role: openai.ChatMessageRoleUser,
				Content: prompt,
		
			},
		},
		Temperature: 0.7, // ユーモアや創造性を出したい場合は多少上げる
		},
	)
	if err!= nil{
		return "","", fmt.Errorf("ChatCompletion API request failed: %w", err)
	}

	if len(resp.Choices)==0{
		return "","",fmt.Errorf("no response from GPT")
	}

	answer :=resp.Choices[0].Message.Content
	return answer, nil

	
}