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

  Future<List<CardItem>> getRoot() async {
    final timer = Stopwatch()..start();
    final result = await _client.getRoot(Empty());
    final dt = timer.elapsedMicroseconds;
    print("getRoot = $dt mcs");

    return result.cards;
  }

  Future<List<ActionItem>> getActions(Int64 cardID) async {
    final timer = Stopwatch()..start();
    final result = await _client.getActions(SelectedCard(cardID: cardID));
    final dt = timer.elapsedMicroseconds;
    print("getActions = $dt mcs");

    return result.items;
  }

  Future<void> executeDefault(Int64 cardID) async {
    final timer = Stopwatch()..start();
    await _client.executeDefault(SelectedCard(cardID: cardID));
    final dt = timer.elapsedMicroseconds;
    print("executeDefault = $dt mcs");

    return;
  }

  Future<void> execute(Int64 cardID, int actionID) async {
    final timer = Stopwatch()..start();
    await _client.execute(SelectedAction(cardID: cardID, actionID: actionID));
    final dt = timer.elapsedMicroseconds;
    print("execute = $dt mcs");

    return;
  }
}
