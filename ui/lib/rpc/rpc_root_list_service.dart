import 'package:runify/rpc/rpc_types.dart';
import 'package:runify/system/logger.dart';
import 'package:runify/global/cast_list.dart';
import 'package:runify/rpc/rpc_proto_client.dart';
import 'package:runify/global/root_list_row.dart';
import 'package:runify/pb/runify.pbgrpc.dart' as pb;
import 'package:runify/rpc/rpc_service_storage.dart';

RootListRow castRootListRow(pb.RootListRow row) {
  RootListRowType rowType;
  switch (row.rowType) {
    case pb.RootListRowType.CALCULATOR:
      rowType = RootListRowType.calculator;
      break;
    case pb.RootListRowType.APPLICATION:
      rowType = RootListRowType.application;
      break;
    case pb.RootListRowType.COMMAND:
      rowType = RootListRowType.command;
      break;
    case pb.RootListRowType.LINK:
      rowType = RootListRowType.link;
      break;
    default:
      rowType = RootListRowType.unknown;
  }

  return RootListRow(
    RootListRowID(row.providerID, row.rowID),
    rowType,
    row.priority,
    row.displayName,
    row.searchNames,
    row.icon,
  );
}

class RLService implements Service {
  final Logger _logger;
  final ProtoClient _pClient;
  final ServiceStorage _storage;
  late final RootListRowFilter _filter;

  RLService(
      this._storage, this._pClient, this._logger, List<pb.RootListRow> rows) {
    _filter = RootListRowFilter();
    _filter.add(CastList(rows, castRootListRow));
    _filter.apply();
  }

  get formID => _pClient.formID;
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
  void onFieldCheckResponse(pb.FieldCheckResponse msg) {
    _logger.error(
        "Unexpected grpc message 'FieldCheckResponse' for root list handler");
  }

  void setFilter(String value) {
    _pClient.filterChanged(value);
    _filter.setFilter(value);
  }

  void execute(RootListRowID id) {
    _pClient.rootListRowActivated(id.providerID, id.rowID);
  }

  void menuActivate(RootListRowID id) {
    _pClient.rootListMenuActivated(id.providerID, id.rowID);
  }

  void formClosed() {
    if (_storage.remove(formID)) {
      _pClient.formClosed();
    }
  }
}
