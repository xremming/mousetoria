import { StoreProvider } from "easy-peasy";
import React, { useMemo, useState } from "react";
import { store, useStoreActions, useStoreState } from "./state";

import { AccountGroup, Transaction } from "gameclient/proto/ledger.pb";
import "gameclient/types/window.d";

const TransactionList: React.FC<unknown> = () => {
    const transactions = useStoreState((state) => state.transactions);

    return (
        <ul>
            {Object.entries(transactions).map(([id, transaction]) => (
                <li key={id}>
                    {id} {JSON.stringify(transaction)}
                </li>
            ))}
        </ul>
    );
};

const AccountNumberSelector: React.FC<{
    setDebitOrCredit: (debitOrCredit: "debit" | "credit") => void;
    setAccountGroup: (group: number) => void;
    setAccountNumber: (number: number) => void;
    setValue: (value: number) => void;
    remove: () => void;
}> = ({
    setDebitOrCredit,
    setAccountGroup,
    setAccountNumber,
    setValue,
    remove,
}) => {
    const [selectedDebitOrCredit, setSelectedDebitOrCredit] = useState<
        "debit" | "credit"
    >("debit");
    const accountGroups = [
        AccountGroup.ASSETS,
        AccountGroup.EQUITY,
        AccountGroup.EXPENSES,
        AccountGroup.LIABILITIES,
        AccountGroup.REVENUES,
    ];
    const [selectedAccountGroup, setSelectedAccountGroup] = useState(0);
    const [selectedAccountNumber, setSelectedAccountNumber] = useState(0);
    const [selectedValue, setSelectedValue] = useState(0);

    return (
        <>
            <select
                value={selectedDebitOrCredit}
                onChange={(e) => {
                    const type = e.target.value as "debit" | "credit";
                    setSelectedDebitOrCredit(type);
                    setDebitOrCredit(type);
                }}
            >
                <option value="debit">Debit</option>
                <option value="credit">Credit</option>
            </select>
            <select
                value={selectedAccountGroup}
                onChange={(e) => {
                    const accountGroup = parseInt(e.target.value);
                    setSelectedAccountGroup(accountGroup);
                    setAccountGroup(accountGroup);
                }}
            >
                {accountGroups.map((accountGroup) => (
                    <option key={accountGroup} value={accountGroup}>
                        {AccountGroup[accountGroup]}
                    </option>
                ))}
            </select>
            <input
                type="number"
                value={selectedAccountNumber}
                onChange={(e) => {
                    const accountNumber = parseInt(e.target.value);
                    setSelectedAccountNumber(accountNumber);
                    setAccountNumber(accountNumber);
                }}
                min={0}
                max={9999}
                step={1}
            />
            <input
                type="number"
                value={selectedValue}
                onChange={(e) => {
                    const value = parseInt(e.target.value);
                    setSelectedValue(value);
                    setValue(value);
                }}
                min={0}
                step={1}
            />
            <button onClick={remove}>x</button>
        </>
    );
};

const RecordsList: React.FC<{
    records: Records;
    setRecords: React.Dispatch<React.SetStateAction<Records>>;
}> = ({ records, setRecords }) => {
    const addRecord = (type: "debit" | "credit") =>
        setRecords((prev) => [
            ...prev,
            {
                type,
                transaction: {
                    accountNumber: { group: 1, number: 0 },
                    value: 0,
                },
            },
        ]);

    const setRecordType = (index: number) => (type: "debit" | "credit") =>
        setRecords((prev) => {
            const newRecords = [...prev];
            newRecords[index].type = type;
            return newRecords;
        });

    const setRecordAccountNumber = (index: number) => (number: number) =>
        setRecords((prev) => {
            const newRecords = [...prev];
            newRecords[index].transaction.accountNumber!.number = number;
            return newRecords;
        });

    const setRecordAccountGroup = (index: number) => (group: number) =>
        setRecords((prev) => {
            const newRecords = [...prev];
            newRecords[index].transaction.accountNumber!.group = group;
            return newRecords;
        });

    const setRecordValue = (index: number) => (value: number) =>
        setRecords((prev) => {
            const newRecords = [...prev];
            newRecords[index].transaction.value = value;
            return newRecords;
        });

    const removeRecord = (index: number) => () =>
        setRecords((prev) => prev.filter((_, i) => i !== index));

    return (
        <>
            <div>
                {records.map((record, index) => (
                    <div key={index}>
                        <AccountNumberSelector
                            setDebitOrCredit={setRecordType(index)}
                            setAccountNumber={setRecordAccountNumber(index)}
                            setAccountGroup={setRecordAccountGroup(index)}
                            setValue={setRecordValue(index)}
                            remove={removeRecord(index)}
                        />
                    </div>
                ))}
                <button onClick={() => addRecord("debit")}>Add record</button>
            </div>
        </>
    );
};

type Records = {
    type: "debit" | "credit";
    transaction: Transaction["recordsDebit"][0];
}[];

const AddTransaction: React.FC<unknown> = () => {
    const [records, setRecords] = useState<Records>([]);
    const transactions = useMemo(
        () => ({
            timestamp: { day: 0, timeOfDay: 0 },
            comment: "test",
            recordsDebit: records
                .filter((record) => record.type === "debit")
                .map((record) => record.transaction),
            recordsCredit: records
                .filter((record) => record.type === "credit")
                .map((record) => record.transaction),
        }),
        [records]
    );
    const _addTransaction = useStoreActions(
        (actions) => actions.addTransaction
    );
    const addTransaction = () =>
        _addTransaction(transactions).then(() => {
            setRecords([]);
        });

    return (
        <>
            <RecordsList records={records} setRecords={setRecords} />
            <button onClick={addTransaction}>Add transaction</button>
        </>
    );
};

export const App: React.FC<unknown> = () => {
    return (
        <StoreProvider store={store}>
            <div>Transactions</div>
            <TransactionList />
            <AddTransaction />
        </StoreProvider>
    );
};
