package repository

import (
	"context"
	"time"

	"github.com/oklog/ulid/v2"
)

type Project struct {
	ID        ulid.ULID `gorm:"type:char(26);not null;primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Thumbnail string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
}

type ProjectRepository interface {
	SelectProjects(ctx context.Context) ([]*Project, error)
	CreateProject(ctx context.Context, id ulid.ULID, name string, thumbnail string) error
	SelectProject(ctx context.Context, id ulid.ULID) (*Project, error)
	UpdateProject(ctx context.Context, id ulid.ULID, name string) (*Project, error)
	DeleteProject(ctx context.Context, id ulid.ULID) error
}
