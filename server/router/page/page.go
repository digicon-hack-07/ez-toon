package page

import (
	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/digicon-hack-07/ez-toon/server/utils/ulid"
)

type Page struct {
	ID        ulid.ULID `json:"id"`
	ProjectID ulid.ULID `json:"project_id"`
	Index     int       `json:"index"`
	Height    int       `json:"height"`
	Width     int       `json:"width"`
}

type Point struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Pressure float64 `json:"pressure"`
}

type Line struct {
	ID      ulid.ULID `json:"id"`
	PenSize int       `json:"penSize"`
	Points  []Point   `json:"points"`
}

type Dialogue struct {
	ID       ulid.ULID `json:"id"`
	Dialogue string    `json:"dialogue"`
	Top      float64   `json:"top"`
	Bottom   float64   `json:"bottom"`
	Left     float64   `json:"left"`
	Right    float64   `json:"right"`
}

type PageWithContents struct {
	ID        ulid.ULID  `json:"id"`
	ProjectID ulid.ULID  `json:"project_id"`
	Index     int        `json:"index"`
	Height    int        `json:"height"`
	Width     int        `json:"width"`
	Lines     []Line     `json:"lines"`
	Dialogues []Dialogue `json:"dialogues"`
}

type PageHandler struct {
	pageRepo repository.PageRepository
	lineRepo repository.LinePageRepository
	dialRepo repository.DialoguePageRepository
}

func NewPageHandler(
	pageRepo repository.PageRepository,
	lineRepo repository.LinePageRepository,
	dialRepo repository.DialoguePageRepository,
) *PageHandler {
	return &PageHandler{
		pageRepo: pageRepo,
		lineRepo: lineRepo,
		dialRepo: dialRepo,
	}
}
