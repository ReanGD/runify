import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/command.dart';
import 'package:runify/screen/router.dart';
import 'package:runify/model/grpc_client.dart';
import 'package:runify/model/data_filter.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/screen/general/gen_screen.dart';

class _ListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<CommandFilter>().visibleItems;
  }
}

class GenController {
  final GrpcClient service;
  final ScreenRouter router;
  final DataFilter<Command> filter;
  final listController = _ListController();

  GenController(this.service, this.router)
      : filter = DataFilter.future(service.getRoot());

  Widget build() {
    return ChangeNotifierProvider.value(
      value: filter,
      child: GenScreen(this),
    );
  }

  void onListItemEvent(
      BuildContext context, DataItemEvent event, Command command) {
    if (event == DataItemEvent.onMenu) {
      router.openGScreenMenu(context, command.id);
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
    await service.executeDefault(cardID);
    return router.hideWindow();
  }
}
