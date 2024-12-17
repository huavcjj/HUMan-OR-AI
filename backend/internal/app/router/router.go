package router

import (
	"Bot-or-Not/internal/app/handler"

	"github.com/labstack/echo/v4"
)

type Root struct {
	Echo *echo.Echo
}

func New(gameHandler handler.IGameHandler, playerHandler handler.IPlayerHandler) *Root {
	e := echo.New()

	e.POST("/game/:id", gameHandler.StartGame)

	e.POST("/player/:gameID", playerHandler.CreatePlayer)
	e.GET("/player/:gameID", playerHandler.GetPlayersByGameID)

	return &Root{
		Echo: e,
	}
}
