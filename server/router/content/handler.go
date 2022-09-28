package content

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
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

	return c.JSON(http.StatusCreated, PostLineResponse{
		ID:      ulid.Make(),
		PageID:  req.PageID,
		PenSize: req.PenSize,
		Points:  req.Points,
	})
}

func (h *LineHandler) DeleteLine(c echo.Context) error {
	id := c.Param("lineID")
	c.Logger().Debug(id)

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

	return c.JSON(http.StatusCreated, PostDialogueResponse{
		ID:       ulid.Make(),
		Dialogue: req.Dialogue,
		Top:      req.Top,
		Bottom:   req.Bottom,
		Left:     req.Left,
		Right:    req.Right,
	})
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
	id := c.Param("dialogueID")

	req := PatchDialogueRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, PatchDialogueResponse{
		ID:       ulid.MustParse(id),
		Dialogue: req.Dialogue,
		Top:      req.Top,
		Bottom:   req.Bottom,
		Left:     req.Left,
		Right:    req.Right,
	})
}

func (h *DialogueHandler) DeleteDialogue(c echo.Context) error {
	id := c.Param("dialogueID")
	c.Logger().Debug(id)

	return c.NoContent(http.StatusNoContent)
}
