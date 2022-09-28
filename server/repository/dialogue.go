package repository

import "github.com/oklog/ulid/v2"

type Dialogue struct {
	ID       ulid.ULID
	PageID   ulid.ULID
	Dialogue string
	Top      float64
	Bottom   float64
	Left     float64
	Right    float64
}
