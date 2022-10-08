import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/screen/router.dart';
import 'package:runify/system/data_filter.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/screen/general/gen_type.dart';
import 'package:runify/screen/general/gen_screen.dart';
import 'package:runify/screen/general/gen_service.dart';

class _ListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<CommandFilter>().visibleItems;
  }
}

class GenController {
  final GenService service;
  final ScreenRouter router;
  final DataFilter<Command> filter;
  final listController = _ListController();

  GenController(this.service, this.router)
      : filter = DataFilter.future(service.openForm());

  Widget build() {
    return ChangeNotifierProvider.value(
      value: filter,
      child: GenScreen(this),
    );
  }

  void onListItemEvent(DataItemEvent event, Command command) {
    if (event == DataItemEvent.onMenu) {
      final controller = router.prepareGScreenMenu(command.id);
      router.openGScreenMenu(controller);
      return;
    }
    if (event == DataItemEvent.onChoice) {
      _execute(command.id);
      return;
    }
  }

  void onApplyFilter(String query) {
    filter.applyFilter(query);
  }

  Future<void> _execute(Int64 cardID) async {
    await service.execute(cardID);
    return router.hideWindow();
  }
}
