package router

import (
	"Bot-or-Not/internal/app/handler"
	"Bot-or-Not/pkg/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Root struct {
	Echo *echo.Echo
}

func New(playerHandler handler.IPlayerHandler) *Root {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", config.FEURL},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	e.POST("/game/start", playerHandler.StartNewGame)                           //request:passcode; response:id
	e.POST("/player/topic", playerHandler.SubmitPlayerTopic)                    //request:id, topic
	e.GET("/opponent/topic", playerHandler.FetchOpponentTopic)                  //request:id, passcode; response:topic
	e.POST("/opponent/answer", playerHandler.SubmitAnswerToOpponent)            //request:id, passcode, answer
	e.GET("/answers", playerHandler.FetchAnswersForComparison)                  //request:id
	e.POST("/answer/is-player", playerHandler.CompareAnswerIsPlayer)            //request:id,select_answer
	e.GET("/opponent/answer/is-player", playerHandler.IsOpponentAnswerByPlayer) //request:id, passcode; response:topic,answer,ai_answer,is_player
	e.DELETE("/game/end", playerHandler.EndGame)                                //request:id

	return &Root{
		Echo: e,
	}
}
