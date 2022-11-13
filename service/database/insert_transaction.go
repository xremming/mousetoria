package database

import (
	"context"
	"database/sql"

	service "github.com/xremming/mousetoria/service"
)

const (
	insertTransaction = `INSERT INTO "transaction" ("timestamp", "comment") VALUES (?, ?)`
	insertDebit       = `INSERT INTO "ledger" ("transactionID", "accountGroup", "accountID", "debit") VALUES (?, ?, ?, ?)`
	insertCredit      = `INSERT INTO "ledger" ("transactionID", "accountGroup", "accountID", "credit") VALUES (?, ?, ?, ?)`
)

func (db Database) InsertTransaction(ctx context.Context, transaction *service.Transaction) (*InsertTransactionResponse, error) {
	return withTransaction(ctx, db.db, func(ctx context.Context, tx *sql.Tx) (*InsertTransactionResponse, error) {
		res, err := tx.ExecContext(
			ctx,
			insertTransaction,
			transaction.GetTimestamp().GetDay(),
			transaction.GetTimestamp().GetTimeOfDay(),
			transaction.GetComment(),
		)
		if err != nil {
			return nil, err
		}

		transactionID, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}

		stmtDebit, err := tx.PrepareContext(ctx, insertDebit)
		if err != nil {
			return nil, err
		}

		for _, debit := range transaction.GetRecordsDebit() {
			_, err = stmtDebit.ExecContext(
				ctx,
				transactionID,
				debit.GetAccountNumber().GetGroup(),
				debit.GetAccountNumber().GetNumber(),
				debit.Value,
			)

			if err != nil {
				return nil, err
			}
		}

		stmtCredit, err := tx.PrepareContext(ctx, insertCredit)
		if err != nil {
			return nil, err
		}

		for _, credit := range transaction.GetRecordsCredit() {
			_, err = stmtCredit.ExecContext(
				ctx,
				transactionID,
				credit.GetAccountNumber().GetGroup(),
				credit.GetAccountNumber().GetNumber(),
				credit.Value,
			)

			if err != nil {
				return nil, err
			}
		}

		return &InsertTransactionResponse{TransactionId: transactionID}, nil
	})
}
