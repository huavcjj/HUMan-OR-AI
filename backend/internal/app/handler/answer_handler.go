package handler

import (
	"Bot-or-Not/internal/app/dto"
	"Bot-or-Not/internal/app/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IAnswerHandler interface {
	CreateAnswer(c echo.Context) error
	GetAnswersByGameID(c echo.Context) error
}

type answerHandler struct {
	as service.IAnswerService
}

func NewAnswerHandler(as service.IAnswerService) IAnswerHandler {
	return &answerHandler{as: as}
}

func (h *answerHandler) CreateAnswer(c echo.Context) error {
	var answerReq dto.Answer
	if err := c.Bind(&answerReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	gameIDStr := c.Param("gameID")
	gameID, err := strconv.ParseUint(gameIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	answerReq.GameID = uint(gameID)

	playerIDStr := c.Param("playerID")
	playerID, err := strconv.ParseUint(playerIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	answerReq.PlayerID = uint(playerID)

	newAnswer, err := h.as.CreateAnswer(c.Request().Context(), &answerReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, newAnswer)
}

func (h *answerHandler) GetAnswersByGameID(c echo.Context) error {
	gameIDStr := c.Param("gameID")
	gameID, err := strconv.ParseUint(gameIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	answers, err := h.as.GetAnswersByGameID(c.Request().Context(), uint(gameID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, answers)
}
