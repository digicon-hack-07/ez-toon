package page

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

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
				Top:      150,
				Bottom:   50,
				Left:     250,
				Right:    100,
			},
		},
	})
}

type PatchIndexRequest struct {
	// Operationは"inc"(加算)か"dec"(減算)のみを許可する
	Operation string `json:"operation,omitempty"`
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
