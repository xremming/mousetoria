package db

import (
	"context"
	_ "embed"

	"github.com/rs/zerolog"
	"github.com/xremming/mousetoria/ledger"
)

//go:embed sql/upsert_account.sql
var upsertAccountSQL string

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
