package repository

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type Page struct {
	ID        ulid.ULID `gorm:"type:char(26);not null;primaryKey"`
	ProjectID ulid.ULID `gorm:"type:char(26);not null"`
	Index     int       `gorm:"type:int;not null"`
	Height    int       `gorm:"type:int;not null"`
	Width     int       `gorm:"type:int;not null"`
	Project   Project   `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ProjectPageRepository interface {
	SelectProjectPageNum(ctx context.Context, projectID ulid.ULID) (int, error)
	SelectProjectPages(ctx context.Context, projectID ulid.ULID) ([]*Page, error)
}
