package ledger

import (
	"fmt"

	"github.com/xremming/mousetoria/validate"
)

type AccountGroup int

const (
	ASSETS      AccountGroup = 1
	LIABILITIES AccountGroup = 2
	EQUITY      AccountGroup = 3
	REVENUES    AccountGroup = 4
	EXPENSES    AccountGroup = 5
)

func (a AccountGroup) Validate() error {
	if a < 1 || a > 5 {
		return fmt.Errorf("AccountGroup must be between 1 and 5 but was %d", a)
	}
	return nil
}

type AccountID int

func (a AccountID) String() string {
	return fmt.Sprintf("%04d", a)
}

func (a AccountID) Validate() error {
	if a < 0 || a > 9999 {
		return fmt.Errorf("AccountID must be between 0 and 9999 but was %d", a)
	}
	return nil
}

type AccountNumber struct {
	AccountGroup AccountGroup
	AccountID    AccountID
}

func (an AccountNumber) Name() string {
	return accountNames[an]
}

func (an AccountNumber) String() string {
	return fmt.Sprintf("%v-%v (%v)", an.AccountGroup, an.AccountID, an.Name())
}

func (an AccountNumber) Validate() error {
	return validate.ValidateAll(an.AccountGroup, an.AccountID)
}
