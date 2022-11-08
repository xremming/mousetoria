package ledger

import (
	"errors"

	"github.com/hashicorp/go-multierror"
)

var (
	ErrRecordNotPositive = errors.New("record must be positive")
	ErrZeroRecords       = errors.New("list of records must not be empty")
)

// Record represents a single debit or credit record in a transaction.
type Record struct {
	AccountNumber
	Value int
}

func (r Record) Validate() error {
	err := r.AccountNumber.Validate()

	if r.Value <= 0 {
		err = multierror.Append(err, ErrRecordNotPositive)
	}

	return err
}

type Records []Record

func (rs Records) Total() int {
	var out int

	for _, r := range rs {
		out += r.Value
	}

	return out
}

func (rs Records) Validate() error {
	if len(rs) == 0 {
		return ErrZeroRecords
	}

	var err error

	for _, r := range rs {
		if errR := r.Validate(); errR != nil {
			err = multierror.Append(err, errR)
		}
	}

	return err
}
