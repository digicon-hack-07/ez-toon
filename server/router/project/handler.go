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
	Name string `json:"name"`
}

type PostProjectResponse Project

func (h *ProjectHandler) PostProject(c echo.Context) error {
	req := PostProjectRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
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

type GetProjectResponse ProjectWithPages

func (h *ProjectHandler) GetProject(c echo.Context) error {
	id := c.Param("projectID")

	return c.JSON(http.StatusOK, GetProjectResponse{
		ID:        ulid.MustParse(id),
		Name:      "test",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Pages: []Page{
			{
				ID:     ulid.Make(),
				Index:  0,
				Height: 300,
				Width:  400,
			},
			{
				ID:     ulid.Make(),
				Index:  1,
				Height: 500,
				Width:  600,
			},
		},
	})
}

type PatchProjectRequest struct {
	Name string `json:"name"`
}

type PatchProjectResponse Project

func (h *ProjectHandler) PatchProject(c echo.Context) error {
	id := c.Param("projectID")

	req := PatchProjectRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res := PatchProjectResponse{
		ID:         ulid.MustParse(id),
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
