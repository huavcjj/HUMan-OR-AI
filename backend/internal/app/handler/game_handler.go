package handler

import (
	"net/http"
	"strconv"

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

	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	existingGame, _ := h.gs.GetGameByID(c.Request().Context(), uint(id))

	if existingGame != nil {
		return c.JSON(http.StatusOK, existingGame)
	}

	newGame, err := h.gs.CreateGame(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newGame)
}
