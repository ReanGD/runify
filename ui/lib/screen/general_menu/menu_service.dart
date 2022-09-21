import 'package:fixnum/fixnum.dart';
import 'package:runify/model/metrics.dart';
import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/screen/general_menu/menu_type.dart';

class MenuService {
  final Metrics metrics;
  final RunifyClient grpcClient;

  MenuService(this.metrics, this.grpcClient);

  Future<List<CommandAction>> openForm(Int64 cardID) async {
    final timer = Stopwatch()..start();
    final grpcResult =
        await grpcClient.getActions(SelectedCard(cardID: cardID));
    metrics.grpcCall(timer.elapsedMicroseconds, 'GetActions');

    final List<CommandAction> result = [];
    for (final it in grpcResult.items) {
      result.add(CommandAction(it.actionID, it.name));
    }

    return result;
  }

  Future<void> execute(Int64 cardID, int actionID) async {
    final timer = Stopwatch()..start();
    await grpcClient
        .execute(SelectedAction(cardID: cardID, actionID: actionID));
    metrics.grpcCall(timer.elapsedMicroseconds, 'Execute');

    return;
  }
}
