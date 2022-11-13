package database

import (
	"context"
	"database/sql"
	"errors"
	"gameserver/proto"

	"github.com/rs/zerolog"
)

const (
	selectTransactionSQL = `SELECT "timestamp", "comment" FROM "transaction" WHERE "transactionID" = ?`
	selectRecordsSQL     = `SELECT "accountGroup", "accountID", "debit", "credit" FROM "ledger" WHERE "transactionID" = ?`
)

func (db Database) GetTransaction(ctx context.Context, transaction *proto.GetTransactionRequest) (*proto.Transaction, error) {
	return withTransaction(ctx, db.db, func(ctx context.Context, tx *sql.Tx) (*proto.Transaction, error) {
		transactionID := transaction.GetTransactionId()

		logger := zerolog.Ctx(ctx)

		row := tx.QueryRowContext(ctx, selectTransactionSQL, transactionID)

		var transaction proto.Transaction
		err := row.Scan(&transaction.Timestamp, &transaction.Comment)
		if err != nil {
			return &proto.Transaction{}, err
		}

		rows, err := tx.QueryContext(ctx, selectRecordsSQL, transactionID)
		if err != nil {
			return &proto.Transaction{}, err
		}

		defer func() {
			err = rows.Close()
			if err != nil {
				logger.Err(err).Msg("failed to close rows")
			}
		}()

		for rows.Next() {
			var (
				accountNumber *proto.AccountNumber
				debit         *int32
				credit        *int32
			)

			err = rows.Scan(accountNumber.Group, accountNumber.Number, &debit, &credit)
			if err != nil {
				return &proto.Transaction{}, err
			}

			if debit != nil && credit == nil {
				transaction.RecordsDebit = append(transaction.RecordsDebit, &proto.Record{
					AccountNumber: accountNumber, Value: *debit,
				})
			} else if debit == nil && credit != nil {
				transaction.RecordsCredit = append(transaction.RecordsCredit, &proto.Record{
					AccountNumber: accountNumber, Value: *credit,
				})
			} else {
				return &proto.Transaction{}, errors.New("invalid record: only one of debit or credit must be set")
			}
		}
		if err = rows.Err(); err != nil {
			return &proto.Transaction{}, err
		}

		return &transaction, nil
	})
}
