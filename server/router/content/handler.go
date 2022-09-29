package content

import (
	"errors"
	"net/http"

	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/digicon-hack-07/ez-toon/server/utils/ulid"
	"github.com/labstack/echo/v4"
)

type PostLineRequest struct {
	PageID  ulid.ULID `json:"page_id"`
	PenSize int       `json:"penSize"`
	Points  []Point   `json:"points"`
}

type PostLineResponse Line

func (h *LineHandler) PostLine(c echo.Context) error {
	req := PostLineRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	points := make(repository.Points, 0, len(req.Points))
	for _, p := range req.Points {
		points = append(points, repository.Point{
			X:        p.X,
			Y:        p.Y,
			Pressure: p.Pressure,
		})
	}

	line, err := h.repo.InsertLine(c.Request().Context(), ulid.Make(), req.PageID, req.PenSize, points)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := PostLineResponse{
		ID:      line.ID,
		PageID:  line.PageID,
		PenSize: line.PenSize,
		Points:  req.Points,
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *LineHandler) DeleteLine(c echo.Context) error {
	id, err := ulid.Parse(c.Param("lineID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.repo.DeleteLine(c.Request().Context(), id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

type PostDialogueRequest struct {
	PageID   ulid.ULID `json:"page_id"`
	Dialogue string    `json:"dialogue"`
	Top      float64   `json:"top"`
	Bottom   float64   `json:"bottom"`
	Left     float64   `json:"left"`
	Right    float64   `json:"right"`
}

type PostDialogueResponse Dialogue

func (h *DialogueHandler) PostDialogue(c echo.Context) error {
	req := PostDialogueRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	dial, err := h.repo.InsertDialogue(c.Request().Context(), ulid.Make(), req.PageID, req.Dialogue, req.Top, req.Bottom, req.Left, req.Right)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := PostDialogueResponse{
		ID:       dial.ID,
		PageID:   dial.PageID,
		Dialogue: dial.Dialogue,
		Top:      dial.Top,
		Bottom:   dial.Bottom,
		Left:     dial.Left,
		Right:    dial.Right,
	}

	return c.JSON(http.StatusCreated, res)
}

type PatchDialogueRequest struct {
	Dialogue string  `json:"dialogue"`
	Top      float64 `json:"top"`
	Bottom   float64 `json:"bottom"`
	Left     float64 `json:"left"`
	Right    float64 `json:"right"`
}

type PatchDialogueResponse Dialogue

func (h *DialogueHandler) PatchDialogue(c echo.Context) error {
	id, err := ulid.Parse(c.Param("dialogueID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	req := PatchDialogueRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	dial, err := h.repo.UpdateDialogue(c.Request().Context(), id, req.Dialogue, req.Top, req.Bottom, req.Left, req.Right)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := PatchDialogueResponse{
		ID:       dial.ID,
		PageID:   dial.PageID,
		Dialogue: dial.Dialogue,
		Top:      dial.Top,
		Bottom:   dial.Bottom,
		Left:     dial.Left,
		Right:    dial.Right,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *DialogueHandler) DeleteDialogue(c echo.Context) error {
	id, err := ulid.Parse(c.Param("dialogueID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.repo.DeleteDialogue(c.Request().Context(), id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
