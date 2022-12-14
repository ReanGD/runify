import 'dart:async';

import 'package:runify/system/logger.dart';
import 'package:runify/rpc/rpc_types.dart';
import 'package:runify/rpc/rpc_proto_client.dart';
import 'package:runify/global/root_list_row.dart';
import 'package:runify/pb/runify.pbgrpc.dart' as pb;
import 'package:runify/global/context_menu_row.dart';

class RootListHandler implements FormHandler, RootListRpcClient {
  final ProtoClient _pClient;
  late final RootListRowFilter _filter;

  RootListHandler(this._pClient, List<pb.RootListRow> rows) {
    _filter = RootListRowFilter.value(
        rows
            .map(
              (it) => RootListRow(RootListRowID(it.providerID, it.rowID),
                  it.priority, it.value, "Application", it.icon),
            )
            .toList(),
        rootListRowComparator);
  }

  @override
  RootListRowFilter get filter => _filter;

  @override
  Future<void> onRootListAddRows(List<pb.RootListRow> rows) async {
    // TODO: implement onRootListAddRows
  }

  @override
  Future<void> onRootListChangeRows(List<pb.RootListRow> rows) async {
    // TODO: implement onRootListChangeRows
  }

  @override
  Future<void> onRootListRemoveRows(List<pb.RootListRowGlobalID> rows) async {
    // TODO: implement onRootListRemoveRows
  }

  @override
  Future<void> execute(RootListRowID id) async {
    _pClient.rootListRowActivated(id.providerID, id.rowID);
  }

  @override
  Future<void> menuActivate(RootListRowID id) async {
    _pClient.rootListMenuActivated(id.providerID, id.rowID);
  }
}

class ContextMenuHandler implements FormHandler, ContextMenuRpcClient {
  final Logger _logger;
  final ProtoClient _pClient;
  late final ContextMenuRowFilter _filter;

  ContextMenuHandler(
      this._pClient, this._logger, List<pb.ContextMenuRow> rows) {
    _filter = ContextMenuRowFilter.value(
        rows
            .map(
              (it) => ContextMenuRow(
                it.rowID,
                it.value,
              ),
            )
            .toList(),
        contextMenuRowComparator);
  }

  @override
  ContextMenuRowFilter get filter => _filter;

  @override
  Future<void> onRootListAddRows(List<pb.RootListRow> rows) async {
    _logger.error(
        "Unexpected grpc message 'RootListAddRows' for context menu handler");
  }

  @override
  Future<void> onRootListChangeRows(List<pb.RootListRow> rows) async {
    _logger.error(
        "Unexpected grpc message 'RootListChangeRows' for context menu handler");
  }

  @override
  Future<void> onRootListRemoveRows(List<pb.RootListRowGlobalID> rows) async {
    _logger.error(
        "Unexpected grpc message 'RootListRemoveRows' for context menu handler");
  }

  @override
  Future<void> execute(int id) async {
    _pClient.contextMenuRowActivated(id);
  }
}
