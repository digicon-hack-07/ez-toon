package router

import (
	"database/sql"
	"net/http"

	"github.com/digicon-hack-07/ez-toon/server/router/content"
	"github.com/digicon-hack-07/ez-toon/server/router/page"
	"github.com/digicon-hack-07/ez-toon/server/router/project"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/srinathgs/mysqlstore"
)

type Router struct {
	e *echo.Echo
}

func NewRouter(
	db *sql.DB,
	prj *project.ProjectHandler,
	page *page.PageHandler,
	line *content.LineHandler,
	dial *content.DialogueHandler,
) (*Router, error) {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetHeader("${time_rfc3339} ${prefix} ${short_file} ${line} |")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} method = ${method} | uri = ${uri} | status = ${status} ${error}\n",
	}))

	store, err := mysqlstore.NewMySQLStoreFromConnection(db, "sessions", "/", 60*60*24*14, []byte("secret-key"))
	if err != nil {
		return nil, err
	}
	e.Use(session.Middleware(store))

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
			pageAPI.POST("", page.PostPage)
			pageAPI.GET("/:pageID", page.GetPage)
			pageAPI.PATCH("/:pageID/index", page.PatchIndex)
		}

		lineAPI := api.Group("/lines")
		{
			lineAPI.POST("", line.PostLine)
			lineAPI.DELETE("/:lineID", line.DeleteLine)
		}

		dialogueAPI := api.Group("/dialogues")
		{
			dialogueAPI.POST("", dial.PostDialogue)
			dialogueAPI.PATCH("/:dialogueID", dial.PatchDialogue)
			dialogueAPI.DELETE("/:dialogueID", dial.DeleteDialogue)
		}
	}

	e.Static("/", "public")
	e.File("/*", "public/index.html")

	return &Router{e: e}, nil
}

func (r *Router) Start() {
	r.e.Logger.Panic(r.e.Start(":8080"))
}
