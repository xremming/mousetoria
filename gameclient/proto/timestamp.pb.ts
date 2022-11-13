/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "";

export enum TimeOfDay {
  /** NIGHT - Night is the time between 00:00 and 05:59 */
  NIGHT = 0,
  /** MORNING - Morning is the time between 06:00 and 11:59 */
  MORNING = 1,
  /** AFTERNOON - Afternoon is the time between 12:00 and 17:59 */
  AFTERNOON = 2,
  /** EVENING - Evening is the time between 18:00 and 23:59 */
  EVENING = 3,
  UNRECOGNIZED = -1,
}

export function timeOfDayFromJSON(object: any): TimeOfDay {
  switch (object) {
    case 0:
    case "NIGHT":
      return TimeOfDay.NIGHT;
    case 1:
    case "MORNING":
      return TimeOfDay.MORNING;
    case 2:
    case "AFTERNOON":
      return TimeOfDay.AFTERNOON;
    case 3:
    case "EVENING":
      return TimeOfDay.EVENING;
    case -1:
    case "UNRECOGNIZED":
    default:
      return TimeOfDay.UNRECOGNIZED;
  }
}

export function timeOfDayToJSON(object: TimeOfDay): string {
  switch (object) {
    case TimeOfDay.NIGHT:
      return "NIGHT";
    case TimeOfDay.MORNING:
      return "MORNING";
    case TimeOfDay.AFTERNOON:
      return "AFTERNOON";
    case TimeOfDay.EVENING:
      return "EVENING";
    case TimeOfDay.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Timestamp {
  day: number;
  timeOfDay: TimeOfDay;
}

function createBaseTimestamp(): Timestamp {
  return { day: 0, timeOfDay: 0 };
}

export const Timestamp = {
  encode(message: Timestamp, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.day !== 0) {
      writer.uint32(8).sint32(message.day);
    }
    if (message.timeOfDay !== 0) {
      writer.uint32(16).int32(message.timeOfDay);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Timestamp {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTimestamp();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.day = reader.sint32();
          break;
        case 2:
          message.timeOfDay = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Timestamp {
    return {
      day: isSet(object.day) ? Number(object.day) : 0,
      timeOfDay: isSet(object.timeOfDay) ? timeOfDayFromJSON(object.timeOfDay) : 0,
    };
  },

  toJSON(message: Timestamp): unknown {
    const obj: any = {};
    message.day !== undefined && (obj.day = Math.round(message.day));
    message.timeOfDay !== undefined && (obj.timeOfDay = timeOfDayToJSON(message.timeOfDay));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Timestamp>, I>>(object: I): Timestamp {
    const message = createBaseTimestamp();
    message.day = object.day ?? 0;
    message.timeOfDay = object.timeOfDay ?? 0;
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
