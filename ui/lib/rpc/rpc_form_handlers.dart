import 'package:runify/system/logger.dart';
import 'package:runify/rpc/rpc_types.dart';
import 'package:runify/global/cast_list.dart';
import 'package:runify/rpc/rpc_proto_client.dart';
import 'package:runify/global/root_list_row.dart';
import 'package:runify/pb/runify.pbgrpc.dart' as pb;
import 'package:runify/global/context_menu_row.dart';

RootListRow castRootListRow(pb.RootListRow row) {
  return RootListRow(RootListRowID(row.providerID, row.rowID), row.priority,
      row.value, "Application", row.icon);
}

class RootListHandler implements FormHandler, RootListRpcClient {
  final ProtoClient _pClient;
  late final RootListRowFilter _filter;

  RootListHandler(this._pClient, List<pb.RootListRow> rows) {
    _filter = RootListRowFilter();
    _filter.add(CastList(rows, castRootListRow));
    _filter.apply();
  }

  @override
  RootListRowFilter get filter => _filter;

  @override
  void onRootListAddRows(List<pb.RootListRow> rows) {
    _filter.add(CastList(rows, castRootListRow));
    _filter.sort(rootListRowComparator);
    _filter.apply();
  }

  @override
  void onRootListChangeRows(List<pb.RootListRow> rows) {
    final items = CastList(rows, castRootListRow).toList(growable: false);
    final keys = items.map((row) => row.id).toSet();

    _filter.remove(keys);
    _filter.add(items);
    _filter.sort(rootListRowComparator);
    _filter.apply();
  }

  @override
  void onRootListRemoveRows(List<pb.RootListRowGlobalID> rows) {
    final keys = CastList(
        rows,
        (pb.RootListRowGlobalID row) =>
            RootListRowID(row.providerID, row.rowID)).toSet();
    _filter.remove(keys);
    _filter.apply();
  }

  @override
  void setFilter(String value) {
    _pClient.filterChanged(value);
    _filter.setFilter(value);
    _filter.apply();
  }

  @override
  void execute(RootListRowID id) {
    _pClient.rootListRowActivated(id.providerID, id.rowID);
  }

  @override
  void menuActivate(RootListRowID id) {
    _pClient.rootListMenuActivated(id.providerID, id.rowID);
  }

  @override
  void formClosed() {
    _pClient.formClosed();
  }
}

class ContextMenuHandler implements FormHandler, ContextMenuRpcClient {
  final Logger _logger;
  final ProtoClient _pClient;
  late final ContextMenuRowFilter _filter;

  ContextMenuHandler(
      this._pClient, this._logger, List<pb.ContextMenuRow> rows) {
    _filter = ContextMenuRowFilter();
    _filter.add(CastList(
        rows, (pb.ContextMenuRow row) => ContextMenuRow(row.rowID, row.value)));
    _filter.apply();
  }

  @override
  ContextMenuRowFilter get filter => _filter;

  @override
  void onRootListAddRows(List<pb.RootListRow> rows) {
    _logger.error(
        "Unexpected grpc message 'RootListAddRows' for context menu handler");
  }

  @override
  void onRootListChangeRows(List<pb.RootListRow> rows) {
    _logger.error(
        "Unexpected grpc message 'RootListChangeRows' for context menu handler");
  }

  @override
  void onRootListRemoveRows(List<pb.RootListRowGlobalID> rows) {
    _logger.error(
        "Unexpected grpc message 'RootListRemoveRows' for context menu handler");
  }

  @override
  void setFilter(String value) {
    _filter.setFilter(value);
    _filter.apply();
  }

  @override
  void execute(int id) {
    _pClient.contextMenuRowActivated(id);
  }

  @override
  void formClosed() {
    _pClient.formClosed();
  }
}
