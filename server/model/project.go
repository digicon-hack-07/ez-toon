package model

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Project struct {
	ID         ulid.ULID
	Name       string
	TotalPages int
	Thumbnail  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
