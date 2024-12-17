package router

import (
	"Bot-or-Not/internal/app/handler"

	"github.com/labstack/echo/v4"
)

type Root struct {
	Echo *echo.Echo
}

func New(gameHandler handler.IGameHandler) *Root {
	e := echo.New()

	e.POST("/game", gameHandler.StartGame)

	return &Root{
		Echo: e,
	}
}
