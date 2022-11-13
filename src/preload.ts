const { BrowserWindow, contextBridge, ipcRenderer } = require("electron");

window.addEventListener("DOMContentLoaded", () => {
  const replaceText = (selector: string, text: string) => {
    const element = document.getElementById(selector);
    if (element) element.innerText = text;
  };

  for (const type of ["chrome", "node", "electron"]) {
    replaceText(`${type}-version`, process.versions[type] ?? "");
  }
});

contextBridge.exposeInMainWorld("api", {
  newWindow: () => ipcRenderer.invoke("new-window"),
});
