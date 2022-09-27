package page

import "github.com/oklog/ulid/v2"

type Page struct {
	ID     ulid.ULID
	Index  int
	Height int
	Width  int
}
