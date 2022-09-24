package model

import "github.com/oklog/ulid/v2"

type Project struct {
	ID        ulid.ULID
	Name      string
	Pages     int
	Thumbnail string
}
