import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:runify/global/root_list_row.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/screen/root_list/rl_screen.dart';

class _ListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<RootListRowFilter>().visibleItems;
  }
}

class RLController {
  final RootListRpcClient _client;
  final listController = _ListController();

  RLController(this._client);

  Widget build() {
    return ChangeNotifierProvider.value(
      value: _client.filter,
      child: RLScreen(this),
    );
  }

  get filter => _client.filter;

  onListItemEvent(DataItemEvent event, RootListRow row) {
    if (event == DataItemEvent.onMenu) {
      _client.menuActivate(row.id);
      return;
    }
    if (event == DataItemEvent.onChoice) {
      _client.execute(row.id);
      return;
    }
  }

  onApplyFilter(String query) {
    _client.setFilter(query);
  }
}
