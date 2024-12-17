package router

import (
	"Bot-or-Not/internal/app/handler"

	"github.com/labstack/echo/v4"
)

type Root struct {
	Echo *echo.Echo
}

func New(gameHandler handler.IGameHandler, playerHandler handler.IPlayerHandler, answerHandler handler.IAnswerHandler) *Root {
	e := echo.New()

	e.POST("/game/:id", gameHandler.StartGame)

	game := e.Group("/game/:gameID")
	game.POST("/player", playerHandler.CreatePlayer)
	game.GET("/player", playerHandler.GetPlayersByGameID)

	e.POST("/game/:gameID/:playerID/answer", answerHandler.CreateAnswer)

	return &Root{
		Echo: e,
	}
}
