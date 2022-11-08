package db

import (
	"context"
	"database/sql"

	"github.com/hashicorp/go-multierror"
)

func withTransaction[T any](ctx context.Context, db *sql.DB, f func(ctx context.Context, tx *sql.Tx) (T, error)) (T, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return *new(T), err
	}

	out, err := f(ctx, tx)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			err = multierror.Append(err, errRollback)
		}

		return *new(T), err
	}

	return out, tx.Commit()
}

func withTransaction2(ctx context.Context, db *sql.DB, f func(ctx context.Context, tx *sql.Tx) error) error {
	_, err := withTransaction(ctx, db, func(ctx context.Context, tx *sql.Tx) (struct{}, error) {
		return struct{}{}, f(ctx, tx)
	})
	if err != nil {
		return err
	}

	return nil
}
