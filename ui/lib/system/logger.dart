import 'dart:async';

import 'package:runify/pb/runify.pbgrpc.dart';

class Logger {
  final StreamController<ServiceMsgUI> _sender;

  Logger(this._sender);

  void _log(LogLevel level, String message) {
    _sender.add(
      ServiceMsgUI(
        writeLog: WriteLog(
          level: level,
          message: message,
        ),
      ),
    );
  }

  void debug(String message) {
    _log(LogLevel.DEBUG, message);
  }

  void info(String message) {
    _log(LogLevel.INFO, message);
  }

  void warn(String message) {
    _log(LogLevel.WARNING, message);
  }

  void error(String message) {
    _log(LogLevel.ERROR, message);
  }
}
