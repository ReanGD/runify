import 'package:flutter/material.dart';

import 'package:runify/system/logger.dart';
import 'package:runify/system/settings.dart';
import 'package:runify/global/router_api.dart';
import 'package:runify/global/root_list_row.dart';
import 'package:runify/plugin/runify_native.dart';
import 'package:runify/global/context_menu_row.dart';
import 'package:runify/screen/root_list/rl_controller.dart';
import 'package:runify/screen/context_menu/cm_controller.dart';

class _Listener extends WindowListener {
  final RunifyNavigator navigator;

  _Listener(this.navigator);

  @override
  void onTryClose() {
    navigator.hideWindow();
  }

  @override
  void onFocusChange(bool focused) {}

  @override
  void onConfigure(int x, int y, int width, int height) {}
}

class NilController extends Controller {
  @override
  int get formID => -1;

  @override
  Widget build() {
    return Container();
  }

  @override
  void onFormClosed() {}
}

class RouteItem {
  final Controller controller;
  final Route<dynamic> route;

  RouteItem(this.controller, this.route);

  int get formID => controller.formID;

  @override
  bool operator ==(other) =>
      other is RouteItem && other.controller.formID == controller.formID;

  @override
  int get hashCode => controller.formID;
}

class RunifyNavigator {
  final Logger _logger;
  final Settings _settings;
  final RunifyNative _runifyPlugin;
  final NavigatorState _navigator;
  final _routes = <RouteItem>[];
  var _isFirstShow = true;

  RunifyNavigator(
      this._settings, this._runifyPlugin, this._navigator, this._logger) {
    _runifyPlugin.addListener(_Listener(this));
  }

  void _push(Controller ctrl, Route<dynamic> route) {
    final item = RouteItem(ctrl, route);

    if (_routes.contains(item)) {
      _logger.error("Form with ID ${item.formID} already exists");
      return;
    }

    _routes.add(item);
    _navigator.push(route).then((value) {
      if (_routes.remove(item)) {
        ctrl.onFormClosed();
      }
    });
  }

  void pushForm(Controller ctrl) {
    final route = MaterialPageRoute(
      settings: RouteSettings(name: "form", arguments: ctrl),
      builder: (context) => ctrl.build(),
    );

    _push(ctrl, route);
  }

  void pushMenu(Controller ctrl) {
    final route = RawDialogRoute(
      barrierColor: null,
      barrierLabel: "Label",
      settings: RouteSettings(name: "menu", arguments: ctrl),
      pageBuilder: (context, animation, secondaryAnimation) => ctrl.build(),
    );

    _push(ctrl, route);
  }

  bool canPop() {
    if (_navigator.canPop() && _routes.isNotEmpty) {
      return true;
    }

    if (_navigator.canPop()) {
      _logger.error("Form without ID in navigator stack");
      return false;
    }

    if (_routes.isNotEmpty) {
      _logger.error("Form without route in navigator stack");
      return false;
    }

    return false;
  }

  void pop() {
    if (canPop()) {
      final item = _routes.removeLast();
      item.controller.onFormClosed();
      _navigator.pop();
    }
  }

  void popForm(int formID) {
    if (_routes.isEmpty) {
      return;
    }

    if (_routes.last.formID == formID) {
      pop();
      return;
    }

    final item = _routes.firstWhere((item) => item.formID == formID,
        orElse: () => RouteItem(NilController(),
            MaterialPageRoute(builder: (context) => Container())));

    if (item.formID == -1) {
      return;
    }

    _routes.remove(item);
    item.controller.onFormClosed();
    _navigator.removeRoute(item.route);
  }

  Future<void> openRootList(RootListRpcClient client) async {
    final controller = RLController(client);
    await showWindow();
    pushForm(controller);
  }

  Future<void> openContexMenu(ContextMenuRpcClient client) async {
    final controller = CMController(client);
    pushMenu(controller);
  }

  Future<void> back() async {
    if (_routes.length <= 1) {
      await _runifyPlugin.hide(wait: true);
    }
    pop();
  }

  Future<void> showWindow() async {
    if (_isFirstShow) {
      _isFirstShow = false;
      return _runifyPlugin.initialShow(
          _settings.windowOffset, _settings.windowSize);
    } else {
      return _runifyPlugin.show();
    }
  }

  Future<void> hideWindow() async {
    _runifyPlugin.hide(wait: true);
    while (canPop()) {
      pop();
    }
  }

  Future<void> closeWindow() async {
    return _runifyPlugin.close();
  }
}
