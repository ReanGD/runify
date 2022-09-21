import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:runify/model/logger.dart';
import 'package:runify/model/metrics.dart';
import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/plugin/runify_native.dart';
import 'package:runify/screen/router_service.dart';
import 'package:runify/screen/general/gen_service.dart';
import 'package:runify/screen/general/gen_controller.dart';
import 'package:runify/screen/general_menu/menu_service.dart';
import 'package:runify/screen/general_menu/menu_controller.dart';

class ScreenRouter {
  final Logger logger;
  final Metrics metrics;
  final RunifyClient grpcClient;
  final RunifyNative runifyPlugin;
  final ScreenRouterService _service;

  ScreenRouter({
    required this.logger,
    required this.metrics,
    required this.grpcClient,
    required this.runifyPlugin,
  }) : _service = ScreenRouterService(logger, grpcClient) {
    _service.waitShowWindow(this);
  }

  Widget openGScreen() {
    final service = GenService(metrics, grpcClient);
    return GenController(service, this).build();
  }

  Future openGScreenMenu(BuildContext context, Int64 itemID) async {
    final service = MenuService(metrics, grpcClient);
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
    return runifyPlugin.hide();
  }

  Future<void> showWindow() async {
    return runifyPlugin.show();
  }
}
