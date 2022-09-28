// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func InjectRouter(c *config.Config) (*Router, error) {
	gorm2Config := gorm2.GetGorm2Config(c)
	repository, err := gorm2.NewGorm2Repository(gorm2Config)
	if err != nil {
		return nil, err
	}
	db, err := gorm2.GetSQLDb(repository)
	if err != nil {
		return nil, err
	}
	projectHandler := project.NewProjectHandler(repository, repository)
	pageHandler := page.NewPageHandler()
	lineHandler := content.NewLineHandler()
	dialogueHandler := content.NewDialogueHandler()
	router, err := NewRouter(db, projectHandler, pageHandler, lineHandler, dialogueHandler)
	if err != nil {
		return nil, err
	}
	return router, nil
}

// wire.go:

var SuperSet = wire.NewSet(gorm2.GetGorm2Config, gorm2.NewGorm2Repository, gorm2.GetSQLDb, wire.Bind(new(repository.ProjectRepository), new(*gorm2.Repository)), wire.Bind(new(repository.ProjectPageRepository), new(*gorm2.Repository)), project.NewProjectHandler, page.NewPageHandler, content.NewLineHandler, content.NewDialogueHandler, NewRouter)
