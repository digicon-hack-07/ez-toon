package project

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Project struct {
	ID        ulid.ULID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Pages     int       `json:"pages,omitempty"`
	Thumbnail string    `json:"thumbnail,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ProjectHandler struct{}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}
