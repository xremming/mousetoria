package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/xremming/mousetoria/db"
	"github.com/xremming/mousetoria/ledger"
	"github.com/xremming/mousetoria/timestamp"
)

func main() {
	logger := zerolog.New(os.Stderr)
	ctx := logger.WithContext(context.Background())

	db, err := db.NewDatabase(ctx, "file::memory:")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	tx, err := ledger.BuildTx().
		WithTimestamp(1, timestamp.Morning).
		// getting revenue worth 100, and paying 30% taxes from it
		WithDebit(ledger.AssetsCash, 70).
		WithDebit(ledger.ExpensesTaxes, 30).
		WithCredit(ledger.RevenuesUnknown, 100).
		// taxman gets money
		WithDebit(ledger.AssetsCash, 30).
		WithCredit(ledger.RevenuesUnknown, 30).
		Build()

	if err != nil {
		panic(err)
	}

	txID, err := db.InsertTransaction(ctx, tx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v - %v\n", txID, tx)

	tx, err = db.GetTransaction(ctx, txID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v - %v\n", txID, tx)

	t := timestamp.Epoch()
	for i := 0; i < 20_00_000; i++ {
		if t.IsLeapYear() {
			fmt.Printf("%v\t%v\n", t, t.IsLeapYear())
		}
		t = t.Next()
	}
}
