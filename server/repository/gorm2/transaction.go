package gorm2

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ctxKey string

const txKey ctxKey = "transaction"

func (repo *Repository) Do(ctx context.Context, options *sql.TxOptions, callBack func(context.Context) error) error {
	txFunc := func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, txKey, tx)

		err := callBack(ctx)
		if err != nil {
			return err
		}

		err = ctx.Err()
		if err != nil {
			return err
		}

		return nil
	}

	if options == nil {
		err := repo.db.Transaction(txFunc)
		if err != nil {
			return fmt.Errorf("failed to execute callback: %w", err)
		}
	} else {
		err := repo.db.Transaction(txFunc, options)
		if err != nil {
			return fmt.Errorf("failed to execute callback:%w", err)
		}
	}

	return nil
}

func (repo *Repository) getTX(ctx context.Context) (*gorm.DB, error) {
	txInterface := ctx.Value(txKey)
	if txInterface == nil {
		return repo.db.WithContext(ctx), nil
	}

	tx, ok := txInterface.(*gorm.DB)
	if !ok {
		return nil, errors.New("failed to cast DB instance to *gorm.DB")
	}

	return tx, nil
}
