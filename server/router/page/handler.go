package page

import (
	"errors"
	"net/http"

	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type PostPageRequest struct {
	ID        ulid.ULID `json:"id"`
	ProjectID ulid.ULID `json:"project_id"`
	Height    int       `json:"height"`
	Width     int       `json:"width"`
}

type PostPageResponse Page

func (h *PageHandler) PostPage(c echo.Context) error {
	req := PostPageRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	id, err := ulid.Parse(c.Param("pageID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	prjID, err := ulid.Parse(c.Param("projectID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	page, err := h.pageRepo.InsertPage(c.Request().Context(), id, prjID, req.Height, req.Width)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := PostPageResponse{
		ID:        page.ID,
		ProjectID: page.ProjectID,
		Index:     page.Index,
		Height:    page.Height,
		Width:     page.Width,
	}

	return c.JSON(http.StatusCreated, res)
}

type GetPageResponse PageWithContents

func (h *PageHandler) GetPage(c echo.Context) error {
	id, err := ulid.Parse(c.Param("pageID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	page, err := h.pageRepo.SelectPage(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	lines, err := h.lineRepo.SelectLines(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	dials, err := h.dialRepo.SelectDialogues(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	resLines := make([]Line, 0, len(lines))
	for _, line := range lines {
		resPoints := make([]Point, 0, len(line.Points))
		for _, point := range line.Points {
			resPoints = append(resPoints, Point{
				X:        point.X,
				Y:        point.Y,
				Pressure: point.Pressure,
			})
		}

		resLines = append(resLines, Line{
			ID:      line.ID,
			PenSize: line.PenSize,
			Points:  resPoints,
		})
	}

	resDials := make([]Dialogue, 0, len(dials))
	for _, dial := range dials {
		resDials = append(resDials, Dialogue{
			ID:       dial.ID,
			Dialogue: dial.Dialogue,
			Top:      dial.Top,
			Bottom:   dial.Bottom,
			Left:     dial.Left,
			Right:    dial.Right,
		})
	}

	res := GetPageResponse{
		ID:        page.ID,
		ProjectID: page.ProjectID,
		Index:     page.Index,
		Height:    page.Height,
		Width:     page.Width,
		Lines:     resLines,
		Dialogues: resDials,
	}

	return c.JSON(http.StatusOK, res)
}

type PatchIndexRequest struct {
	// Operationは"inc"(加算)か"dec"(減算)のみを許可する
	Operation string `json:"operation"`
}

// 更新したページ一覧を返却する
type PatchIndexResponse []Page

func (h *PageHandler) PatchIndex(c echo.Context) error {
	id, err := ulid.Parse(c.Param("pageID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	req := PatchIndexRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if req.Operation != "inc" && req.Operation != "dec" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid operation"))
	}

	page, err := h.pageRepo.UpdateIndex(c.Request().Context(), id, req.Operation)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	pages, err := h.pageRepo.SelectProjectPages(c.Request().Context(), page.ProjectID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := []Page{}
	for _, page := range pages {
		res = append(res, Page{
			ID:     page.ID,
			Index:  page.Index,
			Height: page.Height,
			Width:  page.Width,
		})
	}

	return c.JSON(http.StatusOK, res)
}
