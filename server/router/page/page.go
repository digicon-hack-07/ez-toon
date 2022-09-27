package page

import (
	"github.com/digicon-hack-07/ez-toon/server/router/content"
	"github.com/oklog/ulid/v2"
)

type Page struct {
	ID     ulid.ULID `json:"id,omitempty"`
	Index  int       `json:"index,omitempty"`
	Height int       `json:"height,omitempty"`
	Width  int       `json:"width,omitempty"`
}

type PageWithContents struct {
	ID        ulid.ULID          `json:"id,omitempty"`
	Index     int                `json:"index,omitempty"`
	Height    int                `json:"height,omitempty"`
	Width     int                `json:"width,omitempty"`
	Lines     []content.Line     `json:"lines,omitempty"`
	Dialogues []content.Dialogue `json:"dialogues,omitempty"`
}

type PageHandler struct{}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}
