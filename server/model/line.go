package model

import "github.com/oklog/ulid/v2"

type Point struct {
	X        float64
	Y        float64
	Pressure float64
}

type Line struct {
	ID      ulid.ULID
	PageID  ulid.ULID
	PenSize int
	Points  []Point
}
