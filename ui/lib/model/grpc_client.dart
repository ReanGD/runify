import 'dart:io';

import 'package:grpc/grpc.dart';
import 'package:fixnum/fixnum.dart';
import 'package:runify/pb/runify.pbgrpc.dart';

class GrpcClient {
  late ClientChannel _channel;
  late RunifyClient _client;

  GrpcClient(String address) {
    _channel = ClientChannel(
      InternetAddress(address, type: InternetAddressType.unix),
      options: const ChannelOptions(credentials: ChannelCredentials.insecure()),
    );

    _client = RunifyClient(_channel,
        options: CallOptions(timeout: const Duration(seconds: 30)));
  }

  Future<List<Command>> getRoot() async {
    final timer = Stopwatch()..start();
    final result = await _client.getRoot(Empty());
    final dt = timer.elapsedMicroseconds;
    print("getRoot = $dt mcs");

    return result.data;
  }

  Future<List<Action>> getActions(Int64 commandID) async {
    final timer = Stopwatch()..start();
    final result =
        await _client.getActions(SelectedCommand(commandID: commandID));
    final dt = timer.elapsedMicroseconds;
    print("getActions = $dt mcs");

    return result.data;
  }

  Future<void> execute(Int64 commandID, int actionID) async {
    final timer = Stopwatch()..start();
    await _client
        .execute(SelectedAction(commandID: commandID, actionID: actionID));
    final dt = timer.elapsedMicroseconds;
    print("execute = $dt mcs");

    return;
  }
}
