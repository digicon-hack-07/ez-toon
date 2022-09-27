package router

import (
	"net/http"

	"github.com/digicon-hack-07/ez-toon/server/router/page"
	"github.com/digicon-hack-07/ez-toon/server/router/project"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Router struct {
	e *echo.Echo
}

func NewRouter() *Router {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${time_rfc3339} ${prefix} ${short_file} ${line} |")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} method = ${method} | uri = ${uri} | status = ${status} ${error}\n",
	}))

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	prj := project.NewProjectHandler()
	page := page.NewPageHandler()

	api := e.Group("/api")
	{
		api.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})

		projectAPI := api.Group("/projects")
		{
			projectAPI.GET("", prj.GetProjects)
			projectAPI.POST("", prj.PostProject)
			projectAPI.GET("/:projectID", prj.GetProject)
			projectAPI.PATCH("/:projectID", prj.PatchProject)
			projectAPI.DELETE("/:projectID", prj.DeleteProject)
		}

		pageAPI := api.Group("/pages")
		{
			pageAPI.GET("/:pageID", page.GetPage)
			pageAPI.PATCH("/:pageID/index", page.PatchIndex)
		}
	}

	e.Static("/", "public")

	return &Router{e: e}
}

func (r *Router) Start() {
	r.e.Logger.Panic(r.e.Start(":8080"))
}
