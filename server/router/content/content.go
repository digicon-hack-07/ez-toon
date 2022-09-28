package content

import "github.com/oklog/ulid/v2"

type Point struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Pressure float64 `json:"pressure"`
}

type Line struct {
	ID      ulid.ULID `json:"id"`
	PageID  ulid.ULID `json:"page_id"`
	PenSize int       `json:"penSize"`
	Points  []Point   `json:"points"`
}

type Dialogue struct {
	ID       ulid.ULID `json:"id"`
	PageID   ulid.ULID `json:"page_id"`
	Dialogue string    `json:"dialogue"`
	Top      float64   `json:"top"`
	Bottom   float64   `json:"bottom"`
	Left     float64   `json:"left"`
	Right    float64   `json:"right"`
}

type LineHandler struct{}

func NewLineHandler() *LineHandler {
	return &LineHandler{}
}

type DialogueHandler struct{}

func NewDialogueHandler() *DialogueHandler {
	return &DialogueHandler{}
}
