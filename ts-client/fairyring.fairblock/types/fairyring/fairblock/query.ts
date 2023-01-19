/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { EncryptedTx } from "./encrypted_tx";
import { Params } from "./params";

export const protobufPackage = "fairyring.fairblock";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetEncryptedTxRequest {
  targetHeight: number;
  index: number;
}

export interface QueryGetEncryptedTxResponse {
  encryptedTx: EncryptedTx | undefined;
}

export interface QueryAllEncryptedTxRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEncryptedTxResponse {
  encryptedTx: EncryptedTx[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetEncryptedTxRequest(): QueryGetEncryptedTxRequest {
  return { targetHeight: 0, index: 0 };
}

export const QueryGetEncryptedTxRequest = {
  encode(message: QueryGetEncryptedTxRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.targetHeight !== 0) {
      writer.uint32(8).uint64(message.targetHeight);
    }
    if (message.index !== 0) {
      writer.uint32(16).uint64(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEncryptedTxRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEncryptedTxRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.targetHeight = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.index = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEncryptedTxRequest {
    return {
      targetHeight: isSet(object.targetHeight) ? Number(object.targetHeight) : 0,
      index: isSet(object.index) ? Number(object.index) : 0,
    };
  },

  toJSON(message: QueryGetEncryptedTxRequest): unknown {
    const obj: any = {};
    message.targetHeight !== undefined && (obj.targetHeight = Math.round(message.targetHeight));
    message.index !== undefined && (obj.index = Math.round(message.index));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEncryptedTxRequest>, I>>(object: I): QueryGetEncryptedTxRequest {
    const message = createBaseQueryGetEncryptedTxRequest();
    message.targetHeight = object.targetHeight ?? 0;
    message.index = object.index ?? 0;
    return message;
  },
};

function createBaseQueryGetEncryptedTxResponse(): QueryGetEncryptedTxResponse {
  return { encryptedTx: undefined };
}

export const QueryGetEncryptedTxResponse = {
  encode(message: QueryGetEncryptedTxResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.encryptedTx !== undefined) {
      EncryptedTx.encode(message.encryptedTx, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetEncryptedTxResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetEncryptedTxResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.encryptedTx = EncryptedTx.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEncryptedTxResponse {
    return { encryptedTx: isSet(object.encryptedTx) ? EncryptedTx.fromJSON(object.encryptedTx) : undefined };
  },

  toJSON(message: QueryGetEncryptedTxResponse): unknown {
    const obj: any = {};
    message.encryptedTx !== undefined
      && (obj.encryptedTx = message.encryptedTx ? EncryptedTx.toJSON(message.encryptedTx) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetEncryptedTxResponse>, I>>(object: I): QueryGetEncryptedTxResponse {
    const message = createBaseQueryGetEncryptedTxResponse();
    message.encryptedTx = (object.encryptedTx !== undefined && object.encryptedTx !== null)
      ? EncryptedTx.fromPartial(object.encryptedTx)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEncryptedTxRequest(): QueryAllEncryptedTxRequest {
  return { pagination: undefined };
}

export const QueryAllEncryptedTxRequest = {
  encode(message: QueryAllEncryptedTxRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEncryptedTxRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEncryptedTxRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEncryptedTxRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllEncryptedTxRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEncryptedTxRequest>, I>>(object: I): QueryAllEncryptedTxRequest {
    const message = createBaseQueryAllEncryptedTxRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllEncryptedTxResponse(): QueryAllEncryptedTxResponse {
  return { encryptedTx: [], pagination: undefined };
}

export const QueryAllEncryptedTxResponse = {
  encode(message: QueryAllEncryptedTxResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.encryptedTx) {
      EncryptedTx.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllEncryptedTxResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllEncryptedTxResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.encryptedTx.push(EncryptedTx.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllEncryptedTxResponse {
    return {
      encryptedTx: Array.isArray(object?.encryptedTx)
        ? object.encryptedTx.map((e: any) => EncryptedTx.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllEncryptedTxResponse): unknown {
    const obj: any = {};
    if (message.encryptedTx) {
      obj.encryptedTx = message.encryptedTx.map((e) => e ? EncryptedTx.toJSON(e) : undefined);
    } else {
      obj.encryptedTx = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllEncryptedTxResponse>, I>>(object: I): QueryAllEncryptedTxResponse {
    const message = createBaseQueryAllEncryptedTxResponse();
    message.encryptedTx = object.encryptedTx?.map((e) => EncryptedTx.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a EncryptedTx by index. */
  EncryptedTx(request: QueryGetEncryptedTxRequest): Promise<QueryGetEncryptedTxResponse>;
  /** Queries a list of EncryptedTx items. */
  EncryptedTxAll(request: QueryAllEncryptedTxRequest): Promise<QueryAllEncryptedTxResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.EncryptedTx = this.EncryptedTx.bind(this);
    this.EncryptedTxAll = this.EncryptedTxAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("fairyring.fairblock.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  EncryptedTx(request: QueryGetEncryptedTxRequest): Promise<QueryGetEncryptedTxResponse> {
    const data = QueryGetEncryptedTxRequest.encode(request).finish();
    const promise = this.rpc.request("fairyring.fairblock.Query", "EncryptedTx", data);
    return promise.then((data) => QueryGetEncryptedTxResponse.decode(new _m0.Reader(data)));
  }

  EncryptedTxAll(request: QueryAllEncryptedTxRequest): Promise<QueryAllEncryptedTxResponse> {
    const data = QueryAllEncryptedTxRequest.encode(request).finish();
    const promise = this.rpc.request("fairyring.fairblock.Query", "EncryptedTxAll", data);
    return promise.then((data) => QueryAllEncryptedTxResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
