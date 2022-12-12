import 'dart:async';

import 'package:runify/pb/runify.pbgrpc.dart';

abstract class FormHandler {
  Future<void> onRootListAddRows(List<RootListRow> rows);
  Future<void> onRootListChangeRows(List<RootListRow> rows);
  Future<void> onRootListRemoveRows(List<RootListRowGlobalID> rows);
}
