const { spawn } = require("child_process");
const { app, BrowserWindow, ipcMain } = require("electron");
const net = require("net");
const path = require("path");

class JsonRPCServer {
  constructor(path) {
    this.cbs = {};
    this.counter = 0;
    this.socket = net.createConnection(path);

    this.socket.on("data", (data) => {
      const parsed = JSON.parse(data.toString());

      if (parsed.result !== null) {
        this.cbs[parsed.id].resolve(parsed.result);
      }
      if (parsed.error !== null) {
        this.cbs[parsed.id].reject(parsed.error);
      }

      delete this.cbs[parsed.id];
    });
  }

  call(method, params) {
    return new Promise((resolve, reject) => {
      const id = this.counter++;
      this.cbs[id] = { resolve, reject };
      this.socket.write(JSON.stringify({ id, method, params }));
    });
  }

  notify(method, params) {
    this.socket.write(JSON.stringify({ method, params }));
  }
}

function startServer() {
  try {
    // const p = spawn("./mousetoria", []);
    const s = new JsonRPCServer("/tmp/mousetoria.sock");

    return {
      insertTx(tx) {
        return s.call("Database.InsertTransaction", [tx]);
      },
      stop() {
        p.kill("SIGTERM");
      },
    };
  } catch (e) {
    console.error(e);
  }
}

const server = startServer();

function createWindow(server) {
  const win = new BrowserWindow({
    width: 800,
    height: 600,
    webPreferences: {
      preload: path.join(__dirname, "preload.js"),
    },
  });
  ipcMain.handle("new-window", () => {
    server
      .insertTx({
        Timestamp: 2,
        Comment: "test",
        RecordsDebit: [
          {
            AccountGroup: 1,
            AccountID: 1,
            Value: 100,
          },
        ],
        RecordsCredit: [
          {
            AccountGroup: 1,
            AccountID: 1,
            Value: 100,
          },
        ],
      })
      .then(console.log)
      .catch(console.error);
  });
  win.loadFile("index.html");
}

app.whenReady().then(() => {
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
    server.stop();
  });
});
