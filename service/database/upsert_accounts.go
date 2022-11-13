package database

import (
	"context"
	_ "embed"

	"github.com/rs/zerolog"
	"github.com/xremming/mousetoria/ledger"
)

const upsertAccountSQL = `INSERT OR REPLACE INTO "account" ("accountGroup", "accountID", "name") VALUES (?, ?, ?)`

func (db Database) upsertAllAccounts(ctx context.Context) error {
	zerolog.Ctx(ctx).Info().Msg("upserting all accounts and account names")

	stmt, err := db.db.PrepareContext(ctx, upsertAccountSQL)
	if err != nil {
		return err
	}

	for account, name := range ledger.AccountNames() {
		_, err = stmt.ExecContext(ctx, account.AccountGroup, account.AccountID, name)
		if err != nil {
			return err
		}
	}

	return nil
}
