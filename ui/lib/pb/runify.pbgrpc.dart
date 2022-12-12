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
  static final _$connect = $grpc.ClientMethod<$0.UIMessage, $0.SrvMessage>(
      '/runify.Runify/Connect',
      ($0.UIMessage value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.SrvMessage.fromBuffer(value));

  RunifyClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseStream<$0.SrvMessage> connect(
      $async.Stream<$0.UIMessage> request,
      {$grpc.CallOptions? options}) {
    return $createStreamingCall(_$connect, request, options: options);
  }
}

abstract class RunifyServiceBase extends $grpc.Service {
  $core.String get $name => 'runify.Runify';

  RunifyServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.UIMessage, $0.SrvMessage>(
        'Connect',
        connect,
        true,
        true,
        ($core.List<$core.int> value) => $0.UIMessage.fromBuffer(value),
        ($0.SrvMessage value) => value.writeToBuffer()));
  }

  $async.Stream<$0.SrvMessage> connect(
      $grpc.ServiceCall call, $async.Stream<$0.UIMessage> request);
}
