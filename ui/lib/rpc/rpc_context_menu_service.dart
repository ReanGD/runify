import 'package:runify/system/logger.dart';
import 'package:runify/rpc/rpc_types.dart';
import 'package:runify/global/cast_list.dart';
import 'package:runify/rpc/rpc_proto_client.dart';
import 'package:runify/global/context_menu_row.dart';
import 'package:runify/pb/runify.pbgrpc.dart' as pb;

class CMService implements Service {
  final Logger _logger;
  final ProtoClient _pClient;
  late final ContextMenuRowFilter _filter;

  CMService(this._pClient, this._logger, List<pb.ContextMenuRow> rows) {
    _filter = ContextMenuRowFilter();
    _filter.add(CastList(
        rows, (pb.ContextMenuRow row) => ContextMenuRow(row.rowID, row.value)));
    _filter.apply();
  }

  get formID => _pClient.formID;
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

  void setFilter(String value) {
    _filter.setFilter(value);
    _filter.apply();
  }

  void execute(int id) {
    _pClient.contextMenuRowActivated(id);
  }

  void formClosed() {
    // TODO: remove from storage
    _pClient.formClosed();
  }
}
