package ledger

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/xremming/mousetoria/timestamp"
)

type Transaction struct {
	Timestamp     timestamp.Timestamp
	Comment       *string
	RecordsDebit  Records
	RecordsCredit Records
}

type ErrDebitCreditMismatch struct {
	DebitTotal  int
	CreditTotal int
}

func (e ErrDebitCreditMismatch) Difference() int {
	diff := e.DebitTotal - e.CreditTotal

	if diff < 0 {
		return -diff
	}

	return diff
}

func (e ErrDebitCreditMismatch) Error() string {
	return fmt.Sprintf("debit and credit must match: debit was %d and credit was %d", e.DebitTotal, e.CreditTotal)
}

func (t Transaction) Validate() error {
	var err error

	if errTimestamp := t.Timestamp.Validate(); errTimestamp != nil {
		err = multierror.Append(err, errTimestamp)
	}

	if errDebit := t.RecordsDebit.Validate(); errDebit != nil {
		err = multierror.Append(err, fmt.Errorf("debit records: %w", errDebit))
	}

	if errCredit := t.RecordsCredit.Validate(); errCredit != nil {
		err = multierror.Append(err, fmt.Errorf("credit records: %w", errCredit))
	}

	totalDebit, totalCredit := t.RecordsDebit.Total(), t.RecordsCredit.Total()
	if totalDebit != totalCredit {
		err = multierror.Append(err, ErrDebitCreditMismatch{totalDebit, totalCredit})
	}

	return err
}
