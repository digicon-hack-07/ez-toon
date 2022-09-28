//go:build wireinject
// +build wireinject

package router

import (
	"github.com/digicon-hack-07/ez-toon/server/config"
	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/digicon-hack-07/ez-toon/server/repository/gorm2"
	"github.com/digicon-hack-07/ez-toon/server/router/content"
	"github.com/digicon-hack-07/ez-toon/server/router/page"
	"github.com/digicon-hack-07/ez-toon/server/router/project"
	"github.com/google/wire"
)

var SuperSet = wire.NewSet(
	gorm2.GetGorm2Config,
	gorm2.NewGorm2Repository,
	gorm2.GetSQLDb,

	wire.Bind(new(repository.ProjectRepository), new(*gorm2.Repository)),
	wire.Bind(new(repository.ProjectPageRepository), new(*gorm2.Repository)),
	wire.Bind(new(repository.PageRepository), new(*gorm2.Repository)),
	wire.Bind(new(repository.LinePageRepository), new(*gorm2.Repository)),
	wire.Bind(new(repository.DialoguePageRepository), new(*gorm2.Repository)),

	project.NewProjectHandler,
	page.NewPageHandler,
	content.NewLineHandler,
	content.NewDialogueHandler,

	NewRouter,
)

func InjectRouter(c *config.Config) (*Router, error) {
	wire.Build(SuperSet)
	return nil, nil
}
