package content

import "github.com/oklog/ulid/v2"

type Point struct {
	X        float64 `json:"x,omitempty"`
	Y        float64 `json:"y,omitempty"`
	Pressure float64 `json:"pressure,omitempty"`
}

type Line struct {
	ID      ulid.ULID `json:"id,omitempty"`
	PageID  ulid.ULID `json:"page_id,omitempty"`
	PenSize int       `json:"penSize,omitempty"`
	Points  []Point   `json:"points,omitempty"`
}

type Dialogue struct {
	ID       ulid.ULID `json:"id,omitempty"`
	PageID   ulid.ULID `json:"page_id,omitempty"`
	Dialogue string    `json:"dialogue,omitempty"`
	Top      float64   `json:"top,omitempty"`
	Bottom   float64   `json:"bottom,omitempty"`
	Left     float64   `json:"left,omitempty"`
	Right    float64   `json:"right,omitempty"`
}
