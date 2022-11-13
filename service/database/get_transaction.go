package database

// func (db *Database) GetTransaction(ctx context.Context, transactionID int64) (service.Transaction, error) {
// 	return withTransaction(ctx, db.db, func(ctx context.Context, tx *sql.Tx) (service.Transaction, error) {
// 		logger := zerolog.Ctx(ctx)

// 		row := tx.QueryRowContext(ctx, `SELECT "timestamp", "comment" FROM "transaction" WHERE "transactionID" = ?`, transactionID)

// 		var transaction service.Transaction
// 		err := row.Scan(&transaction.Timestamp, &transaction.Comment)
// 		if err != nil {
// 			return service.Transaction{}, err
// 		}

// 		rows, err := tx.QueryContext(ctx, `SELECT "accountGroup", "accountID", "debit", "credit" FROM "ledger" WHERE "transactionID" = ?`, transactionID)
// 		if err != nil {
// 			return service.Transaction{}, err
// 		}

// 		defer func() {
// 			err = rows.Close()
// 			if err != nil {
// 				logger.Err(err).Msg("failed to close rows")
// 			}
// 		}()

// 		for rows.Next() {
// 			var (
// 				accountNumber service.AccountNumber
// 				debit         *int32
// 				credit        *int32
// 			)
// 			err = rows.Scan(&accountNumber.AccountGroup, &accountNumber.AccountID, &debit, &credit)
// 			if err != nil {
// 				return service.Transaction{}, err
// 			}

// 			if debit != nil && credit == nil {
// 				transaction.RecordsDebit = append(transaction.RecordsDebit, &service.Record{
// 					AccountNumber: &service.AccountNumber{Group: service.AccountGroup()}, Value: *debit,
// 				})
// 			}
// 			if debit == nil && credit != nil {
// 				transaction.RecordsCredit = append(transaction.RecordsCredit, &service.Record{
// 					AccountNumber: accountNumber, Value: *credit,
// 				})
// 			}
// 		}
// 		if err = rows.Err(); err != nil {
// 			return service.Transaction{}, err
// 		}

// 		return transaction, nil
// 	})
// }
