package content

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type PostLineRequest struct {
	PageID  ulid.ULID `json:"page_id,omitempty"`
	PenSize int       `json:"penSize,omitempty"`
	Points  []Point   `json:"points,omitempty"`
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
