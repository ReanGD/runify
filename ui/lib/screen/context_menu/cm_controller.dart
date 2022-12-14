import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:runify/screen/router.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/global/context_menu_row.dart';
import 'package:runify/screen/context_menu/cm_screen.dart';

class _ListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<ContextMenuRowFilter>().visibleItems;
  }
}

class CMController {
  final ContextMenuRpcClient _client;
  final ScreenRouter router;
  final ContextMenuRowFilter filter;
  final listController = _ListController();

  CMController(this._client, this.router) : filter = _client.filter;

  Widget build() {
    return ChangeNotifierProvider.value(
      value: filter,
      child: CMScreen(this),
    );
  }

  void onListItemEvent(DataItemEvent event, ContextMenuRow row) {
    if (event == DataItemEvent.onChoice) {
      _client.execute(row.id);
      return;
    }
  }

  void onApplyFilter(String query) {
    filter.setFilter(query);
  }
}
