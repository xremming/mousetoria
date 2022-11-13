package database

import "gameserver/proto"

// const selectTransactionsSQL = `SELECT "transactionID", "timestamp", "comment" FROM "ledger" AS l NATURAL JOIN "transaction" AS tx`

func (db Database) GetAllTransactions(req *proto.GetAllTransactionsRequest, out proto.Database_GetAllTransactionsServer) error {
	return nil
}
