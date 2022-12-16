import 'package:runify/pb/runify.pbgrpc.dart';

abstract class FormHandler {
  onRootListAddRows(List<RootListRow> rows);
  onRootListChangeRows(List<RootListRow> rows);
  onRootListRemoveRows(List<RootListRowGlobalID> rows);
}
