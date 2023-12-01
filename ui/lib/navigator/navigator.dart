import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:runify/system/logger.dart';
import 'package:runify/system/settings.dart';
import 'package:runify/global/shortcuts.dart';
import 'package:runify/plugin/runify_native.dart';
import 'package:runify/rpc/rpc_form_service.dart';
import 'package:runify/screen/form/fm_controller.dart';
import 'package:runify/rpc/rpc_root_list_service.dart';
import 'package:runify/rpc/rpc_context_menu_service.dart';
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

typedef FormClosedFn = void Function();

class RouteItem {
  final int formID;
  final FormClosedFn formClosed;
  final Route<dynamic> route;

  RouteItem(this.formID, this.formClosed, this.route);

  @override
  bool operator ==(other) => other is RouteItem && other.formID == formID;

  @override
  int get hashCode => formID;
}

class RunifyNavigator {
  final Logger _logger;
  final Settings _settings;
  final ShortcutStorage _shortcuts;
  final RunifyNative _runifyPlugin;
  final NavigatorState _navigator;
  final _routes = <RouteItem>[];
  var _isFirstShow = true;

  RunifyNavigator(this._settings, this._shortcuts, this._runifyPlugin,
      this._navigator, this._logger) {
    _runifyPlugin.addListener(_Listener(this));
  }

  void _push(int formID, FormClosedFn formClosed, Route<dynamic> route) {
    final item = RouteItem(formID, formClosed, route);

    if (_routes.contains(item)) {
      _logger.error("Form with ID ${item.formID} already exists");
      return;
    }

    _routes.add(item);
    _navigator.push(route).then((value) {
      if (_routes.remove(item)) {
        formClosed();
      }
    });
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
      item.formClosed();
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

    for (RouteItem item in _routes) {
      if (item.formID == formID) {
        _routes.remove(item);
        item.formClosed();
        _navigator.removeRoute(item.route);
        break;
      }
    }
  }

  Widget shortcutsWrapper(Widget child) {
    return ChangeNotifierProvider.value(
      value: _shortcuts,
      child: child,
    );
  }

  Future<void> openForm(FMService service) async {
    final ctrl = FMController(service);

    final route = MaterialPageRoute(
      builder: (context) => shortcutsWrapper(ctrl.build()),
    );

    _push(service.formID, service.formClosed, route);
  }

  Future<void> openRootList(RLService service) async {
    final ctrl = RLController(service);
    await showWindow();

    final route = MaterialPageRoute(
      builder: (context) => shortcutsWrapper(ctrl.build()),
    );

    _push(service.formID, service.formClosed, route);
  }

  Future<void> openContexMenu(CMService service) async {
    final ctrl = CMController(service);

    final route = RawDialogRoute(
      barrierColor: null,
      barrierLabel: "Label",
      pageBuilder: (context, animation, secondaryAnimation) =>
          shortcutsWrapper(ctrl.build()),
    );

    _push(service.formID, service.formClosed, route);
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
