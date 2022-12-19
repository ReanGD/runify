import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:runify/global/router_api.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/global/context_menu_row.dart';
import 'package:runify/screen/context_menu/cm_screen.dart';

class _ListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<ContextMenuRowFilter>().visibleItems;
  }
}

class CMController implements Controller {
  final ContextMenuRpcClient _client;
  final listController = _ListController();

  CMController(this._client);

  @override
  int get formID => _client.formID;

  @override
  Widget build() {
    return ChangeNotifierProvider.value(
      value: _client.filter,
      child: CMScreen(this),
    );
  }

  get filter => _client.filter;

  void onListItemEvent(DataItemEvent event, ContextMenuRow row) {
    if (event == DataItemEvent.onChoice) {
      _client.execute(row.id);
      return;
    }
  }

  void onApplyFilter(String query) {
    _client.setFilter(query);
  }

  @override
  void onFormClosed() {
    _client.formClosed();
  }
}
