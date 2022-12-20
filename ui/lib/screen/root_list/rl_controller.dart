import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:runify/global/root_list_row.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/rpc/rpc_root_list_service.dart';
import 'package:runify/screen/root_list/rl_screen.dart';

class _ListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<RootListRowFilter>().visibleItems;
  }
}

class RLController {
  final RLService _service;
  final listController = _ListController();

  RLController(this._service);

  Widget build() {
    return ChangeNotifierProvider.value(
      value: _service.filter,
      child: RLScreen(this),
    );
  }

  get filter => _service.filter;

  void onListItemEvent(DataItemEvent event, RootListRow row) {
    if (event == DataItemEvent.onMenu) {
      _service.menuActivate(row.id);
      return;
    }
    if (event == DataItemEvent.onChoice) {
      _service.execute(row.id);
      return;
    }
  }

  void onApplyFilter(String query) {
    _service.setFilter(query);
  }
}
