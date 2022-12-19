import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'package:runify/system/logger.dart';
import 'package:runify/system/metrics.dart';
import 'package:runify/system/settings.dart';
import 'package:runify/navigator/navigator.dart';
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
  late final Logger _logger;
  late final RunifyNavigator _navigator;
  late final GrpcClient _grpcClient;
  var _isFirstShow = true;

  ScreenRouter({super.key}) {
    _metrics = Metrics(_settings.metricsEnabled);
  }

  Future<void> init() async {
    final pluginFuture = _runifyPlugin.initPlugin();
    _runifyPlugin.addListener(_Listener(this));
    _grpcClient = GrpcClient(_settings, this);
    _logger = _grpcClient.logger;

    return pluginFuture;
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

  Future<void> back() async {
    if (_navigator.isLast()) {
      await _runifyPlugin.hide(wait: true);
    }
    _navigator.pop();
  }

  Future<void> openRootList(RootListRpcClient client) async {
    final controller = RLController(client);
    if (_isFirstShow) {
      _isFirstShow = false;
      await _runifyPlugin.initialShow(
          _settings.windowOffset, _settings.windowSize);
    } else {
      await _runifyPlugin.show();
    }
    _navigator.pushForm(controller);
  }

  Future<void> openContexMenu(ContextMenuRpcClient client) async {
    final controller = CMController(client);
    _navigator.pushMenu(controller);
  }

  Future<void> hideWindow() async {
    _runifyPlugin.hide(wait: true);
    _navigator.popAll();
  }

  Future<void> closeWindow() async {
    return _runifyPlugin.close();
  }

  @override
  Widget build(BuildContext context) {
    _navigator = RunifyNavigator(_logger, Navigator.of(context));
    _grpcClient.connect();
    final theme = Theme.of(context);
    return Container(
      color: theme.backgroundColor,
    );
  }
}
