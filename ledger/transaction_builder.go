package ledger

import (
	"github.com/xremming/mousetoria/timestamp"
)

type TxBuilder Transaction

func BuildTx() *TxBuilder {
	return &TxBuilder{}
}

func (b *TxBuilder) WithTimestamp(day int, timeOfDay timestamp.TimeOfDay) *TxBuilder {
	b.Timestamp = timestamp.New(day, timeOfDay)
	return b
}

func (b *TxBuilder) WithComment(comment string) *TxBuilder {
	b.Comment = &comment
	return b
}

func (b *TxBuilder) WithDebit(an AccountNumber, debit int) *TxBuilder {
	b.RecordsDebit = append(b.RecordsDebit, Record{an, debit})
	return b
}

func (b *TxBuilder) WithCredit(an AccountNumber, credit int) *TxBuilder {
	b.RecordsCredit = append(b.RecordsCredit, Record{an, credit})
	return b
}

func (b *TxBuilder) Build() (Transaction, error) {
	out := Transaction{b.Timestamp, b.Comment, b.RecordsDebit, b.RecordsCredit}
	err := out.Validate()

	return out, err
}
