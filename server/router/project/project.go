package project

import (
	"time"

	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/digicon-hack-07/ez-toon/server/utils/ulid"
)

type Project struct {
	ID         ulid.ULID `json:"id"`
	Name       string    `json:"name"`
	TotalPages int       `json:"pages"`
	Thumbnail  string    `json:"thumbnail"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Page struct {
	ID     ulid.ULID `json:"id"`
	Index  int       `json:"index"`
	Height int       `json:"height"`
	Width  int       `json:"width"`
}

type ProjectWithPages struct {
	ID        ulid.ULID `json:"id"`
	Name      string    `json:"name"`
	Pages     []Page    `json:"pages"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectHandler struct {
	prjRepo  repository.ProjectRepository
	pageRepo repository.ProjectPageRepository
}

func NewProjectHandler(prjRepo repository.ProjectRepository, pageRepo repository.ProjectPageRepository) *ProjectHandler {
	return &ProjectHandler{
		prjRepo:  prjRepo,
		pageRepo: pageRepo,
	}
}
