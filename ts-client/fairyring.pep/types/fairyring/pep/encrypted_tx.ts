/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "fairyring.pep";

export interface EncryptedTx {
  targetHeight: number;
  index: number;
  data: string;
  creator: string;
  chargedGas: Coin | undefined;
  processedAtChainHeight: number;
  expired: boolean;
}

export interface EncryptedTxArray {
  encryptedTx: EncryptedTx[];
}

export interface GeneralEncryptedTx {
  identity: string;
  index: number;
  data: string;
  creator: string;
  chargedGas: Coin | undefined;
}

export interface GeneralEncryptedTxArray {
  encryptedTx: GeneralEncryptedTx[];
}

export interface GenEncTxExecutionQueue {
  creator: string;
  requestId: string;
  identity: string;
  pubkey: string;
  txList: GeneralEncryptedTxArray | undefined;
  aggrKeyshare: string;
}

function createBaseEncryptedTx(): EncryptedTx {
  return {
    targetHeight: 0,
    index: 0,
    data: "",
    creator: "",
    chargedGas: undefined,
    processedAtChainHeight: 0,
    expired: false,
  };
}

export const EncryptedTx = {
  encode(message: EncryptedTx, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.targetHeight !== 0) {
      writer.uint32(8).uint64(message.targetHeight);
    }
    if (message.index !== 0) {
      writer.uint32(16).uint64(message.index);
    }
    if (message.data !== "") {
      writer.uint32(26).string(message.data);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    if (message.chargedGas !== undefined) {
      Coin.encode(message.chargedGas, writer.uint32(42).fork()).ldelim();
    }
    if (message.processedAtChainHeight !== 0) {
      writer.uint32(48).uint64(message.processedAtChainHeight);
    }
    if (message.expired === true) {
      writer.uint32(56).bool(message.expired);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EncryptedTx {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEncryptedTx();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.targetHeight = longToNumber(reader.uint64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.index = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.data = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.chargedGas = Coin.decode(reader, reader.uint32());
          continue;
        case 6:
          if (tag !== 48) {
            break;
          }

          message.processedAtChainHeight = longToNumber(reader.uint64() as Long);
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.expired = reader.bool();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EncryptedTx {
    return {
      targetHeight: isSet(object.targetHeight) ? Number(object.targetHeight) : 0,
      index: isSet(object.index) ? Number(object.index) : 0,
      data: isSet(object.data) ? String(object.data) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      chargedGas: isSet(object.chargedGas) ? Coin.fromJSON(object.chargedGas) : undefined,
      processedAtChainHeight: isSet(object.processedAtChainHeight) ? Number(object.processedAtChainHeight) : 0,
      expired: isSet(object.expired) ? Boolean(object.expired) : false,
    };
  },

  toJSON(message: EncryptedTx): unknown {
    const obj: any = {};
    if (message.targetHeight !== 0) {
      obj.targetHeight = Math.round(message.targetHeight);
    }
    if (message.index !== 0) {
      obj.index = Math.round(message.index);
    }
    if (message.data !== "") {
      obj.data = message.data;
    }
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.chargedGas !== undefined) {
      obj.chargedGas = Coin.toJSON(message.chargedGas);
    }
    if (message.processedAtChainHeight !== 0) {
      obj.processedAtChainHeight = Math.round(message.processedAtChainHeight);
    }
    if (message.expired === true) {
      obj.expired = message.expired;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<EncryptedTx>, I>>(base?: I): EncryptedTx {
    return EncryptedTx.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<EncryptedTx>, I>>(object: I): EncryptedTx {
    const message = createBaseEncryptedTx();
    message.targetHeight = object.targetHeight ?? 0;
    message.index = object.index ?? 0;
    message.data = object.data ?? "";
    message.creator = object.creator ?? "";
    message.chargedGas = (object.chargedGas !== undefined && object.chargedGas !== null)
      ? Coin.fromPartial(object.chargedGas)
      : undefined;
    message.processedAtChainHeight = object.processedAtChainHeight ?? 0;
    message.expired = object.expired ?? false;
    return message;
  },
};

function createBaseEncryptedTxArray(): EncryptedTxArray {
  return { encryptedTx: [] };
}

export const EncryptedTxArray = {
  encode(message: EncryptedTxArray, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.encryptedTx) {
      EncryptedTx.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EncryptedTxArray {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEncryptedTxArray();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.encryptedTx.push(EncryptedTx.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): EncryptedTxArray {
    return {
      encryptedTx: Array.isArray(object?.encryptedTx)
        ? object.encryptedTx.map((e: any) => EncryptedTx.fromJSON(e))
        : [],
    };
  },

  toJSON(message: EncryptedTxArray): unknown {
    const obj: any = {};
    if (message.encryptedTx?.length) {
      obj.encryptedTx = message.encryptedTx.map((e) => EncryptedTx.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<EncryptedTxArray>, I>>(base?: I): EncryptedTxArray {
    return EncryptedTxArray.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<EncryptedTxArray>, I>>(object: I): EncryptedTxArray {
    const message = createBaseEncryptedTxArray();
    message.encryptedTx = object.encryptedTx?.map((e) => EncryptedTx.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGeneralEncryptedTx(): GeneralEncryptedTx {
  return { identity: "", index: 0, data: "", creator: "", chargedGas: undefined };
}

export const GeneralEncryptedTx = {
  encode(message: GeneralEncryptedTx, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.identity !== "") {
      writer.uint32(10).string(message.identity);
    }
    if (message.index !== 0) {
      writer.uint32(16).uint64(message.index);
    }
    if (message.data !== "") {
      writer.uint32(26).string(message.data);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    if (message.chargedGas !== undefined) {
      Coin.encode(message.chargedGas, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GeneralEncryptedTx {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGeneralEncryptedTx();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.identity = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.index = longToNumber(reader.uint64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.data = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.chargedGas = Coin.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GeneralEncryptedTx {
    return {
      identity: isSet(object.identity) ? String(object.identity) : "",
      index: isSet(object.index) ? Number(object.index) : 0,
      data: isSet(object.data) ? String(object.data) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      chargedGas: isSet(object.chargedGas) ? Coin.fromJSON(object.chargedGas) : undefined,
    };
  },

  toJSON(message: GeneralEncryptedTx): unknown {
    const obj: any = {};
    if (message.identity !== "") {
      obj.identity = message.identity;
    }
    if (message.index !== 0) {
      obj.index = Math.round(message.index);
    }
    if (message.data !== "") {
      obj.data = message.data;
    }
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.chargedGas !== undefined) {
      obj.chargedGas = Coin.toJSON(message.chargedGas);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GeneralEncryptedTx>, I>>(base?: I): GeneralEncryptedTx {
    return GeneralEncryptedTx.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GeneralEncryptedTx>, I>>(object: I): GeneralEncryptedTx {
    const message = createBaseGeneralEncryptedTx();
    message.identity = object.identity ?? "";
    message.index = object.index ?? 0;
    message.data = object.data ?? "";
    message.creator = object.creator ?? "";
    message.chargedGas = (object.chargedGas !== undefined && object.chargedGas !== null)
      ? Coin.fromPartial(object.chargedGas)
      : undefined;
    return message;
  },
};

function createBaseGeneralEncryptedTxArray(): GeneralEncryptedTxArray {
  return { encryptedTx: [] };
}

export const GeneralEncryptedTxArray = {
  encode(message: GeneralEncryptedTxArray, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.encryptedTx) {
      GeneralEncryptedTx.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GeneralEncryptedTxArray {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGeneralEncryptedTxArray();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.encryptedTx.push(GeneralEncryptedTx.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GeneralEncryptedTxArray {
    return {
      encryptedTx: Array.isArray(object?.encryptedTx)
        ? object.encryptedTx.map((e: any) => GeneralEncryptedTx.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GeneralEncryptedTxArray): unknown {
    const obj: any = {};
    if (message.encryptedTx?.length) {
      obj.encryptedTx = message.encryptedTx.map((e) => GeneralEncryptedTx.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GeneralEncryptedTxArray>, I>>(base?: I): GeneralEncryptedTxArray {
    return GeneralEncryptedTxArray.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GeneralEncryptedTxArray>, I>>(object: I): GeneralEncryptedTxArray {
    const message = createBaseGeneralEncryptedTxArray();
    message.encryptedTx = object.encryptedTx?.map((e) => GeneralEncryptedTx.fromPartial(e)) || [];
    return message;
  },
};

function createBaseGenEncTxExecutionQueue(): GenEncTxExecutionQueue {
  return { creator: "", requestId: "", identity: "", pubkey: "", txList: undefined, aggrKeyshare: "" };
}

export const GenEncTxExecutionQueue = {
  encode(message: GenEncTxExecutionQueue, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.requestId !== "") {
      writer.uint32(18).string(message.requestId);
    }
    if (message.identity !== "") {
      writer.uint32(26).string(message.identity);
    }
    if (message.pubkey !== "") {
      writer.uint32(34).string(message.pubkey);
    }
    if (message.txList !== undefined) {
      GeneralEncryptedTxArray.encode(message.txList, writer.uint32(42).fork()).ldelim();
    }
    if (message.aggrKeyshare !== "") {
      writer.uint32(50).string(message.aggrKeyshare);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenEncTxExecutionQueue {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenEncTxExecutionQueue();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.creator = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.requestId = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.identity = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.pubkey = reader.string();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.txList = GeneralEncryptedTxArray.decode(reader, reader.uint32());
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.aggrKeyshare = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GenEncTxExecutionQueue {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      requestId: isSet(object.requestId) ? String(object.requestId) : "",
      identity: isSet(object.identity) ? String(object.identity) : "",
      pubkey: isSet(object.pubkey) ? String(object.pubkey) : "",
      txList: isSet(object.txList) ? GeneralEncryptedTxArray.fromJSON(object.txList) : undefined,
      aggrKeyshare: isSet(object.aggrKeyshare) ? String(object.aggrKeyshare) : "",
    };
  },

  toJSON(message: GenEncTxExecutionQueue): unknown {
    const obj: any = {};
    if (message.creator !== "") {
      obj.creator = message.creator;
    }
    if (message.requestId !== "") {
      obj.requestId = message.requestId;
    }
    if (message.identity !== "") {
      obj.identity = message.identity;
    }
    if (message.pubkey !== "") {
      obj.pubkey = message.pubkey;
    }
    if (message.txList !== undefined) {
      obj.txList = GeneralEncryptedTxArray.toJSON(message.txList);
    }
    if (message.aggrKeyshare !== "") {
      obj.aggrKeyshare = message.aggrKeyshare;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GenEncTxExecutionQueue>, I>>(base?: I): GenEncTxExecutionQueue {
    return GenEncTxExecutionQueue.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GenEncTxExecutionQueue>, I>>(object: I): GenEncTxExecutionQueue {
    const message = createBaseGenEncTxExecutionQueue();
    message.creator = object.creator ?? "";
    message.requestId = object.requestId ?? "";
    message.identity = object.identity ?? "";
    message.pubkey = object.pubkey ?? "";
    message.txList = (object.txList !== undefined && object.txList !== null)
      ? GeneralEncryptedTxArray.fromPartial(object.txList)
      : undefined;
    message.aggrKeyshare = object.aggrKeyshare ?? "";
    return message;
  },
};

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const tsProtoGlobalThis: any = (() => {
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
    throw new tsProtoGlobalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
