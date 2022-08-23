///
//  Generated code. Do not modify.
//  source: runify.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'runify.pb.dart' as $0;
export 'runify.pb.dart';

class RunifyClient extends $grpc.Client {
  static final _$getRoot = $grpc.ClientMethod<$0.Empty, $0.Commands>(
      '/runify.Runify/GetRoot',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Commands.fromBuffer(value));
  static final _$getActions = $grpc.ClientMethod<$0.CommandID, $0.Actions>(
      '/runify.Runify/GetActions',
      ($0.CommandID value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Actions.fromBuffer(value));
  static final _$execute = $grpc.ClientMethod<$0.ActionID, $0.Result>(
      '/runify.Runify/Execute',
      ($0.ActionID value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Result.fromBuffer(value));

  RunifyClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.Commands> getRoot($0.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getRoot, request, options: options);
  }

  $grpc.ResponseFuture<$0.Actions> getActions($0.CommandID request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getActions, request, options: options);
  }

  $grpc.ResponseFuture<$0.Result> execute($0.ActionID request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$execute, request, options: options);
  }
}

abstract class RunifyServiceBase extends $grpc.Service {
  $core.String get $name => 'runify.Runify';

  RunifyServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $0.Commands>(
        'GetRoot',
        getRoot_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($0.Commands value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.CommandID, $0.Actions>(
        'GetActions',
        getActions_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CommandID.fromBuffer(value),
        ($0.Actions value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.ActionID, $0.Result>(
        'Execute',
        execute_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.ActionID.fromBuffer(value),
        ($0.Result value) => value.writeToBuffer()));
  }

  $async.Future<$0.Commands> getRoot_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return getRoot(call, await request);
  }

  $async.Future<$0.Actions> getActions_Pre(
      $grpc.ServiceCall call, $async.Future<$0.CommandID> request) async {
    return getActions(call, await request);
  }

  $async.Future<$0.Result> execute_Pre(
      $grpc.ServiceCall call, $async.Future<$0.ActionID> request) async {
    return execute(call, await request);
  }

  $async.Future<$0.Commands> getRoot($grpc.ServiceCall call, $0.Empty request);
  $async.Future<$0.Actions> getActions(
      $grpc.ServiceCall call, $0.CommandID request);
  $async.Future<$0.Result> execute($grpc.ServiceCall call, $0.ActionID request);
}
