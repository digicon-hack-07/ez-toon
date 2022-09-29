package repository

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/oklog/ulid/v2"
)

type Point struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Pressure float64 `json:"pressure"`
}

type Points []Point

func (p *Points) Scan(value interface{}) error {
	b, ok := value.(string)
	if !ok {
		return errors.New("failed to scan Point")
	}

	return json.Unmarshal([]byte(b), p)
}

func (p Points) Value() (driver.Value, error) {
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

type Line struct {
	ID      ulid.ULID `gorm:"primaryKey;type:binary(16);not null"`
	PageID  ulid.ULID `gorm:"type:binary(16);not null"`
	PenSize int       `gorm:"type:int;not null"`
	Points  Points    `gorm:"type:text;not null"`
	Page    Page      `gorm:"foreignKey:PageID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type LinePageRepository interface {
	SelectLines(ctx context.Context, pageID ulid.ULID) ([]*Line, error)
}

type LineRepository interface {
	InsertLine(ctx context.Context, id ulid.ULID, pageID ulid.ULID, penSize int, points []Point) (*Line, error)
	DeleteLine(ctx context.Context, id ulid.ULID) error
}
