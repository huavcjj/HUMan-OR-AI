package handler

import (
	"Bot-or-Not/internal/app/dto"
	"Bot-or-Not/internal/app/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IPlayerHandler interface {
	CreatePlayer(c echo.Context) error
	GetPlayersByGameID(c echo.Context) error
}

type playerHandler struct {
	ps service.IPlayerService
}

func NewPlayerHandler(ps service.IPlayerService) IPlayerHandler {
	return &playerHandler{ps: ps}
}

func (h *playerHandler) CreatePlayer(c echo.Context) error {
	var playerReq dto.Player
	if err := c.Bind(&playerReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	gameIDStr := c.Param("gameID")
	gameID, err := strconv.ParseUint(gameIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	playerReq.GameID = uint(gameID)

	newPlayer, err := h.ps.CreatePlayer(c.Request().Context(), &playerReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, newPlayer)
}

func (h *playerHandler) GetPlayersByGameID(c echo.Context) error {
	gameIDStr := c.Param("gameID")
	gameID, err := strconv.ParseUint(gameIDStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	players, err := h.ps.GetPlayersByGameID(c.Request().Context(), uint(gameID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, players)
}
