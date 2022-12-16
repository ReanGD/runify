import 'package:runify/pb/runify.pbgrpc.dart';

abstract class FormHandler {
  void onRootListAddRows(List<RootListRow> rows);
  void onRootListChangeRows(List<RootListRow> rows);
  void onRootListRemoveRows(List<RootListRowGlobalID> rows);
}
