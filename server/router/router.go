package router

import (
	"net/http"

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

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	p := project.NewProjectHandler()

	api := e.Group("/api")
	{
		projectAPI := api.Group("/projects")
		{
			projectAPI.GET("", p.GetProjects)
			projectAPI.POST("", p.PostProject)
		}
	}

	e.Static("/", "public")

	return &Router{e: e}
}

func (r *Router) Start() {
	r.e.Logger.Panic(r.e.Start(":8080"))
}
