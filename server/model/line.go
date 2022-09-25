package model

import "github.com/oklog/ulid/v2"

type Line struct {
	ID     ulid.ULID
	PageID ulid.ULID
}
