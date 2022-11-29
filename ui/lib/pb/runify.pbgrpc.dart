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
  static final _$serviceChannel =
      $grpc.ClientMethod<$0.ServiceMsgUI, $0.ServiceMsgSrv>(
          '/runify.Runify/ServiceChannel',
          ($0.ServiceMsgUI value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.ServiceMsgSrv.fromBuffer(value));
  static final _$formDataChannel =
      $grpc.ClientMethod<$0.FormDataMsgUI, $0.FormDataMsgSrv>(
          '/runify.Runify/FormDataChannel',
          ($0.FormDataMsgUI value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.FormDataMsgSrv.fromBuffer(value));
  static final _$getRoot = $grpc.ClientMethod<$0.Empty, $0.Form>(
      '/runify.Runify/GetRoot',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Form.fromBuffer(value));
  static final _$getActions = $grpc.ClientMethod<$0.SelectedCard, $0.Actions>(
      '/runify.Runify/GetActions',
      ($0.SelectedCard value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Actions.fromBuffer(value));
  static final _$executeDefault =
      $grpc.ClientMethod<$0.SelectedCard, $0.Result>(
          '/runify.Runify/ExecuteDefault',
          ($0.SelectedCard value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Result.fromBuffer(value));
  static final _$execute = $grpc.ClientMethod<$0.SelectedAction, $0.Result>(
      '/runify.Runify/Execute',
      ($0.SelectedAction value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.Result.fromBuffer(value));

  RunifyClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseStream<$0.ServiceMsgSrv> serviceChannel(
      $async.Stream<$0.ServiceMsgUI> request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$serviceChannel, request, options: options);
  }

  $grpc.ResponseStream<$0.FormDataMsgSrv> formDataChannel(
      $async.Stream<$0.FormDataMsgUI> request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$formDataChannel, request, options: options);
  }

  $grpc.ResponseFuture<$0.Form> getRoot($0.Empty request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getRoot, request, options: options);
  }

  $grpc.ResponseFuture<$0.Actions> getActions($0.SelectedCard request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getActions, request, options: options);
  }

  $grpc.ResponseFuture<$0.Result> executeDefault($0.SelectedCard request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$executeDefault, request, options: options);
  }

  $grpc.ResponseFuture<$0.Result> execute($0.SelectedAction request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$execute, request, options: options);
  }
}

abstract class RunifyServiceBase extends $grpc.Service {
  $core.String get $name => 'runify.Runify';

  RunifyServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.ServiceMsgUI, $0.ServiceMsgSrv>(
        'ServiceChannel',
        serviceChannel,
        true,
        true,
        ($core.List<$core.int> value) => $0.ServiceMsgUI.fromBuffer(value),
        ($0.ServiceMsgSrv value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.FormDataMsgUI, $0.FormDataMsgSrv>(
        'FormDataChannel',
        formDataChannel,
        true,
        true,
        ($core.List<$core.int> value) => $0.FormDataMsgUI.fromBuffer(value),
        ($0.FormDataMsgSrv value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Empty, $0.Form>(
        'GetRoot',
        getRoot_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($0.Form value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.SelectedCard, $0.Actions>(
        'GetActions',
        getActions_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.SelectedCard.fromBuffer(value),
        ($0.Actions value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.SelectedCard, $0.Result>(
        'ExecuteDefault',
        executeDefault_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.SelectedCard.fromBuffer(value),
        ($0.Result value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.SelectedAction, $0.Result>(
        'Execute',
        execute_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.SelectedAction.fromBuffer(value),
        ($0.Result value) => value.writeToBuffer()));
  }

  $async.Future<$0.Form> getRoot_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return getRoot(call, await request);
  }

  $async.Future<$0.Actions> getActions_Pre(
      $grpc.ServiceCall call, $async.Future<$0.SelectedCard> request) async {
    return getActions(call, await request);
  }

  $async.Future<$0.Result> executeDefault_Pre(
      $grpc.ServiceCall call, $async.Future<$0.SelectedCard> request) async {
    return executeDefault(call, await request);
  }

  $async.Future<$0.Result> execute_Pre(
      $grpc.ServiceCall call, $async.Future<$0.SelectedAction> request) async {
    return execute(call, await request);
  }

  $async.Stream<$0.ServiceMsgSrv> serviceChannel(
      $grpc.ServiceCall call, $async.Stream<$0.ServiceMsgUI> request);
  $async.Stream<$0.FormDataMsgSrv> formDataChannel(
      $grpc.ServiceCall call, $async.Stream<$0.FormDataMsgUI> request);
  $async.Future<$0.Form> getRoot($grpc.ServiceCall call, $0.Empty request);
  $async.Future<$0.Actions> getActions(
      $grpc.ServiceCall call, $0.SelectedCard request);
  $async.Future<$0.Result> executeDefault(
      $grpc.ServiceCall call, $0.SelectedCard request);
  $async.Future<$0.Result> execute(
      $grpc.ServiceCall call, $0.SelectedAction request);
}
