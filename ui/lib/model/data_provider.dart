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
      res.add(Command(it.id, it.name, "Application", it.icon));
    }

    return res;
  }

  Future<List<CommandAction>> getActions(Int64 commandID) async {
    final List<CommandAction> res = [];
    for (final it in await _client.getActions(commandID)) {
      res.add(CommandAction(it.id, it.name));
    }

    return res;
  }
}
