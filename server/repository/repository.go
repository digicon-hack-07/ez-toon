package repository

import (
	"context"
	"database/sql"
)

type Repository interface {
	Do(context.Context, *sql.TxOptions, func(context.Context) error) error
}
