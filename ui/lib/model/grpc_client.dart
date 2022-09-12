import 'dart:io';
import 'package:grpc/grpc.dart';
import 'package:fixnum/fixnum.dart';
import 'package:runify/model/logger.dart';
import 'package:runify/router/router.dart';
import 'package:runify/model/metrics.dart';
import 'package:runify/pb/runify.pbgrpc.dart';

class GrpcClient {
  final Logger logger;
  final Metrics metrics;
  final RunifyRouter router;
  late RunifyClient _client;

  GrpcClient(this.logger, this.metrics, this.router, String address) {
    final channel = ClientChannel(
      InternetAddress(address, type: InternetAddressType.unix),
      options: const ChannelOptions(credentials: ChannelCredentials.insecure()),
    );

    _client = RunifyClient(channel);
    waitShowWindow();
  }

  waitShowWindow() async {
    try {
      final stream = _client.waitShowWindow(Empty());
      await for (var _ in stream) {
        router.show();
      }
    } catch (e) {
      // TODO: add policy for reconnect
      logger.log('gRPC WaitShowWindow method ended with error: $e');
    }
  }

  Future<List<CardItem>> getRoot() async {
    final timer = Stopwatch()..start();
    final result = await _client.getRoot(Empty());
    metrics.grpcCall(timer.elapsedMicroseconds, 'GetRoot');

    return result.cards;
  }

  Future<List<ActionItem>> getActions(Int64 cardID) async {
    final timer = Stopwatch()..start();
    final result = await _client.getActions(SelectedCard(cardID: cardID));
    metrics.grpcCall(timer.elapsedMicroseconds, 'GetActions');

    return result.items;
  }

  Future<void> executeDefault(Int64 cardID) async {
    final timer = Stopwatch()..start();
    await _client.executeDefault(SelectedCard(cardID: cardID));
    metrics.grpcCall(timer.elapsedMicroseconds, 'ExecuteDefault');

    return;
  }

  Future<void> execute(Int64 cardID, int actionID) async {
    final timer = Stopwatch()..start();
    await _client.execute(SelectedAction(cardID: cardID, actionID: actionID));
    metrics.grpcCall(timer.elapsedMicroseconds, 'Execute');

    return;
  }
}
