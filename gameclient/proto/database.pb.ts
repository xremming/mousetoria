/* eslint-disable */
import {
  CallOptions,
  ChannelCredentials,
  ChannelOptions,
  Client,
  ClientReadableStream,
  ClientUnaryCall,
  handleServerStreamingCall,
  handleUnaryCall,
  makeGenericClientConstructor,
  Metadata,
  ServiceError,
  UntypedServiceImplementation,
} from "@grpc/grpc-js";
import * as Long from "long";
import * as _m0 from "protobufjs/minimal";
import { Transaction } from "./ledger.pb";

export const protobufPackage = "";

export interface InsertTransactionResponse {
  transactionId: number;
}

export interface GetTransactionRequest {
  transactionId: number;
}

export interface GetAllTransactionsRequest {
  limit?: number | undefined;
}

function createBaseInsertTransactionResponse(): InsertTransactionResponse {
  return { transactionId: 0 };
}

export const InsertTransactionResponse = {
  encode(message: InsertTransactionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.transactionId !== 0) {
      writer.uint32(8).int64(message.transactionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InsertTransactionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInsertTransactionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.transactionId = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InsertTransactionResponse {
    return { transactionId: isSet(object.transactionId) ? Number(object.transactionId) : 0 };
  },

  toJSON(message: InsertTransactionResponse): unknown {
    const obj: any = {};
    message.transactionId !== undefined && (obj.transactionId = Math.round(message.transactionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InsertTransactionResponse>, I>>(object: I): InsertTransactionResponse {
    const message = createBaseInsertTransactionResponse();
    message.transactionId = object.transactionId ?? 0;
    return message;
  },
};

function createBaseGetTransactionRequest(): GetTransactionRequest {
  return { transactionId: 0 };
}

export const GetTransactionRequest = {
  encode(message: GetTransactionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.transactionId !== 0) {
      writer.uint32(8).int64(message.transactionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetTransactionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetTransactionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.transactionId = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetTransactionRequest {
    return { transactionId: isSet(object.transactionId) ? Number(object.transactionId) : 0 };
  },

  toJSON(message: GetTransactionRequest): unknown {
    const obj: any = {};
    message.transactionId !== undefined && (obj.transactionId = Math.round(message.transactionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetTransactionRequest>, I>>(object: I): GetTransactionRequest {
    const message = createBaseGetTransactionRequest();
    message.transactionId = object.transactionId ?? 0;
    return message;
  },
};

function createBaseGetAllTransactionsRequest(): GetAllTransactionsRequest {
  return { limit: undefined };
}

export const GetAllTransactionsRequest = {
  encode(message: GetAllTransactionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.limit !== undefined) {
      writer.uint32(8).int64(message.limit);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetAllTransactionsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetAllTransactionsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.limit = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetAllTransactionsRequest {
    return { limit: isSet(object.limit) ? Number(object.limit) : undefined };
  },

  toJSON(message: GetAllTransactionsRequest): unknown {
    const obj: any = {};
    message.limit !== undefined && (obj.limit = Math.round(message.limit));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetAllTransactionsRequest>, I>>(object: I): GetAllTransactionsRequest {
    const message = createBaseGetAllTransactionsRequest();
    message.limit = object.limit ?? undefined;
    return message;
  },
};

export type DatabaseService = typeof DatabaseService;
export const DatabaseService = {
  insertTransaction: {
    path: "/Database/InsertTransaction",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: Transaction) => Buffer.from(Transaction.encode(value).finish()),
    requestDeserialize: (value: Buffer) => Transaction.decode(value),
    responseSerialize: (value: InsertTransactionResponse) =>
      Buffer.from(InsertTransactionResponse.encode(value).finish()),
    responseDeserialize: (value: Buffer) => InsertTransactionResponse.decode(value),
  },
  getTransaction: {
    path: "/Database/GetTransaction",
    requestStream: false,
    responseStream: false,
    requestSerialize: (value: GetTransactionRequest) => Buffer.from(GetTransactionRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => GetTransactionRequest.decode(value),
    responseSerialize: (value: Transaction) => Buffer.from(Transaction.encode(value).finish()),
    responseDeserialize: (value: Buffer) => Transaction.decode(value),
  },
  getAllTransactions: {
    path: "/Database/GetAllTransactions",
    requestStream: false,
    responseStream: true,
    requestSerialize: (value: GetAllTransactionsRequest) =>
      Buffer.from(GetAllTransactionsRequest.encode(value).finish()),
    requestDeserialize: (value: Buffer) => GetAllTransactionsRequest.decode(value),
    responseSerialize: (value: Transaction) => Buffer.from(Transaction.encode(value).finish()),
    responseDeserialize: (value: Buffer) => Transaction.decode(value),
  },
} as const;

export interface DatabaseServer extends UntypedServiceImplementation {
  insertTransaction: handleUnaryCall<Transaction, InsertTransactionResponse>;
  getTransaction: handleUnaryCall<GetTransactionRequest, Transaction>;
  getAllTransactions: handleServerStreamingCall<GetAllTransactionsRequest, Transaction>;
}

export interface DatabaseClient extends Client {
  insertTransaction(
    request: Transaction,
    callback: (error: ServiceError | null, response: InsertTransactionResponse) => void,
  ): ClientUnaryCall;
  insertTransaction(
    request: Transaction,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: InsertTransactionResponse) => void,
  ): ClientUnaryCall;
  insertTransaction(
    request: Transaction,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: InsertTransactionResponse) => void,
  ): ClientUnaryCall;
  getTransaction(
    request: GetTransactionRequest,
    callback: (error: ServiceError | null, response: Transaction) => void,
  ): ClientUnaryCall;
  getTransaction(
    request: GetTransactionRequest,
    metadata: Metadata,
    callback: (error: ServiceError | null, response: Transaction) => void,
  ): ClientUnaryCall;
  getTransaction(
    request: GetTransactionRequest,
    metadata: Metadata,
    options: Partial<CallOptions>,
    callback: (error: ServiceError | null, response: Transaction) => void,
  ): ClientUnaryCall;
  getAllTransactions(
    request: GetAllTransactionsRequest,
    options?: Partial<CallOptions>,
  ): ClientReadableStream<Transaction>;
  getAllTransactions(
    request: GetAllTransactionsRequest,
    metadata?: Metadata,
    options?: Partial<CallOptions>,
  ): ClientReadableStream<Transaction>;
}

export const DatabaseClient = makeGenericClientConstructor(DatabaseService, "Database") as unknown as {
  new (address: string, credentials: ChannelCredentials, options?: Partial<ChannelOptions>): DatabaseClient;
  service: typeof DatabaseService;
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
