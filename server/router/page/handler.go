package page

import (
	"errors"
	"net/http"

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
	id := c.Param("pageID")

	return c.JSON(http.StatusOK, GetPageResponse{
		ID:        ulid.MustParse(id),
		ProjectID: ulid.Make(),
		Index:     0,
		Height:    300,
		Width:     400,
		Lines: []Line{
			{
				ID:      ulid.Make(),
				PenSize: 3,
				Points: []Point{
					{
						X:        0,
						Y:        0,
						Pressure: 1.500,
					},
					{
						X:        14.7,
						Y:        22.3,
						Pressure: 2.500,
					},
					{
						X:        44.7,
						Y:        35.2,
						Pressure: 1.500,
					},
				},
			},
			{
				ID:      ulid.Make(),
				PenSize: 4,
				Points: []Point{
					{
						X:        100.2,
						Y:        353.6,
						Pressure: 1.500,
					},
					{
						X:        114.6,
						Y:        373.3,
						Pressure: 2.500,
					},
					{
						X:        125.7,
						Y:        367.3,
						Pressure: 1.500,
					},
				},
			},
		},
		Dialogues: []Dialogue{
			{
				ID:       ulid.Make(),
				Dialogue: "あいうえおかきくけこ",
				Top:      50,
				Bottom:   150,
				Left:     100,
				Right:    250,
			},
		},
	})
}

type PatchIndexRequest struct {
	// Operationは"inc"(加算)か"dec"(減算)のみを許可する
	Operation string `json:"operation"`
}

// 更新したページ一覧を返却する
type PatchIndexResponse []Page

func (h *PageHandler) PatchIndex(c echo.Context) error {
	id := c.Param("pageID")

	req := PatchIndexRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if req.Operation != "inc" && req.Operation != "dec" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid operation"))
	}

	return c.JSON(http.StatusOK, PatchIndexResponse{
		{
			ID:        ulid.Make(),
			ProjectID: ulid.Make(),
			Index:     0,
			Height:    300,
			Width:     400,
		},
		{
			ID:        ulid.MustParse(id),
			ProjectID: ulid.Make(),
			Index:     1,
			Height:    500,
			Width:     600,
		},
	})
}