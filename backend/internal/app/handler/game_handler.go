package handler

import (
	"net/http"
	"strconv"

	"Bot-or-Not/internal/app/dto"
	"Bot-or-Not/internal/app/service"

	"github.com/labstack/echo/v4"
)

type IGameHandler interface {
	StartGame(c echo.Context) error
	// GetAIAnswer(c echo.Context) error
	VerifyAIAnswer(c echo.Context) error
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

// func (h *gameHandler) GetAIAnswer(c echo.Context) error {
// 	gameIDStr := c.Param("gameID")
// 	gameID, err := strconv.ParseUint(gameIDStr, 10, 32)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	game, err := h.gs.GetGameByID(c.Request().Context(), uint(gameID))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	return c.JSON(http.StatusOK, game)
// }

func (h *gameHandler) VerifyAIAnswer(c echo.Context) error {
	var gameReq dto.Game
	if err := c.Bind(&gameReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	gameIDStr := c.Param("gameID")
	gameID, err := strconv.ParseUint(gameIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	game, err := h.gs.GetGameByID(c.Request().Context(), uint(gameID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if game.AIAnswer == gameReq.AIAnswer {
		return c.JSON(http.StatusOK, "Correct")
	}
	return c.JSON(http.StatusOK, "Incorrect")
}
