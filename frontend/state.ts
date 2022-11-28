import {
  action,
  createStore,
  Action,
  Thunk,
  thunk,
  createTypedHooks,
} from "easy-peasy";
import { type Transaction } from "gameclient/proto/ledger.pb";

type Store = {
  transactions: Record<number, Transaction>;
  _addTransaction: Action<Store, [number, Transaction]>;
  addTransaction: Thunk<Store, Transaction>;
};

export const store = createStore<Store>({
  transactions: [],
  _addTransaction: action((state, [id, tx]) => {
    state.transactions[id] = tx;
  }),
  addTransaction: thunk(async (actions, payload) => {
    const resp = await window.api.insertTransaction(payload);
    actions._addTransaction([resp.transactionId, payload]);
  }),
});

export const { useStoreActions, useStoreState, useStoreDispatch, useStore } =
  createTypedHooks<Store>();
