import 'package:flutter/material.dart';

import 'package:runify/system/logger.dart';
import 'package:runify/global/router_api.dart';

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
  final NavigatorState _navigator;
  final _routes = <RouteItem>[];

  RunifyNavigator(this._logger, this._navigator);

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

  bool isLast() {
    return _routes.length <= 1;
  }

  void pop() {
    if (canPop()) {
      final item = _routes.removeLast();
      item.controller.onFormClosed();
      _navigator.pop();
    }
  }

  void popAll() {
    while (canPop()) {
      pop();
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
}
