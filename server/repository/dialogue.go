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

type DialogueRepository interface {
	InsertDialogue(ctx context.Context, id ulid.ULID, pageID ulid.ULID, dialogue string, top float64, bottom float64, left float64, right float64) (*Dialogue, error)
	UpdateDialogue(ctx context.Context, id ulid.ULID, dialogue string, top float64, bottom float64, left float64, right float64) (*Dialogue, error)
	DeleteDialogue(ctx context.Context, id ulid.ULID) error
}
