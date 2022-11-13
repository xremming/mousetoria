import { contextBridge, ipcRenderer } from "electron";
import { DatabaseClient } from "./proto/database.pb";
import { startServer } from "./server";

import { promisify } from "node:util";

window.addEventListener("DOMContentLoaded", () => {
  const replaceText = (selector: string, text: string) => {
    const element = document.getElementById(selector);
    if (element) element.innerText = text;
  };

  for (const type of ["chrome", "node", "electron"]) {
    replaceText(`${type}-version`, process.versions[type] ?? "");
  }
});

startServer().then(([client, stopServer]: [DatabaseClient, () => void]) => {
  contextBridge.exposeInMainWorld("api", {
    insertTransaction: promisify(client.insertTransaction).bind(client),
  });
});
