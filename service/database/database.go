package database

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/hashicorp/go-multierror"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
)

type Database struct {
	UnimplementedDatabaseServer
	db *sql.DB
}

func NewDatabase(ctx context.Context, dataSourceName string) (Database, error) {
	logger := zerolog.Ctx(ctx).With().Str("dataSourceName", dataSourceName).Logger()

	logger.Info().Msg("creating database")

	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		logger.Err(err).Msg("failed to open database")
		return Database{}, err
	}

	closeDB := func(err error) error {
		errClose := db.Close()
		if errClose != nil {
			logger.Err(errClose).
				AnErr("originalError", err).
				Msg("failed to close database after an error happened")
			err = multierror.Append(errClose)
		}

		return err
	}

	out := Database{db: db}

	err = out.runMigrations(ctx)
	if err != nil {
		logger.Err(err).Msg("failed to run migrations")
		return Database{}, closeDB(err)
	}

	err = out.upsertAllAccounts(ctx)
	if err != nil {
		logger.Err(err).Msg("failed to upsert accounts and account names")
		return Database{}, closeDB(err)
	}

	return out, nil
}

func (db Database) Close() error {
	return db.db.Close()
}
