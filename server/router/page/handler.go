package page

import (
	"net/http"

	"github.com/digicon-hack-07/ez-toon/server/router/content"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type GetPageResponse PageWithContents

func (h *PageHandler) GetPage(c echo.Context) error {
	return c.JSON(http.StatusOK, GetPageResponse{
		ID:     ulid.Make(),
		Index:  0,
		Height: 300,
		Width:  400,
		Lines: []content.Line{
			{
				ID:      ulid.Make(),
				PenSize: 3,
				Points: []content.Point{
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
				Points: []content.Point{
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
		Dialogues: []content.Dialogue{
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
