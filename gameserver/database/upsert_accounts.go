package database

import (
	"context"
	_ "embed"

	"github.com/rs/zerolog"
)

const upsertAccountSQL = `INSERT OR REPLACE INTO "account" ("accountGroup", "accountID", "name") VALUES (?, ?, ?)`

func (db Database) upsertAllAccounts(ctx context.Context) error {
	zerolog.Ctx(ctx).Info().Msg("upserting all accounts and account names")

	stmt, err := db.db.PrepareContext(ctx, upsertAccountSQL)
	if err != nil {
		return err
	}

	for accountGroup, accounts := range accounts {
		for accountID, name := range accounts {
			_, err = stmt.ExecContext(ctx, accountGroup, accountID, name)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
