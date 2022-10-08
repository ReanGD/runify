import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:runify/system/logger.dart';
import 'package:runify/system/metrics.dart';
import 'package:runify/system/settings.dart';
import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/system/grpc_client.dart';
import 'package:runify/plugin/runify_native.dart';
import 'package:runify/screen/router_service.dart';
import 'package:runify/screen/general/gen_service.dart';
import 'package:runify/screen/general/gen_controller.dart';
import 'package:runify/screen/general_menu/menu_service.dart';
import 'package:runify/screen/general_menu/menu_controller.dart';

class OnBackIntent extends Intent {
  const OnBackIntent();
}

class OnBackAction extends ContextAction<OnBackIntent> {
  final ScreenRouter router;

  OnBackAction(this.router);

  @override
  Object? invoke(covariant OnBackIntent intent, [BuildContext? context]) {
    if (context != null) {
      router.back(context);
    }

    return null;
  }
}

class _Listener extends WindowListener {
  final ScreenRouter router;

  _Listener(this.router);

  @override
  void onTryClose() {
    router.hideWindow();
  }

  @override
  void onFocusChange(bool focused) {}

  @override
  void onConfigure(int x, int y, int width, int height) {}
}

class ScreenRouter extends StatelessWidget {
  final _logger = Logger();
  final _settings = Settings();
  final _runifyPlugin = RunifyNative();
  late final Metrics _metrics;
  late final RunifyClient _grpcClient;
  late final ScreenRouterService _service;

  ScreenRouter({super.key}) {
    _metrics = Metrics(_settings.metricsEnabled);
  }

  Future<void> init() async {
    final initFuture =
        _runifyPlugin.initPlugin(_settings.windowOffset, _settings.windowSize);
    _grpcClient = newGrpcClient(_settings);
    _service = ScreenRouterService(_logger, _grpcClient, this);
    _service.waitShowWindow(this);
    _runifyPlugin.addListener(_Listener(this));
    return initFuture;
  }

  Widget openGScreen() {
    final service = GenService(_metrics, _grpcClient);
    return GenController(service, this).build();
  }

  Future openGScreenMenu(BuildContext context, Int64 itemID) async {
    final service = MenuService(_metrics, _grpcClient);
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
    return _runifyPlugin.hide();
  }

  Future<void> showWindow() async {
    return _runifyPlugin.show();
  }

  Future<void> closeWindow() async {
    return _runifyPlugin.close();
  }

  Map<Type, Action<Intent>> getActions() {
    return <Type, Action<Intent>>{
      OnBackIntent: OnBackAction(this),
    };
  }

  Map<LogicalKeySet, Intent> getShortcuts() {
    return <LogicalKeySet, Intent>{
      LogicalKeySet(LogicalKeyboardKey.escape): const OnBackIntent(),
    };
  }

  NavigatorState getNavigator(BuildContext context) {
    return Navigator.of(context);
  }

  Future<void> backAndHide(NavigatorState navigator) async {
    if (navigator.canPop()) {
      navigator.pop();
    }
    return _runifyPlugin.hide();
  }

  Future<void> back(BuildContext context) async {
    if (Navigator.of(context).canPop()) {
      Navigator.of(context).pop();
      return;
    }

    return _runifyPlugin.hide();
  }

  @override
  Widget build(BuildContext context) {
    return openGScreen();
  }
}
