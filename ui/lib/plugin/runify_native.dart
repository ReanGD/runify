import 'dart:async';

import 'package:flutter/services.dart';
import 'package:flutter/foundation.dart';

abstract class WindowListener {
  // Emitted when the window is going to be closed by user, but disable by plugin.
  void onTryClose();

  // Emitted when the window has changed focus state
  void onFocusChange(bool focused);

  // Emitted when the window has changed position or size
  // All values are in pixels
  void onConfigure(int x, int y, int width, int height);
}

class Geomtry {
  final int x;
  final int y;
  final int width;
  final int height;

  Geomtry(this.x, this.y, this.width, this.height);
}

class RunifyNative {
  final MethodChannel _channel = const MethodChannel('runify_native');
  final ObserverList<WindowListener> _listeners =
      ObserverList<WindowListener>();
  bool? _hasFocus;
  Geomtry? _geomtry;

  RunifyNative() {
    _channel.setMethodCallHandler(_methodCallHandler);
  }

  void addListener(WindowListener listener) {
    _listeners.add(listener);
  }

  void removeListener(WindowListener listener) {
    _listeners.remove(listener);
  }

  Future<void> _methodCallHandler(MethodCall call) async {
    for (final WindowListener listener in _listeners) {
      switch (call.method) {
        case "onTryClose":
          listener.onTryClose();
          break;
        case "onFocusChange":
          final hasFocus = call.arguments["hasFocus"] as bool;
          if (_hasFocus == null || _hasFocus != hasFocus) {
            _hasFocus = hasFocus;
            listener.onFocusChange(hasFocus);
          }
          break;
        case "onConfigure":
          final geomtry = Geomtry(
            call.arguments["x"] as int,
            call.arguments["y"] as int,
            call.arguments["width"] as int,
            call.arguments["height"] as int,
          );
          if (_geomtry == null || _geomtry != geomtry) {
            _geomtry = geomtry;
            listener.onConfigure(
                geomtry.x, geomtry.y, geomtry.width, geomtry.height);
          }
          break;
        default:
          break;
      }
    }
  }

  // Initializes the plugin.
  Future<void> initPlugin({bool wait = false}) async {
    await _channel.invokeMethod('initPlugin');
    if (wait) {
      await _waitVisibleState(false);
    }
  }

  // Turns on callbacks for the listener.
  Future<void> _closePlugin() async {
    await _channel.invokeMethod('closePlugin');
  }

  // Whether the window is visible to the user.
  Future<bool> isVisible() async {
    return await _channel.invokeMethod('isVisible');
  }

  // Shows and gives focus to the window with initial position and size.
  //    position = {0, 0} - is center of screen
  //    position and size values are in millimeters.
  // If wait is true, the method will wait until the window is visible.
  Future<void> initialShow(Offset position, Size size,
      {bool wait = false}) async {
    final Map<String, dynamic> arguments = {
      'x': position.dx,
      'y': position.dy,
      'width': size.width,
      'height': size.height,
      'min_width': 480,
      'min_height': 640,
    };

    await _channel.invokeMethod('initialShow', arguments);

    if (wait) {
      await _waitVisibleState(true);
    }
  }

  // Shows and gives focus to the window.
  // If wait is true, the method will wait until the window is visible.
  Future<void> show({bool wait = false}) async {
    await _channel.invokeMethod('show');

    if (wait) {
      await _waitVisibleState(true);
    }
  }

  // Hides the window.
  // If wait is true, the method will wait until the window is hidden.
  Future<void> hide({bool wait = false}) async {
    // Before hiding the window,
    // let's wait until we have received all the KeyUp messages.
    for (;
        RawKeyboard.instance.keysPressed.isNotEmpty;
        await Future.delayed(const Duration(milliseconds: 1))) {}

    await _channel.invokeMethod('hide');
    if (wait) {
      await _waitVisibleState(false);
    }
  }

  Future<void> _waitVisibleState(bool targetState) async {
    var state = !targetState;
    while (state != targetState) {
      await Future.delayed(const Duration(milliseconds: 1));
      state = await isVisible();
    }
  }

  // Close the window.
  Future<void> close() async {
    await _closePlugin();
    await SystemNavigator.pop();
  }

  // Whether window is focused.
  Future<bool> isFocused() async {
    return await _channel.invokeMethod('isFocused');
  }

  // Focuses on the window.
  Future<void> focus() async {
    await _channel.invokeMethod('focus');
  }

  // Sets the opacity of the window.
  // opacity - between 0.0 (fully transparent) and 1.0 (fully opaque)
  Future<void> setOpacity(double opacity) async {
    final Map<String, dynamic> arguments = {
      'opacity': opacity,
    };
    await _channel.invokeMethod('setOpacity', arguments);
  }

  // The bounds of the window.
  // All values are in millimetrs.
  Future<Rect> getGeometry() async {
    final Map<dynamic, dynamic> resultData =
        await _channel.invokeMethod('getGeometry');

    return Rect.fromLTWH(
      resultData['x'],
      resultData['y'],
      resultData['width'],
      resultData['height'],
    );
  }

  // Resizes and moves the window.
  // position = {0, 0} - is center of screen
  // All parameters are in millimeters.
  Future<void> setGeometry(Offset position, Size size) async {
    final Map<String, dynamic> arguments = {
      'x': position.dx,
      'y': position.dy,
      'width': size.width,
      'height': size.height,
    };
    await _channel.invokeMethod('setGeometry', arguments);
  }
}
