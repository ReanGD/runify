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

class OnBackAction extends Action<OnBackIntent> {
  final ScreenRouter router;

  OnBackAction(this.router);

  @override
  Object? invoke(covariant OnBackIntent intent) {
    router.back();
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
  final _settings = Settings();
  final _runifyPlugin = RunifyNative();
  late final Logger _logger;
  late final Metrics _metrics;
  late final RunifyClient _grpcClient;
  late final NavigatorState _navigator;
  late final ScreenRouterService _service;

  ScreenRouter({super.key}) {
    _metrics = Metrics(_settings.metricsEnabled);
  }

  Future<void> init() async {
    final initFuture =
        _runifyPlugin.initPlugin(_settings.windowOffset, _settings.windowSize);
    _grpcClient = newGrpcClient(_settings);
    _service = ScreenRouterService(_grpcClient, this);
    _service.serviceChannel(this);
    _logger = _service.logger;
    _runifyPlugin.addListener(_Listener(this));
    return initFuture;
  }

  GenController prepareGScreen() {
    final service = GenService(_metrics, _grpcClient);
    return GenController(service, this);
  }

  Future openGScreen(GenController controller) async {
    return _navigator.pushAndRemoveUntil(
      MaterialPageRoute(
        builder: (BuildContext context) {
          return controller.build();
        },
      ),
      (route) => false,
    );
  }

  MenuController prepareGScreenMenu(Int64 itemID) {
    final service = MenuService(_metrics, _grpcClient);
    return MenuController(service, this, itemID: itemID);
  }

  Future openGScreenMenu(MenuController controller) async {
    return _navigator.push(RawDialogRoute(
      barrierColor: null,
      barrierLabel: "Label",
      pageBuilder: (BuildContext a, Animation<double> b, Animation<double> c) {
        return controller.build();
      },
    ));
  }

  Future<void> hideWindow() async {
    return _runifyPlugin.hide();
  }

  Future<void> showWindow() async {
    final controller = prepareGScreen();
    await _runifyPlugin.show();
    return openGScreen(controller);
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

  Future<void> back({bool forceHide = false}) async {
    bool hide = forceHide;
    if (_navigator.canPop()) {
      _navigator.pop();
    } else {
      hide = true;
    }

    if (hide) {
      return _runifyPlugin.hide();
    }
  }

  @override
  Widget build(BuildContext context) {
    _navigator = Navigator.of(context);
    return prepareGScreen().build();
  }
}
