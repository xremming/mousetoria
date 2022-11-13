package database

import (
	"gameserver/proto"
)

var accountsAssets = map[int]string{
	0: "Unknown Assets",
	1: "Cash",
	2: "Receivables",
}

var accountsLiabilities = map[int]string{
	0: "Unknown Liabilities",
	1: "Loans",
}

var accountsEquity = map[int]string{
	0: "Unknown Equity",
}

var accountsRevenues = map[int]string{
	0: "Unknown Revenues",
	1: "Minting",
}

var accountsExpenses = map[int]string{
	0: "Unknown Expenses",
	1: "Taxes",
}

var accounts = map[int]map[int]string{
	int(proto.AccountGroup_ASSETS):      accountsAssets,
	int(proto.AccountGroup_LIABILITIES): accountsLiabilities,
	int(proto.AccountGroup_EQUITY):      accountsEquity,
	int(proto.AccountGroup_REVENUES):    accountsRevenues,
	int(proto.AccountGroup_EXPENSES):    accountsExpenses,
}
