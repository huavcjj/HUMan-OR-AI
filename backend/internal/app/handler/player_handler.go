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
	CompareAnswerIsPlayer(c echo.Context) error     // 選択した回答がAIか人間かを判定する
	IsOpponentAnswerByPlayer(c echo.Context) error  // 相手が選択した回答がAIか人間かの情報を取得する
	EndGame(c echo.Context) error                   // ゲームを終了する
}

type playerHandler struct {
	ps service.IPlayerService
}

func NewPlayerHandler(ps service.IPlayerService) IPlayerHandler {
	return &playerHandler{ps: ps}
}

func (ph *playerHandler) StartNewGame(c echo.Context) error {
	var passcodeReq dto.PasscodeReq
	if err := c.Bind(&passcodeReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if passcodeReq.Passcode == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "合言葉を入力してください",
		})
	}

	player := dto.NewPlayer(passcodeReq.Passcode)

	newPlayer, err := ph.ps.CreatePlayer(c.Request().Context(), player)
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
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "対戦相手が見つかりませんでした",
		})
	}

	passcodeResp := dto.PasscodeResp{
		ID:       newPlayer.ID,
		Passcode: newPlayer.Passcode,
	}
	return c.JSON(http.StatusOK, passcodeResp)
}

func (ph *playerHandler) SubmitPlayerTopic(c echo.Context) error {
	var topicReq dto.TopicReq

	if err := c.Bind(&topicReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id := topicReq.ID
	topic := topicReq.Topic

	_, err := ph.ps.UpdateTopicAndAIAnswer(c.Request().Context(), id, topic)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
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
		return c.JSON(http.StatusNoContent, err)
	}

	topicResp := dto.TopicResp{
		Topic: opponentPlayer.Topic,
	}

	return c.JSON(http.StatusOK, topicResp)
}

func (ph *playerHandler) SubmitAnswerToOpponent(c echo.Context) error {
	var answerReq dto.AnswerReq
	if err := c.Bind(&answerReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := answerReq.ID
	passcode := answerReq.Passcode

	opponentPlayer, err := ph.ps.FindAvailableOpponentByPasscode(c.Request().Context(), id, passcode)
	if opponentPlayer == nil || err != nil {
		return c.JSON(http.StatusOK, err)
	}
	opponentPlayer.OpponentAnswer = answerReq.Answer

	_, err = ph.ps.UpdateOpponentAnswer(c.Request().Context(), opponentPlayer.ID, opponentPlayer.OpponentAnswer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}
func (ph *playerHandler) FetchAnswersForComparison(c echo.Context) error {
	idStr := c.QueryParam("id")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	time.Sleep(5 * time.Second)

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
	var answerIsPlayerReq dto.AnswerIsPlayerReq
	if err := c.Bind(&answerIsPlayerReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	id := answerIsPlayerReq.ID
	selectAnswer := answerIsPlayerReq.SelectAnswer

	opponentPlayer, err := ph.ps.GetPlayerByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if opponentPlayer.OpponentAnswer == selectAnswer {
		if _, err := ph.ps.UpdateSelectAnswerIsPlayer(c.Request().Context(), opponentPlayer.ID, true); err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, true)
	}
	if _, err := ph.ps.UpdateSelectAnswerIsPlayer(c.Request().Context(), opponentPlayer.ID, false); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, false)
}

func (ph *playerHandler) IsOpponentAnswerByPlayer(c echo.Context) error {
	idStr := c.QueryParam("id")
	passcode := c.QueryParam("passcode")

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	player, err := ph.ps.GetPlayerByID(c.Request().Context(), uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	decodedPasscode, err := url.QueryUnescape(passcode)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	opponentPlayer, err := ph.ps.FindAvailableOpponentByPasscode(c.Request().Context(), uint(id), decodedPasscode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	pponentAnswerIsPlayerResp := dto.OpponentAnswerIsPlayerResp{
		Topic:    player.Topic,
		Anser:    player.OpponentAnswer,
		AIAnswer: player.AIAnswer,
		IsPlayer: opponentPlayer.SelectAnswerIsPlayer,
	}
	return c.JSON(http.StatusOK, pponentAnswerIsPlayerResp)

}

func (ph *playerHandler) EndGame(c echo.Context) error {
	var idReq dto.IDReq
	if err := c.Bind(&idReq); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := ph.ps.DeletePlayerByID(c.Request().Context(), idReq.ID); err != nil {

		return c.JSON(http.StatusInternalServerError, err)

	}
	return c.NoContent(http.StatusOK)
}
