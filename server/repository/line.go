package repository

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/digicon-hack-07/ez-toon/server/utils/ulid"
)

type Point struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Pressure float64 `json:"pressure"`
}

type Points []Point

func (p *Points) Scan(value interface{}) error {
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		return json.Unmarshal([]byte(v), p)
	case []byte:
		return json.Unmarshal(v, p)
	}

	return errors.New("invalid type")
}

func (p Points) Value() (driver.Value, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

type Line struct {
	ID      ulid.ULID `gorm:"type:char(26);not null;primaryKey"`
	PageID  ulid.ULID `gorm:"type:char(26);not null"`
	PenSize int       `gorm:"type:int;not null"`
	Points  Points    `gorm:"type:text;not null"`
	Page    Page      `gorm:"foreignKey:PageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type LinePageRepository interface {
	SelectLines(ctx context.Context, pageID ulid.ULID) ([]*Line, error)
}

type LineRepository interface {
	InsertLine(ctx context.Context, id ulid.ULID, pageID ulid.ULID, penSize int, points Points) (*Line, error)
	DeleteLine(ctx context.Context, id ulid.ULID) error
}
