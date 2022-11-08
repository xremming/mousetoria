package db

import (
	"context"
	"database/sql"

	"github.com/xremming/mousetoria/ledger"
)

func (db Database) GetTransaction(ctx context.Context, transactionID int64) (ledger.Transaction, error) {
	return withTransaction(ctx, db.db, func(ctx context.Context, tx *sql.Tx) (ledger.Transaction, error) {
		row := tx.QueryRowContext(ctx, `SELECT "timestamp", "comment" FROM "transaction" WHERE "transactionID" = ?`, transactionID)

		var transaction ledger.Transaction
		err := row.Scan(&transaction.Timestamp, &transaction.Comment)
		if err != nil {
			return ledger.Transaction{}, err
		}

		rows, err := tx.QueryContext(ctx, `SELECT "accountGroup", "accountID", "debit", "credit" FROM "ledger" WHERE "transactionID" = ?`, transactionID)
		if err != nil {
			return ledger.Transaction{}, err
		}

		for rows.Next() {
			var (
				accountNumber ledger.AccountNumber
				debit         *int
				credit        *int
			)
			err = rows.Scan(&accountNumber.AccountGroup, &accountNumber.AccountID, &debit, &credit)
			if err != nil {
				return ledger.Transaction{}, err
			}

			if debit != nil && credit == nil {
				transaction.RecordsDebit = append(transaction.RecordsDebit, ledger.Record{
					AccountNumber: accountNumber, Value: *debit,
				})
			}
			if debit == nil && credit != nil {
				transaction.RecordsCredit = append(transaction.RecordsCredit, ledger.Record{
					AccountNumber: accountNumber, Value: *credit,
				})
			}
		}
		if err = rows.Err(); err != nil {
			return ledger.Transaction{}, err
		}

		return transaction, nil
	})
}

func (db Database) InsertTransaction(ctx context.Context, transaction ledger.Transaction) (int64, error) {
	return withTransaction(ctx, db.db, func(ctx context.Context, tx *sql.Tx) (int64, error) {
		res, err := tx.ExecContext(ctx, `INSERT INTO "transaction" (
			"timestamp", "comment"
		) VALUES (?, ?)`, transaction.Timestamp, transaction.Comment)
		if err != nil {
			return 0, err
		}

		transactionID, err := res.LastInsertId()
		if err != nil {
			return 0, err
		}

		stmtDebit, err := tx.PrepareContext(ctx, `INSERT INTO "ledger" (
			"transactionID", "accountGroup", "accountID", "debit"
		) VALUES (?, ?, ?, ?)`)
		if err != nil {
			return 0, err
		}

		stmtCredit, err := tx.PrepareContext(ctx, `INSERT INTO "ledger" (
			"transactionID", "accountGroup", "accountID", "credit"
		) VALUES (?, ?, ?, ?)`)
		if err != nil {
			return 0, err
		}

		for _, debit := range transaction.RecordsDebit {
			_, err = stmtDebit.ExecContext(ctx, transactionID, debit.AccountGroup, debit.AccountID, debit.Value)
			if err != nil {
				return 0, err
			}
		}

		for _, credit := range transaction.RecordsCredit {
			_, err = stmtCredit.ExecContext(ctx, transactionID, credit.AccountGroup, credit.AccountID, credit.Value)
			if err != nil {
				return 0, err
			}
		}

		return transactionID, nil
	})
}
