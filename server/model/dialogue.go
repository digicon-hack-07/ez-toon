package model

import "github.com/oklog/ulid/v2"

type Dialogue struct {
	ID       ulid.ULID
	PageID   ulid.ULID
	Dialogue string
}
