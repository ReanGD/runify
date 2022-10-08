import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/screen/router.dart';
import 'package:runify/system/data_filter.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/screen/general_menu/menu_type.dart';
import 'package:runify/screen/general_menu/menu_screen.dart';
import 'package:runify/screen/general_menu/menu_service.dart';

class _ListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<CommandActionFilter>().visibleItems;
  }
}

class MenuController {
  final MenuService service;
  final ScreenRouter router;
  final Int64 itemID;
  final DataFilter<CommandAction> filter;
  final listController = _ListController();

  MenuController(this.service, this.router, {required this.itemID})
      : filter = DataFilter.future(service.openForm(itemID));

  Widget build() {
    return ChangeNotifierProvider.value(
      value: filter,
      child: MenuScreen(this),
    );
  }

  void onListItemEvent(DataItemEvent event, CommandAction action) {
    if (event == DataItemEvent.onChoice) {
      _execute(action.id);
      return;
    }
  }

  void onApplyFilter(String query) {
    filter.applyFilter(query);
  }

  Future<void> _execute(int actionID) async {
    await service.execute(itemID, actionID);
    return router.back(forceHide: true);
  }
}
