import 'dart:io';
import 'package:grpc/grpc.dart';
import 'package:fixnum/fixnum.dart';
import 'package:runify/model/logger.dart';
import 'package:runify/model/command.dart';
import 'package:runify/router/router.dart';
import 'package:runify/model/metrics.dart';
import 'package:runify/model/settings.dart';
import 'package:runify/model/form_data.dart';
import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/model/data_filter.dart';

class GrpcClient {
  final Logger logger;
  final Metrics metrics;
  final Settings settings;
  final RunifyRouter router;
  late RunifyClient _client;

  GrpcClient(this.logger, this.metrics, this.settings, this.router) {
    final channel = ClientChannel(
      InternetAddress(settings.grpcAddress, type: InternetAddressType.unix),
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

  FormData openForm() {
    final dataFilter = DataFilter.future(getRoot());
    return FormData(dataFilter);
  }

  Future<List<Command>> getRoot() async {
    final timer = Stopwatch()..start();
    final grpcResult = await _client.getRoot(Empty());
    metrics.grpcCall(timer.elapsedMicroseconds, 'GetRoot');

    final List<Command> result = [];
    for (final it in grpcResult.cards) {
      result.add(Command(it.cardID, it.name, "Application", it.icon));
    }

    return result;
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
