package handler

import (
	"Bot-or-Not/internal/app/dto"
	"Bot-or-Not/internal/app/service"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type IPlayerHandler interface {
	StartNewGame(c echo.Context) error              // 新しいゲームを開始する
	SubmitPlayerTopic(c echo.Context) error         // プレイヤーがお題を提出する
	FetchOpponentTopic(c echo.Context) error        // 相手のお題を取得する
	SubmitAnswerToOpponent(c echo.Context) error    // 相手のお題に対する回答を提出する
	FetchAnswersForComparison(c echo.Context) error // AIの回答と相手の回答を取得する
	CompareAnswerIsPlayer(c echo.Context) error     // 送信した回答がAIか人間かを判定する
	EndGame(c echo.Context) error
}

type playerHandler struct {
	ps service.IPlayerService
}

func NewPlayerHandler(ps service.IPlayerService) IPlayerHandler {
	return &playerHandler{ps: ps}
}

func (ph *playerHandler) StartNewGame(c echo.Context) error {
	var player dto.Player
	if err := c.Bind(&player); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if player.Passcode == "" {
		return c.JSON(http.StatusBadRequest, "合言葉を入力してください")
	}

	newPlayer, err := ph.ps.CreatePlayer(c.Request().Context(), &player)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	time.Sleep(5 * time.Second)

	opponent, err := ph.ps.FindAvailableOpponentByPasscode(c.Request().Context(), newPlayer.ID, newPlayer.Passcode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if opponent == nil {

		if err := ph.ps.DeletePlayerByID(c.Request().Context(), newPlayer.ID); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusNotFound, "対戦相手が見つかりませんでした")
	}

	return c.JSON(http.StatusOK, newPlayer.ID)
}

func (ph *playerHandler) SubmitPlayerTopic(c echo.Context) error {
	var player dto.Player
	if err := c.Bind(&player); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := player.ID
	topic := player.Topic

	_, err := ph.ps.UpdateTopicAndAIAnswer(c.Request().Context(), id, topic)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "お題を提出しました")
}

func (ph *playerHandler) FetchOpponentTopic(c echo.Context) error {
	idStr := c.QueryParam("id")
	passcode := c.QueryParam("passcode")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	decodedPasscode, err := url.QueryUnescape(passcode)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	opponentPlayer, err := ph.ps.FindAvailableOpponentByPasscode(c.Request().Context(), uint(id), decodedPasscode)
	if opponentPlayer == nil || err != nil {
		return c.JSON(http.StatusOK, err)
	}
	return c.JSON(http.StatusOK, opponentPlayer.Topic)
}

func (ph *playerHandler) SubmitAnswerToOpponent(c echo.Context) error {
	var player dto.Player
	if err := c.Bind(&player); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := player.ID
	passcode := player.Passcode

	opponentPlayer, err := ph.ps.FindAvailableOpponentByPasscode(c.Request().Context(), id, passcode)
	if opponentPlayer == nil || err != nil {
		return c.JSON(http.StatusOK, err)
	}
	opponentPlayer.OpponentAnswer = player.Answer

	_, err = ph.ps.UpdateOpponentAnswer(c.Request().Context(), opponentPlayer.ID, opponentPlayer.OpponentAnswer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "回答を提出しました")
}
func (ph *playerHandler) FetchAnswersForComparison(c echo.Context) error {
	idStr := c.QueryParam("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newPlayer, err := ph.ps.GetPlayerByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	answersResp := dto.AnswersResp{
		AIAnswer:       newPlayer.AIAnswer,
		OpponentAnswer: newPlayer.OpponentAnswer,
	}
	return c.JSON(http.StatusOK, answersResp)
}

func (ph *playerHandler) CompareAnswerIsPlayer(c echo.Context) error {
	var player dto.Player
	if err := c.Bind(&player); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := player.ID
	selectAnswer := player.SelectAnswer

	newPlayer, err := ph.ps.GetPlayerByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if newPlayer.OpponentAnswer == selectAnswer {
		return c.JSON(http.StatusOK, "正解です")
	}
	return c.JSON(http.StatusOK, "不正解です")
}

func (ph *playerHandler) EndGame(c echo.Context) error {

	var request struct {
		ID uint `json:"id"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := ph.ps.DeletePlayerByID(c.Request().Context(), request.ID); err != nil {

		return c.JSON(http.StatusInternalServerError, err)

	}
	return c.JSON(http.StatusOK, "削除に成功しました")
}
