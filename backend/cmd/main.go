package main

import (
	"Bot-or-Not/internal/di"
	"log"
	"net/http"

	//以下を追加(12/18大道)
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	_ "Bot-or-Not/backend/pkg/openai"
	//ここまで

	_ "Bot-or-Not/internal/migration"
)

func main() {
	root := di.New()

	if err := http.ListenAndServe(":8080", root.Echo); err != nil {
		log.Fatal("Error starting server: ", err)
	}

	//以下を追加(12/18大道)
	err := godotenv.Load()
	if err != nil{
		log.Fatalf("Error loading .env file: %v",err)
	}

	apiKey:=os.Getenv("OPENAI_API_KEY")
	if apiKey ==""{
		log.Fatal("OPENAI_API_KEY not set in environment variables")
	}

	//OpenAIインスタンス作成
	oai := openai.NewOpenAI(apiKey)

	// コンテキストを用意
	ctx := context.Background()

	//実際に呼び出してみる（仮のパラメータ）
	gameID:="game_1"
	themaNum:=1
	themaTxt:="「地元で驚きのイベント」が！どんなイベント？"

	answer,err:=oai.GetGPTResponse(ctx, gameID,themaNum,themaTxt)

	if err != nil{
		log.Fatalf("Error from OpenAI: %v",err)
	}

	//fmt.Println("Prompt:",prompt)
	fmt.Println("Answer:",answer)
}
