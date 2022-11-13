package database

import (
	"context"
	"database/sql"
	"gameserver/proto"
)

const (
	insertTransactionSQL = `INSERT INTO "transaction" ("timestamp", "comment") VALUES (?, ?)`
	insertDebitSQL       = `INSERT INTO "ledger" ("transactionID", "accountGroup", "accountID", "debit") VALUES (?, ?, ?, ?)`
	insertCreditSQL      = `INSERT INTO "ledger" ("transactionID", "accountGroup", "accountID", "credit") VALUES (?, ?, ?, ?)`
)

func (db Database) InsertTransaction(ctx context.Context, transaction *proto.Transaction) (*proto.InsertTransactionResponse, error) {
	return withTransaction(ctx, db.db, func(ctx context.Context, tx *sql.Tx) (*proto.InsertTransactionResponse, error) {
		res, err := tx.ExecContext(
			ctx,
			insertTransactionSQL,
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

		stmtDebit, err := tx.PrepareContext(ctx, insertDebitSQL)
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

		stmtCredit, err := tx.PrepareContext(ctx, insertCreditSQL)
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

		return &proto.InsertTransactionResponse{TransactionId: transactionID}, nil
	})
}
