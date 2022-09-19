import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:runify/model/grpc_client.dart';
import 'package:runify/plugin/runify_native.dart';
import 'package:runify/screen/general/gen_controller.dart';
import 'package:runify/screen/general_menu/menu_controller.dart';

class ScreenRouter {
  final GrpcClient service;
  final RunifyNative runifyNative;

  ScreenRouter(this.service, this.runifyNative);

  Widget openGScreen() {
    return GenController(service, this).build();
  }

  Future openGScreenMenu(BuildContext context, Int64 itemID) async {
    final controller = MenuController(service, this, itemID: itemID);
    return showDialog(
      context: context,
      barrierColor: null,
      builder: (BuildContext context) {
        return controller.build();
      },
    );
  }

  Future<void> hideWindow() async {
    return runifyNative.hide();
  }

  Future<void> showWindow() async {
    return runifyNative.show();
  }
}
