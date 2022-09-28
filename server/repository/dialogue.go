package repository

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type Dialogue struct {
	ID       ulid.ULID `gorm:"primaryKey;type:char(26);not null"`
	PageID   ulid.ULID `gorm:"type:char(26);not null"`
	Dialogue string    `gorm:"type:text;not null"`
	Top      float64   `gorm:"type:float;not null"`
	Bottom   float64   `gorm:"type:float;not null"`
	Left     float64   `gorm:"type:float;not null"`
	Right    float64   `gorm:"type:float;not null"`
	Page     Page      `gorm:"foreignKey:PageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type DialoguePageRepository interface {
	SelectDialogues(ctx context.Context, pageID ulid.ULID) ([]*Dialogue, error)
}
