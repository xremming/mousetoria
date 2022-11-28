CREATE TABLE "account" (
	"accountGroup" INTEGER NOT NULL CHECK("accountGroup" BETWEEN 1 AND 5),
	"accountID"	   INTEGER NOT NULL CHECK("accountID"    BETWEEN 0 AND 9999),
	"name"         TEXT NOT NULL UNIQUE,

	UNIQUE("accountGroup", "accountID")
);

CREATE INDEX "account_group_id"
	ON "account" ("accountGroup", "accountID");

CREATE TABLE "entity" (
	"entityID" INTEGER NOT NULL,
	"type"     INTEGER NOT NULL,
	"born"     INTEGER NOT NULL CHECK("born" >= 0),
	"name"     TEXT,

	PRIMARY KEY ("entityID")
);

CREATE TABLE "transaction" (
	"transactionID" INTEGER NOT NULL,
	"timestamp"     INTEGER NOT NULL CHECK("timestamp" >= 0),
	"comment"       TEXT,

	PRIMARY KEY ("transactionID")
);

CREATE TABLE "ledger" (
	"recordID"      INTEGER NOT NULL,
	"transactionID" INTEGER NOT NULL,
	"accountGroup"  INTEGER NOT NULL,
	"accountID"     INTEGER NOT NULL,
	"entityID"      INTEGER, -- NOT NULL,
	"debit"         INTEGER CHECK("debit"  IS NULL OR "debit"  > 0),
	"credit"        INTEGER CHECK("credit" IS NULL OR "credit" > 0),

	PRIMARY KEY ("recordID"),

	FOREIGN KEY ("transactionID")
		REFERENCES "transaction" ("transactionID"),
	FOREIGN KEY ("accountID", "accountGroup")
		REFERENCES "account" ("accountID", "accountGroup"),
	FOREIGN KEY ("entityID")
		REFERENCES "entity" ("entityID"),

	CONSTRAINT "debit_xor_credit"
		CHECK(("debit" IS     NULL AND "credit" IS NOT NULL)
		OR    ("debit" IS NOT NULL AND "credit" IS     NULL))
);

CREATE INDEX "ledger_transactionID"
	ON "ledger" ("transactionID");
CREATE INDEX "ledger_accountGroup_accountID"
	ON "ledger" ("accountGroup", "accountID");

CREATE TABLE "node" (
    "nodeId" INTEGER NOT NULL,
    PRIMARY KEY ("nodeId")
);

CREATE TABLE "edge" (
    "from" INTEGER NOT NULL,
    "to"   INTEGER NOT NULL,
    FOREIGN KEY ("from") REFERENCES "node" ("nodeId"),
    FOREIGN KEY ("to")   REFERENCES "node" ("nodeId"),
    CONSTRAINT "from_to_unique"
        UNIQUE ("from", "to"),
    CONSTRAINT "from_to_not_equal"
        CHECK ("from" != "to"),
    CONSTRAINT "from_to_ordered"
        CHECK ("from" < "to")
);
