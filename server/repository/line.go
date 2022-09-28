package repository

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type Point struct {
	X        float64
	Y        float64
	Pressure float64
}

type Line struct {
	ID      ulid.ULID `gorm:"primaryKey;type:char(26);not null"`
	PageID  ulid.ULID `gorm:"type:char(26);not null"`
	PenSize int       `gorm:"type:int;not null"`
	Points  []Point   `gorm:"type:json;not null"`
	Page    Page      `gorm:"foreignKey:PageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type LinePageRepository interface {
	SelectLines(ctx context.Context, pageID ulid.ULID) ([]*Line, error)
}
