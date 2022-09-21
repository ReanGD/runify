import 'package:fixnum/fixnum.dart';
import 'package:runify/model/metrics.dart';
import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/screen/general/gen_types.dart';

class GenService {
  final Metrics metrics;
  final RunifyClient grpcClient;

  GenService(this.metrics, this.grpcClient);

  Future<List<Command>> openForm() async {
    final timer = Stopwatch()..start();
    final grpcResult = await grpcClient.getRoot(Empty());
    metrics.grpcCall(timer.elapsedMicroseconds, 'GetRoot');

    final List<Command> result = [];
    for (final it in grpcResult.cards) {
      result.add(Command(it.cardID, it.name, "Application", it.icon));
    }

    return result;
  }

  Future<void> execute(Int64 cardID) async {
    final timer = Stopwatch()..start();
    await grpcClient.executeDefault(SelectedCard(cardID: cardID));
    metrics.grpcCall(timer.elapsedMicroseconds, 'ExecuteDefault');

    return;
  }
}
