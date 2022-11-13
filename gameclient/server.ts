import { spawn } from "node:child_process";
import { mkdtemp, stat } from "node:fs/promises";
import * as path from "node:path";
import * as os from "node:os";

import { z } from "zod";
import * as grpc from "@grpc/grpc-js";

import { DatabaseClient } from "./proto/database.pb";

const ErrorWithCode = z.object({ code: z.string() });

export async function startServer(): Promise<[DatabaseClient, () => void]> {
  const sockDir = await mkdtemp(path.join(os.tmpdir(), "mousetoria"));
  const jsonRpcPath = path.join(sockDir, "grpc.sock");

  const p = spawn(path.join(__dirname, "./gameserver"), [jsonRpcPath]);

  // busy loop until the socket is ready
  while (true) {
    try {
      await stat(jsonRpcPath);
      break;
    } catch (err: unknown) {
      const parsedErr = ErrorWithCode.parse(err);
      console.error(parsedErr);

      if (parsedErr.code !== "ENOENT") {
        throw parsedErr;
      }
    }
  }

  return new Promise((resolve, _reject) => {
    const client = new DatabaseClient(
      "unix:" + jsonRpcPath,
      grpc.credentials.createInsecure()
    );

    resolve([client, () => p.kill()]);
  });
}
