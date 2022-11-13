/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Timestamp } from "./timestamp.pb";

export const protobufPackage = "";

export enum AccountGroup {
  INVALID = 0,
  ASSETS = 1,
  LIABILITIES = 2,
  EQUITY = 3,
  REVENUES = 4,
  EXPENSES = 5,
  UNRECOGNIZED = -1,
}

export function accountGroupFromJSON(object: any): AccountGroup {
  switch (object) {
    case 0:
    case "INVALID":
      return AccountGroup.INVALID;
    case 1:
    case "ASSETS":
      return AccountGroup.ASSETS;
    case 2:
    case "LIABILITIES":
      return AccountGroup.LIABILITIES;
    case 3:
    case "EQUITY":
      return AccountGroup.EQUITY;
    case 4:
    case "REVENUES":
      return AccountGroup.REVENUES;
    case 5:
    case "EXPENSES":
      return AccountGroup.EXPENSES;
    case -1:
    case "UNRECOGNIZED":
    default:
      return AccountGroup.UNRECOGNIZED;
  }
}

export function accountGroupToJSON(object: AccountGroup): string {
  switch (object) {
    case AccountGroup.INVALID:
      return "INVALID";
    case AccountGroup.ASSETS:
      return "ASSETS";
    case AccountGroup.LIABILITIES:
      return "LIABILITIES";
    case AccountGroup.EQUITY:
      return "EQUITY";
    case AccountGroup.REVENUES:
      return "REVENUES";
    case AccountGroup.EXPENSES:
      return "EXPENSES";
    case AccountGroup.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface AccountNumber {
  group: AccountGroup;
  number: number;
}

export interface Record {
  accountNumber: AccountNumber | undefined;
  value: number;
}

export interface Transaction {
  timestamp: Timestamp | undefined;
  comment?: string | undefined;
  recordsDebit: Record[];
  recordsCredit: Record[];
}

function createBaseAccountNumber(): AccountNumber {
  return { group: 0, number: 0 };
}

export const AccountNumber = {
  encode(message: AccountNumber, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.group !== 0) {
      writer.uint32(8).int32(message.group);
    }
    if (message.number !== 0) {
      writer.uint32(16).sint32(message.number);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AccountNumber {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccountNumber();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.group = reader.int32() as any;
          break;
        case 2:
          message.number = reader.sint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccountNumber {
    return {
      group: isSet(object.group) ? accountGroupFromJSON(object.group) : 0,
      number: isSet(object.number) ? Number(object.number) : 0,
    };
  },

  toJSON(message: AccountNumber): unknown {
    const obj: any = {};
    message.group !== undefined && (obj.group = accountGroupToJSON(message.group));
    message.number !== undefined && (obj.number = Math.round(message.number));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AccountNumber>, I>>(object: I): AccountNumber {
    const message = createBaseAccountNumber();
    message.group = object.group ?? 0;
    message.number = object.number ?? 0;
    return message;
  },
};

function createBaseRecord(): Record {
  return { accountNumber: undefined, value: 0 };
}

export const Record = {
  encode(message: Record, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountNumber !== undefined) {
      AccountNumber.encode(message.accountNumber, writer.uint32(10).fork()).ldelim();
    }
    if (message.value !== 0) {
      writer.uint32(16).sint32(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Record {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountNumber = AccountNumber.decode(reader, reader.uint32());
          break;
        case 2:
          message.value = reader.sint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Record {
    return {
      accountNumber: isSet(object.accountNumber) ? AccountNumber.fromJSON(object.accountNumber) : undefined,
      value: isSet(object.value) ? Number(object.value) : 0,
    };
  },

  toJSON(message: Record): unknown {
    const obj: any = {};
    message.accountNumber !== undefined &&
      (obj.accountNumber = message.accountNumber ? AccountNumber.toJSON(message.accountNumber) : undefined);
    message.value !== undefined && (obj.value = Math.round(message.value));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Record>, I>>(object: I): Record {
    const message = createBaseRecord();
    message.accountNumber = (object.accountNumber !== undefined && object.accountNumber !== null)
      ? AccountNumber.fromPartial(object.accountNumber)
      : undefined;
    message.value = object.value ?? 0;
    return message;
  },
};

function createBaseTransaction(): Transaction {
  return { timestamp: undefined, comment: undefined, recordsDebit: [], recordsCredit: [] };
}

export const Transaction = {
  encode(message: Transaction, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.timestamp !== undefined) {
      Timestamp.encode(message.timestamp, writer.uint32(10).fork()).ldelim();
    }
    if (message.comment !== undefined) {
      writer.uint32(18).string(message.comment);
    }
    for (const v of message.recordsDebit) {
      Record.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.recordsCredit) {
      Record.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Transaction {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransaction();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.timestamp = Timestamp.decode(reader, reader.uint32());
          break;
        case 2:
          message.comment = reader.string();
          break;
        case 3:
          message.recordsDebit.push(Record.decode(reader, reader.uint32()));
          break;
        case 4:
          message.recordsCredit.push(Record.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Transaction {
    return {
      timestamp: isSet(object.timestamp) ? Timestamp.fromJSON(object.timestamp) : undefined,
      comment: isSet(object.comment) ? String(object.comment) : undefined,
      recordsDebit: Array.isArray(object?.recordsDebit) ? object.recordsDebit.map((e: any) => Record.fromJSON(e)) : [],
      recordsCredit: Array.isArray(object?.recordsCredit)
        ? object.recordsCredit.map((e: any) => Record.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Transaction): unknown {
    const obj: any = {};
    message.timestamp !== undefined &&
      (obj.timestamp = message.timestamp ? Timestamp.toJSON(message.timestamp) : undefined);
    message.comment !== undefined && (obj.comment = message.comment);
    if (message.recordsDebit) {
      obj.recordsDebit = message.recordsDebit.map((e) => e ? Record.toJSON(e) : undefined);
    } else {
      obj.recordsDebit = [];
    }
    if (message.recordsCredit) {
      obj.recordsCredit = message.recordsCredit.map((e) => e ? Record.toJSON(e) : undefined);
    } else {
      obj.recordsCredit = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Transaction>, I>>(object: I): Transaction {
    const message = createBaseTransaction();
    message.timestamp = (object.timestamp !== undefined && object.timestamp !== null)
      ? Timestamp.fromPartial(object.timestamp)
      : undefined;
    message.comment = object.comment ?? undefined;
    message.recordsDebit = object.recordsDebit?.map((e) => Record.fromPartial(e)) || [];
    message.recordsCredit = object.recordsCredit?.map((e) => Record.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
