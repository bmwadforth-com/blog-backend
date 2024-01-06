import * as jspb from 'google-protobuf'



export class QueryRequest extends jspb.Message {
  getQuery(): string;
  setQuery(value: string): QueryRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QueryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: QueryRequest): QueryRequest.AsObject;
  static serializeBinaryToWriter(message: QueryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QueryRequest;
  static deserializeBinaryFromReader(message: QueryRequest, reader: jspb.BinaryReader): QueryRequest;
}

export namespace QueryRequest {
  export type AsObject = {
    query: string,
  }
}

export class QueryResponse extends jspb.Message {
  getOption(): number;
  setOption(value: number): QueryResponse;

  getResponse(): string;
  setResponse(value: string): QueryResponse;

  getError(): string;
  setError(value: string): QueryResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): QueryResponse.AsObject;
  static toObject(includeInstance: boolean, msg: QueryResponse): QueryResponse.AsObject;
  static serializeBinaryToWriter(message: QueryResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): QueryResponse;
  static deserializeBinaryFromReader(message: QueryResponse, reader: jspb.BinaryReader): QueryResponse;
}

export namespace QueryResponse {
  export type AsObject = {
    option: number,
    response: string,
    error: string,
  }
}

