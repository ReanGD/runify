import 'dart:io';
import 'dart:async';

import 'package:grpc/grpc.dart';

import 'package:runify/system/logger.dart';
import 'package:runify/system/settings.dart';
import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/navigator/navigator.dart';
import 'package:runify/rpc/rpc_proto_client.dart';
import 'package:runify/rpc/rpc_service_storage.dart';

class GrpcClient {
  final _outCh = StreamController<UIMessage>();
  late final Logger _logger;
  late final RunifyClient _grpcClient;

  GrpcClient(Settings settings) {
    final channel = ClientChannel(
      InternetAddress(
        settings.grpcAddress,
        type: InternetAddressType.unix,
      ),
      options: const ChannelOptions(
        credentials: ChannelCredentials.insecure(),
      ),
    );

    _grpcClient = RunifyClient(channel);
    _logger = Logger(ProtoClient(0, _outCh));
  }

  Logger get logger => _logger;

  Future<void> start(RunifyNavigator navigator) async {
    try {
      final storage = ServiceStorage(_outCh, navigator, _logger);
      final msg = await _callConnect(storage);
      // ignore: avoid_print
      print("gRPC call 'connect' ended. $msg. Stop runify UI.");
    } catch (e) {
      final msg = e.toString();
      // ignore: avoid_print
      print("gRPC call 'connect' ended with exception: $msg. Stop runify UI.");
    }

    return navigator.closeWindow();
  }

  Future<String> _callConnect(ServiceStorage storage) async {
    final stream = _grpcClient.connect(_outCh.stream);
    await for (var msg in stream) {
      switch (msg.whichPayload()) {
        case SrvMessage_Payload.formOpen:
          storage.addForm(msg.formID, msg.formOpen);
          break;
        case SrvMessage_Payload.rootListOpen:
          storage.addRootListForm(msg.formID, msg.rootListOpen);
          break;
        case SrvMessage_Payload.rootListAddRows:
          storage.onRootListAddRows(msg.formID, msg.rootListAddRows);
          break;
        case SrvMessage_Payload.rootListChangeRows:
          storage.onRootListChangeRows(msg.formID, msg.rootListChangeRows);
          break;
        case SrvMessage_Payload.rootListRemoveRows:
          storage.onRootListRemoveRows(msg.formID, msg.rootListRemoveRows);
          break;
        case SrvMessage_Payload.contextMenuOpen:
          storage.addContextMenu(msg.formID, msg.contextMenuOpen);
          break;
        case SrvMessage_Payload.userMessage:
          storage.onUserMessage(msg.userMessage);
          break;
        case SrvMessage_Payload.closeForm:
          storage.onCloseForm(msg.formID);
          break;
        case SrvMessage_Payload.hideUI:
          storage.onHideUI(msg.hideUI);
          break;
        case SrvMessage_Payload.closeUI:
          storage.onCloseUI();
          break;
        case SrvMessage_Payload.notSet:
          return "Recv unknown message";
        default:
          return "Recv undefined message";
      }
    }

    return "stream ended";
  }
}
