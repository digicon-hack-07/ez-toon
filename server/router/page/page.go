package page

import (
	"github.com/oklog/ulid/v2"
)

type Page struct {
	ID        ulid.ULID `json:"id,omitempty"`
	ProjectID ulid.ULID `json:"project_id,omitempty"`
	Index     int       `json:"index,omitempty"`
	Height    int       `json:"height,omitempty"`
	Width     int       `json:"width,omitempty"`
}

type Point struct {
	X        float64 `json:"x,omitempty"`
	Y        float64 `json:"y,omitempty"`
	Pressure float64 `json:"pressure,omitempty"`
}

type Line struct {
	ID      ulid.ULID `json:"id,omitempty"`
	PenSize int       `json:"penSize,omitempty"`
	Points  []Point   `json:"points,omitempty"`
}

type Dialogue struct {
	ID       ulid.ULID `json:"id,omitempty"`
	Dialogue string    `json:"dialogue,omitempty"`
	Top      float64   `json:"top,omitempty"`
	Bottom   float64   `json:"bottom,omitempty"`
	Left     float64   `json:"left,omitempty"`
	Right    float64   `json:"right,omitempty"`
}

type PageWithContents struct {
	ID        ulid.ULID  `json:"id,omitempty"`
	ProjectID ulid.ULID  `json:"project_id,omitempty"`
	Index     int        `json:"index,omitempty"`
	Height    int        `json:"height,omitempty"`
	Width     int        `json:"width,omitempty"`
	Lines     []Line     `json:"lines,omitempty"`
	Dialogues []Dialogue `json:"dialogues,omitempty"`
}

type PageHandler struct{}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}
