import { contextBridge } from "electron";
import { DatabaseClient, InsertTransactionResponse } from "./proto/database.pb";
import { startServer } from "./server";

import { promisify } from "node:util";
import { Transaction } from "./proto/ledger.pb";

window.addEventListener("DOMContentLoaded", () => {
  const replaceText = (selector: string, text: string) => {
    const element = document.getElementById(selector);
    if (element) element.innerText = text;
  };

  for (const type of ["chrome", "node", "electron"]) {
    replaceText(`${type}-version`, process.versions[type] ?? "");
  }
});

const createApi = (client: DatabaseClient) => ({
  insertTransaction: promisify<Transaction, InsertTransactionResponse>(
    client.insertTransaction
  ).bind(client),
});

export type Api = ReturnType<typeof createApi>;

startServer().then(([client, stopServer]: [DatabaseClient, () => void]) => {
  contextBridge.exposeInMainWorld("api", createApi(client));

  window.addEventListener("beforeunload", () => {
    stopServer();
  });
});
