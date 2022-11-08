package ledger

import (
	_ "embed"
)

var (
	AssetsUnknown     = AccountNumber{ASSETS, 0}
	AssetsCash        = AccountNumber{ASSETS, 1}
	AssetsReceivables = AccountNumber{ASSETS, 2}
)

var (
	LiabilitiesUnknown = AccountNumber{LIABILITIES, 0}
	LiabilitiesLoans   = AccountNumber{LIABILITIES, 1}
)

var (
	EquityUnknown = AccountNumber{EQUITY, 0}
)

var (
	RevenuesUnknown = AccountNumber{REVENUES, 0}
	RevenuesMinting = AccountNumber{REVENUES, 1}
)

var (
	ExpensesUnknown = AccountNumber{EXPENSES, 0}
	ExpensesTaxes   = AccountNumber{EXPENSES, 1}
)

var accountNames map[AccountNumber]string = map[AccountNumber]string{
	AssetsUnknown:     "Unknown Assets",
	AssetsCash:        "Cash",
	AssetsReceivables: "Receivables",

	LiabilitiesUnknown: "Unknown Liabilities",
	LiabilitiesLoans:   "Loans",

	EquityUnknown: "Unknown Equity",

	RevenuesUnknown: "Unknown Revenues",
	RevenuesMinting: "Minting",

	ExpensesUnknown: "Unknown Expenses",
	ExpensesTaxes:   "Taxes",
}

func AccountNames() map[AccountNumber]string {
	return accountNames
}
