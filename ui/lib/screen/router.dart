import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'package:runify/system/metrics.dart';
import 'package:runify/system/settings.dart';
import 'package:runify/global/router_api.dart';
import 'package:runify/rpc/rpc_grpc_client.dart';
import 'package:runify/global/root_list_row.dart';
import 'package:runify/plugin/runify_native.dart';
import 'package:runify/global/context_menu_row.dart';
import 'package:runify/screen/root_list/rl_controller.dart';
import 'package:runify/screen/context_menu/cm_controller.dart';

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
  late final Metrics _metrics;
  late final NavigatorState _navigator;

  ScreenRouter({super.key}) {
    _metrics = Metrics(_settings.metricsEnabled);
  }

  Future<void> init() async {
    final initFuture =
        _runifyPlugin.initPlugin(_settings.windowOffset, _settings.windowSize);
    final grpcClient = GrpcClient(_settings, this);
    grpcClient.connect();
    // _logger = _service.logger;
    _runifyPlugin.addListener(_Listener(this));
    return initFuture;
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

  _showForm(RLController controller) {
    // TODO: need only push
    _navigator.pushAndRemoveUntil(
      MaterialPageRoute(
        builder: (BuildContext context) {
          return controller.build();
        },
      ),
      (route) => false,
    ).then((value) => controller.onFormClosed());
  }

  _showMenu(Controller controller) {
    _navigator
        .push(
          RawDialogRoute(
            barrierColor: null,
            barrierLabel: "Label",
            pageBuilder:
                (BuildContext a, Animation<double> b, Animation<double> c) {
              return controller.build();
            },
          ),
        )
        .then((value) => controller.onFormClosed());
  }

  Future<void> openRootList(RootListRpcClient client) async {
    final controller = RLController(client);
    await _runifyPlugin.show();
    _showForm(controller);
  }

  Future<void> openContexMenu(ContextMenuRpcClient client) async {
    final controller = CMController(client);
    _showMenu(controller);
  }

  Future<void> hideWindow() async {
    // TODO: remove all routes
    return _runifyPlugin.hide();
  }

  Future<void> closeWindow() async {
    return _runifyPlugin.close();
  }

  @override
  Widget build(BuildContext context) {
    _navigator = Navigator.of(context);
    return const Text("Hello");
  }
}
