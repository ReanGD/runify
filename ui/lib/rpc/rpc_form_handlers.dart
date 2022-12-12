import 'dart:async';

import 'package:runify/system/logger.dart';
import 'package:runify/rpc/rpc_types.dart';
import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/system/data_filter.dart';
import 'package:runify/rpc/rpc_proto_client.dart';
import 'package:runify/screen/general/gen_type.dart';
import 'package:runify/screen/general_menu/menu_type.dart';

class RootListHandler implements FormHandler, RootListRpcClient {
  final ProtoClient _pClient;
  late final DataFilter<Command> _filter;

  RootListHandler(this._pClient, List<RootListRow> rows) {
    _filter = DataFilter<Command>.value(rows
        .map(
          (it) => Command(it.providerID, it.rowID, it.priority, it.value,
              "Application", it.icon),
        )
        .toList());
  }

  @override
  DataFilter<Command> get filter => _filter;

  @override
  Future<void> onRootListAddRows(List<RootListRow> rows) async {
    // TODO: implement onRootListAddRows
  }

  @override
  Future<void> onRootListChangeRows(List<RootListRow> rows) async {
    // TODO: implement onRootListChangeRows
  }

  @override
  Future<void> onRootListRemoveRows(List<RootListRowGlobalID> rows) async {
    // TODO: implement onRootListRemoveRows
  }

  @override
  Future<void> execute(Command cmd) async {
    _pClient.rootListRowActivated(cmd.providerID, cmd.rowID);
  }

  @override
  Future<void> menuActivate(Command cmd) async {
    _pClient.rootListMenuActivated(cmd.providerID, cmd.rowID);
  }
}

class ContextMenuHandler implements FormHandler, ContextMenuRpcClient {
  final Logger _logger;
  final ProtoClient _pClient;
  late final DataFilter<CommandAction> _filter;

  ContextMenuHandler(this._pClient, this._logger, List<ContextMenuRow> rows) {
    _filter = DataFilter<CommandAction>.value(rows
        .map(
          (it) => CommandAction(
            it.rowID,
            it.value,
          ),
        )
        .toList());
  }

  @override
  DataFilter<CommandAction> get filter => _filter;

  @override
  Future<void> onRootListAddRows(List<RootListRow> rows) async {
    _logger.error(
        "Unexpected grpc message 'RootListAddRows' for context menu handler");
  }

  @override
  Future<void> onRootListChangeRows(List<RootListRow> rows) async {
    _logger.error(
        "Unexpected grpc message 'RootListChangeRows' for context menu handler");
  }

  @override
  Future<void> onRootListRemoveRows(List<RootListRowGlobalID> rows) async {
    _logger.error(
        "Unexpected grpc message 'RootListRemoveRows' for context menu handler");
  }

  @override
  Future<void> execute(CommandAction cmd) async {
    _pClient.contextMenuRowActivated(cmd.id);
  }
}
