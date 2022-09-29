package project

import (
	"errors"
	"net/http"

	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/digicon-hack-07/ez-toon/server/utils/ulid"
	"github.com/labstack/echo/v4"
)

type GetProjectsResponse []Project

func (h *ProjectHandler) GetProjects(c echo.Context) error {
	projects, err := h.prjRepo.SelectProjects(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := GetProjectsResponse{}
	for _, project := range projects {
		totalPages, err := h.pageRepo.SelectProjectPageNum(c.Request().Context(), project.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}

		res = append(res, Project{
			ID:         project.ID,
			Name:       project.Name,
			TotalPages: totalPages,
			Thumbnail:  project.Thumbnail,
			CreatedAt:  project.CreatedAt,
			UpdatedAt:  project.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, res)
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

	project, err := h.prjRepo.InsertProject(c.Request().Context(), ulid.Make(), req.Name, "")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := PostProjectResponse{
		ID:         project.ID,
		Name:       project.Name,
		TotalPages: 0,
		Thumbnail:  project.Thumbnail,
		CreatedAt:  project.CreatedAt,
		UpdatedAt:  project.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, res)
}

type GetProjectResponse ProjectWithPages

func (h *ProjectHandler) GetProject(c echo.Context) error {
	id, err := ulid.Parse(c.Param("projectID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	project, err := h.prjRepo.SelectProject(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	pages, err := h.pageRepo.SelectProjectPages(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	resPages := []Page{}
	for _, page := range pages {
		resPages = append(resPages, Page{
			ID:     page.ID,
			Index:  page.Index,
			Height: page.Height,
			Width:  page.Width,
		})
	}

	res := GetProjectResponse{
		ID:        project.ID,
		Name:      project.Name,
		Thumbnail: project.Thumbnail,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
		Pages:     resPages,
	}

	return c.JSON(http.StatusOK, res)
}

type PatchProjectRequest struct {
	Name string `json:"name"`
}

type PatchProjectResponse Project

func (h *ProjectHandler) PatchProject(c echo.Context) error {
	id, err := ulid.Parse(c.Param("projectID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	req := PatchProjectRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	project, err := h.prjRepo.UpdateProject(c.Request().Context(), id, req.Name)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	totalPages, err := h.pageRepo.SelectProjectPageNum(c.Request().Context(), project.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	res := PatchProjectResponse{
		ID:         project.ID,
		Name:       project.Name,
		TotalPages: totalPages,
		Thumbnail:  project.Thumbnail,
		CreatedAt:  project.CreatedAt,
		UpdatedAt:  project.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *ProjectHandler) DeleteProject(c echo.Context) error {
	id, err := ulid.Parse(c.Param("projectID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.prjRepo.DeleteProject(c.Request().Context(), id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
