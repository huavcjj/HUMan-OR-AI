package handler

import (
	"net/http"

	"Bot-or-Not/internal/app/service"

	"github.com/labstack/echo/v4"
)

type IGameHandler interface {
	StartGame(c echo.Context) error
}

type gameHandler struct {
	gs service.IGameService
}

func NewGameHandler(gs service.IGameService) IGameHandler {
	return &gameHandler{gs: gs}
}

func (h *gameHandler) StartGame(c echo.Context) error {
	newGame, err := h.gs.CreateGame(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, newGame)
}
