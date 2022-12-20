import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/global/context_menu_row.dart';
import 'package:runify/rpc/rpc_context_menu_service.dart';
import 'package:runify/screen/context_menu/cm_screen.dart';

class _ListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<ContextMenuRowFilter>().visibleItems;
  }
}

class CMController {
  final CMService _service;
  final listController = _ListController();

  CMController(this._service);

  Widget build() {
    return ChangeNotifierProvider.value(
      value: _service.filter,
      child: CMScreen(this),
    );
  }

  get filter => _service.filter;

  void onListItemEvent(DataItemEvent event, ContextMenuRow row) {
    if (event == DataItemEvent.onChoice) {
      _service.execute(row.id);
      return;
    }
  }

  void onApplyFilter(String query) {
    _service.setFilter(query);
  }
}
