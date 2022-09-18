import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:runify/form/meta/meta_model.dart';
import 'package:runify/model/command.dart';
import 'package:runify/model/grpc_client.dart';
import 'package:runify/model/data_filter.dart';
import 'package:runify/plugin/runify_native.dart';
import 'package:runify/form/meta/meta_view.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/form/meta/meta_view_menu.dart';

class MetaController {
  final GrpcClient service;
  final RunifyNative runifyNative;
  Int64 _cardID = Int64(0);

  MetaController(this.service, this.runifyNative);

  Widget openForm() {
    final model = MetaModel(service.getRoot());
    return MetaView(model, this);
  }

  Future openMenu(BuildContext context, Command command) {
    return showDialog(
      context: context,
      barrierColor: null,
      builder: (BuildContext context) {
        _cardID = command.id;
        final model = MetaMenuModel(service.getActions(_cardID));
        return MetaViewMenu(model, this);
      },
    );
  }

  void onFormListItemEvent(
      BuildContext context, DataItemEvent event, Command command) {
    if (event == DataItemEvent.onMenu) {
      openMenu(context, command);
      return;
    }
    if (event == DataItemEvent.onChoice) {
      _executeDefault(command.id);
      return;
    }
  }

  void onMenuListItemEvent(
      BuildContext context, DataItemEvent event, CommandAction action) {
    if (event == DataItemEvent.onChoice) {
      _execute(_cardID, action.id);
      return;
    }
  }

  Future<void> _executeDefault(Int64 cardID) async {
    await service.executeDefault(cardID);
    return runifyNative.hide();
  }

  Future<void> _execute(Int64 cardID, int actionID) async {
    await service.execute(cardID, actionID);
    return runifyNative.hide();
  }
}
