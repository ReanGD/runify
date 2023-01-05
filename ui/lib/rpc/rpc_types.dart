import 'package:runify/pb/runify.pbgrpc.dart';

abstract class Service {
  void onRootListAddRows(List<RootListRow> rows);
  void onRootListChangeRows(List<RootListRow> rows);
  void onRootListRemoveRows(List<RootListRowGlobalID> rows);
  void onFieldCheckResponse(FieldCheckResponse msg);
}
