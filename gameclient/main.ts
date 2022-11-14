import * as path from "node:path";

import { app, BrowserWindow, ipcMain } from "electron";

function createWindow() {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      // based on docs, the nodeIntegration should not be needed in this case
      // but the code will not work without it nor can I see any reason why it should not be enabled
      // https://github.com/electron/electron/issues/36330
      nodeIntegration: true,
      preload: path.join(__dirname, "preload.js"),
    },
  });

  win.loadFile("../index.html");
  win.on("ready-to-show", () => win.show());
}

app.whenReady().then(() => {
  createWindow();

  app.on("activate", () => {
    if (BrowserWindow.getAllWindows().length === 0) {
      createWindow();
    }
  });

  app.on("window-all-closed", () => {
    if (process.platform !== "darwin") {
      app.quit();
    }
  });
});
