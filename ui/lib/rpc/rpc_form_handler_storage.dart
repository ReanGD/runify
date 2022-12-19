import 'dart:async';

import 'package:runify/system/logger.dart';
import 'package:runify/rpc/rpc_types.dart';
import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/navigator/navigator.dart';
import 'package:runify/rpc/rpc_proto_client.dart';
import 'package:runify/rpc/rpc_form_handlers.dart';

class FormHandlerStorage {
  final _handlers = <int, FormHandler>{};

  final Logger _logger;
  final RunifyNavigator _navigator;
  final StreamController<UIMessage> _outCh;

  FormHandlerStorage(this._outCh, this._navigator, this._logger);

  Future<void> addRootListForm(int formID, RootListOpen msg) async {
    final pClient = ProtoClient(formID, _outCh);
    final handler = RootListHandler(pClient, msg.rows);
    _handlers[formID] = handler;
    await _navigator.openRootList(handler);
  }

  Future<void> addContextMenu(int formID, ContextMenuOpen msg) async {
    final pClient = ProtoClient(formID, _outCh);
    final handler = ContextMenuHandler(pClient, _logger, msg.rows);
    _handlers[formID] = handler;
    _navigator.openContexMenu(handler);
  }

  FormHandler getForHandle(int formID, String msgName) {
    final handler = _handlers[formID];
    if (handler == null) {
      _logger
          .debug("Grpc message for unknown form = $formID, message = $msgName");
    }
    return handler!;
  }

  Future<void> onRootListAddRows(int formID, RootListAddRows msg) async {
    getForHandle(formID, "RootListAddRows").onRootListAddRows(msg.rows);
  }

  Future<void> onRootListChangeRows(int formID, RootListChangeRows msg) async {
    getForHandle(formID, "RootListChangeRows").onRootListChangeRows(msg.rows);
  }

  Future<void> onRootListRemoveRows(int formID, RootListRemoveRows msg) async {
    getForHandle(formID, "RootListRemoveRows").onRootListRemoveRows(msg.rows);
  }

  Future<void> onUserMessage(UserMessage msg) async {
    // TODO: implement onUserMessage
  }

  Future<void> onCloseForm(int formID) async {
    // TODO: remove handler
    _navigator.popForm(formID);
  }

  Future<void> onHideUI(HideUI msg) async {
    // TODO: show message
    return _navigator.hideWindow();
  }

  Future<void> onCloseUI() async {
    return _navigator.closeWindow();
  }
}
