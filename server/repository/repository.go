package repository

import (
	"context"
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("not found")

type Repository interface {
	Do(context.Context, *sql.TxOptions, func(context.Context) error) error
}
