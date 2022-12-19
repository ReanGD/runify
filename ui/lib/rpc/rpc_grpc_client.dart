import 'dart:io';
import 'dart:async';

import 'package:grpc/grpc.dart';

import 'package:runify/screen/router.dart';
import 'package:runify/system/logger.dart';
import 'package:runify/system/settings.dart';
import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/rpc/rpc_proto_client.dart';
import 'package:runify/rpc/rpc_form_handler_storage.dart';

class GrpcClient {
  final _outCh = StreamController<UIMessage>();
  final ScreenRouter _router;
  late final Logger _logger;
  late final RunifyClient _grpcClient;
  late final FormHandlerStorage _handlers;

  GrpcClient(Settings settings, this._router) {
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
    final client = ProtoClient(0, _outCh);
    _logger = Logger(client);
    _handlers = FormHandlerStorage(_outCh, _router, _logger);
  }

  Logger get logger => _logger;

  Future<void> _exit(String error) async {
    // ignore: avoid_print
    print("gRPC stream ended with error: $error. Stop runify.");
    return _router.closeWindow();
  }

  Future<void> connect() async {
    try {
      final stream = _grpcClient.connect(_outCh.stream);
      await for (var msg in stream) {
        switch (msg.whichPayload()) {
          case SrvMessage_Payload.rootListOpen:
            _handlers.addRootListForm(msg.formID, msg.rootListOpen);
            break;
          case SrvMessage_Payload.rootListAddRows:
            _handlers.onRootListAddRows(msg.formID, msg.rootListAddRows);
            break;
          case SrvMessage_Payload.rootListChangeRows:
            _handlers.onRootListChangeRows(msg.formID, msg.rootListChangeRows);
            break;
          case SrvMessage_Payload.rootListRemoveRows:
            _handlers.onRootListRemoveRows(msg.formID, msg.rootListRemoveRows);
            break;
          case SrvMessage_Payload.contextMenuOpen:
            _handlers.addContextMenu(msg.formID, msg.contextMenuOpen);
            break;
          case SrvMessage_Payload.userMessage:
            _handlers.onUserMessage(msg.userMessage);
            break;
          case SrvMessage_Payload.closeForm:
            _handlers.onCloseForm(msg.formID);
            break;
          case SrvMessage_Payload.hideUI:
            _handlers.onHideUI(msg.hideUI);
            break;
          case SrvMessage_Payload.closeUI:
            _handlers.onCloseUI();
            break;

          case SrvMessage_Payload.notSet:
            _exit("notSet");
            break;
          default:
            _exit("unknown method");
            break;
        }
      }
    } catch (e) {
      _exit(e.toString());
    }
  }
}
