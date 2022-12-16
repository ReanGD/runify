import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/rpc/rpc_proto_client.dart';

class Logger {
  final ProtoClient _pClient;

  Logger(this._pClient);

  void debug(String message) {
    _pClient.writeLog(LogLevel.DEBUG, message);
  }

  void info(String message) {
    _pClient.writeLog(LogLevel.INFO, message);
  }

  void warn(String message) {
    _pClient.writeLog(LogLevel.WARNING, message);
  }

  void error(String message) {
    _pClient.writeLog(LogLevel.ERROR, message);
  }
}
