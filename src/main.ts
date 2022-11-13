import * as path from "node:path";
import { promisify } from "node:util";

import { app, BrowserWindow, ipcMain } from "electron";

import { TimeOfDay } from "./gen/timestamp.pb";
import { AccountGroup } from "./gen/ledger.pb";
import { DatabaseClient } from "./gen/database.pb";
import { startServer } from "./server";

function createWindow(server: DatabaseClient) {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      preload: path.join(__dirname, "preload.js"),
    },
  });
  ipcMain.handle("new-window", () => {
    new Promise((resolve, reject) => {
      server.insertTransaction(
        {
          timestamp: { day: 0, timeOfDay: TimeOfDay.NIGHT },
          comment: "test",
          recordsDebit: [
            {
              accountNumber: { group: AccountGroup.ASSETS, number: 1 },
              value: 100,
            },
          ],
          recordsCredit: [
            {
              accountNumber: { group: AccountGroup.REVENUE, number: 1 },
              value: 100,
            },
          ],
        },
        (err, resp) => {
          if (err) reject(err);
          else resolve(resp);
        }
      );
    })
      .then(console.log)
      .catch(console.error);
  });
  win.loadFile("../index.html");
}

app.whenReady().then(async () => {
  const [server, stopServer] = await startServer();
  createWindow(server);

  app.on("activate", () => {
    if (BrowserWindow.getAllWindows().length === 0) {
      createWindow(server);
    }
  });

  app.on("window-all-closed", () => {
    if (process.platform !== "darwin") {
      app.quit();
    }
  });

  app.on("before-quit", () => {
    stopServer();
  });
});
