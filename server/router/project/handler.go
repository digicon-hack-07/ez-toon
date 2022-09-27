package project

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type GetProjectsResponse []Project

func (h *ProjectHandler) GetProjects(c echo.Context) error {
	return c.JSON(http.StatusOK, GetProjectsResponse{
		{
			ID:         ulid.Make(),
			Name:       "test",
			TotalPages: 5,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			ID:         ulid.Make(),
			Name:       "test2",
			TotalPages: 5,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	})
}

type PostProjectRequest struct {
	Name string `json:"name,omitempty"`
}

type PostProjectResponse Project

func (h *ProjectHandler) PostProject(c echo.Context) error {
	req := PostProjectRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	res := PostProjectResponse{
		ID:         ulid.Make(),
		Name:       req.Name,
		TotalPages: 0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return c.JSON(http.StatusCreated, res)
}

type PatchProjectRequest struct {
	Name string `json:"name,omitempty"`
}

func (h *ProjectHandler) PatchProject(c echo.Context) error {
	id := c.Param("projectID")
	c.Logger().Debug(id)

	req := PatchProjectRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	res := PostProjectResponse{
		ID:         ulid.Make(),
		Name:       req.Name,
		TotalPages: 0,
		CreatedAt:  time.Now().Add(-time.Hour),
		UpdatedAt:  time.Now(),
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *ProjectHandler) DeleteProject(c echo.Context) error {
	id := c.Param("projectID")
	c.Logger().Debug(id)

	return c.NoContent(http.StatusNoContent)
}
