package project

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Project struct {
	ID         ulid.ULID `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	TotalPages int       `json:"pages,omitempty"`
	Thumbnail  string    `json:"thumbnail,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type Page struct {
	ID     ulid.ULID `json:"id,omitempty"`
	Index  int       `json:"index,omitempty"`
	Height int       `json:"height,omitempty"`
	Width  int       `json:"width,omitempty"`
}

type ProjectWithPages struct {
	ID        ulid.ULID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Pages     []Page    `json:"pages,omitempty"`
	Thumbnail string    `json:"thumbnail,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ProjectHandler struct{}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}
