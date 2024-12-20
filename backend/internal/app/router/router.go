package router

import (
	"Bot-or-Not/internal/app/handler"

	"github.com/labstack/echo/v4"
)

type Root struct {
	Echo *echo.Echo
}

func New(playerHandler handler.IPlayerHandler) *Root {
	e := echo.New()

	e.POST("/game/start", playerHandler.StartNewGame)                //request:passcode; response:id
	e.POST("/player/topic", playerHandler.SubmitPlayerTopic)         //request:id, topic
	e.GET("/opponent/topic", playerHandler.FetchOpponentTopic)       //request:id passcode; response:topic
	e.POST("/opponent/answer", playerHandler.SubmitAnswerToOpponent) //request:id, passcode, answer
	e.GET("/answers", playerHandler.FetchAnswersForComparison)       //request:id
	e.POST("/answer/is-player", playerHandler.CompareAnswerIsPlayer) //request:id,select_answer
	e.POST("/game/end", playerHandler.EndGame)                       //request:id

	return &Root{
		Echo: e,
	}
}
