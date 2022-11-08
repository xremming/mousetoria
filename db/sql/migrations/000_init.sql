CREATE TABLE "account" (
	"accountGroup" INTEGER NOT NULL CHECK("accountGroup" BETWEEN 1 AND 5),
	"accountID"	   INTEGER NOT NULL CHECK("accountID"    BETWEEN 0 AND 9999),
	"name"         TEXT NOT NULL UNIQUE,

	UNIQUE("accountGroup", "accountID")
);

CREATE INDEX "account_group_id" ON "account" ("accountGroup", "accountID");

CREATE TABLE "transaction" (
	"transactionID" INTEGER NOT NULL,
	"timestamp"     INTEGER NOT NULL CHECK("timestamp" > 0),
	"comment"       TEXT,

	PRIMARY KEY ("transactionID")
);

CREATE TABLE "ledger" (
	"recordID"      INTEGER NOT NULL,
	"transactionID" INTEGER NOT NULL,
	"accountGroup"  INTEGER NOT NULL,
	"accountID"     INTEGER NOT NULL,
	"debit"         INTEGER CHECK("debit"  IS NULL OR "debit" > 0),
	"credit"        INTEGER CHECK("credit" IS NULL OR "credit" > 0),

	PRIMARY KEY ("recordID"),

	FOREIGN KEY ("transactionID")
		REFERENCES "transaction" ("transactionID"),
	FOREIGN KEY ("accountID", "accountGroup")
		REFERENCES "account" ("accountID", "accountGroup"),

	CONSTRAINT "debit_xor_credit"
		CHECK(("debit" IS     NULL AND "credit" IS NOT NULL)
		OR    ("debit" IS NOT NULL AND "credit" IS     NULL))
);

CREATE INDEX "ledger_transactionID"          ON "ledger" ("transactionID");
CREATE INDEX "ledger_accountGroup_accountID" ON "ledger" ("accountGroup", "accountID");
