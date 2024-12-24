/**
 * @fileoverview gRPC-Web generated client stub for asaycle.motify.v1
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v3.19.1
// source: v1/user.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as v1_user_pb from '../v1/user_pb'; // proto import: "v1/user.proto"


export class UserServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorGetUser = new grpcWeb.MethodDescriptor(
    '/asaycle.motify.v1.UserService/GetUser',
    grpcWeb.MethodType.UNARY,
    v1_user_pb.GetUserRequest,
    v1_user_pb.User,
    (request: v1_user_pb.GetUserRequest) => {
      return request.serializeBinary();
    },
    v1_user_pb.User.deserializeBinary
  );

  getUser(
    request: v1_user_pb.GetUserRequest,
    metadata?: grpcWeb.Metadata | null): Promise<v1_user_pb.User>;

  getUser(
    request: v1_user_pb.GetUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: v1_user_pb.User) => void): grpcWeb.ClientReadableStream<v1_user_pb.User>;

  getUser(
    request: v1_user_pb.GetUserRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: v1_user_pb.User) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/asaycle.motify.v1.UserService/GetUser',
        request,
        metadata || {},
        this.methodDescriptorGetUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/asaycle.motify.v1.UserService/GetUser',
    request,
    metadata || {},
    this.methodDescriptorGetUser);
  }

  methodDescriptorListUsers = new grpcWeb.MethodDescriptor(
    '/asaycle.motify.v1.UserService/ListUsers',
    grpcWeb.MethodType.UNARY,
    v1_user_pb.ListUsersRequest,
    v1_user_pb.ListUsersResponse,
    (request: v1_user_pb.ListUsersRequest) => {
      return request.serializeBinary();
    },
    v1_user_pb.ListUsersResponse.deserializeBinary
  );

  listUsers(
    request: v1_user_pb.ListUsersRequest,
    metadata?: grpcWeb.Metadata | null): Promise<v1_user_pb.ListUsersResponse>;

  listUsers(
    request: v1_user_pb.ListUsersRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: v1_user_pb.ListUsersResponse) => void): grpcWeb.ClientReadableStream<v1_user_pb.ListUsersResponse>;

  listUsers(
    request: v1_user_pb.ListUsersRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: v1_user_pb.ListUsersResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/asaycle.motify.v1.UserService/ListUsers',
        request,
        metadata || {},
        this.methodDescriptorListUsers,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/asaycle.motify.v1.UserService/ListUsers',
    request,
    metadata || {},
    this.methodDescriptorListUsers);
  }

}
