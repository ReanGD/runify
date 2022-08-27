import 'package:fixnum/fixnum.dart';
import 'package:runify/model/command.dart';
import 'package:runify/model/grpc_client.dart';

class DataProvider {
  final _client = GrpcClient("/tmp/runify.socket");

  DataProvider._();

  static final DataProvider instance = DataProvider._();

  Future<List<Command>> getRoot() async {
    final List<Command> res = [];
    for (final it in await _client.getRoot()) {
      res.add(Command(it.cardID, it.name, "Application", it.icon));
    }

    return res;
  }

  Future<List<CommandAction>> getActions(Int64 cardID) async {
    final List<CommandAction> res = [];
    for (final it in await _client.getActions(cardID)) {
      res.add(CommandAction(it.actionID, it.name));
    }

    return res;
  }

  Future<void> executeDefault(Int64 cardID) async {
    await _client.executeDefault(cardID);
  }

  Future<void> execute(Int64 cardID, int actionID) async {
    await _client.execute(cardID, actionID);
  }
}
