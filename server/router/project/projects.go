package project

import (
	"time"

	"github.com/labstack/echo/v4"
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

func (h *ProjectHandler) GetProjects(c echo.Context) error {
	return c.JSON(200, []Project{
		{
			ID:        ulid.Make(),
			Name:      "test",
			Pages:     5,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        ulid.Make(),
			Name:      "test2",
			Pages:     5,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	})
}

func (h *ProjectHandler) PostProject(c echo.Context) error {
	project := Project{}
	if err := c.Bind(&project); err != nil {
		return err
	}

	project.ID = ulid.Make()

	return c.JSON(201, project)
}
