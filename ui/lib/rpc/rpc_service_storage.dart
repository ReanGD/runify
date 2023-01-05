import 'dart:async';

import 'package:runify/system/logger.dart';
import 'package:runify/rpc/rpc_types.dart';
import 'package:runify/pb/runify.pbgrpc.dart';
import 'package:runify/navigator/navigator.dart';
import 'package:runify/rpc/rpc_proto_client.dart';
import 'package:runify/rpc/rpc_form_service.dart';
import 'package:runify/rpc/rpc_root_list_service.dart';
import 'package:runify/rpc/rpc_context_menu_service.dart';

class ServiceStorage {
  final _services = <int, Service>{};

  final Logger _logger;
  final RunifyNavigator _navigator;
  final StreamController<UIMessage> _outCh;

  ServiceStorage(this._outCh, this._navigator, this._logger);

  Future<void> addForm(int formID, FormOpen msg) async {
    final pClient = ProtoClient(formID, _outCh);
    final service = FMService(this, pClient, _logger, msg.markup, msg.model);
    _services[formID] = service;
    await _navigator.openForm(service);
  }

  Future<void> addRootListForm(int formID, RootListOpen msg) async {
    final pClient = ProtoClient(formID, _outCh);
    final service = RLService(this, pClient, _logger, msg.rows);
    _services[formID] = service;
    await _navigator.openRootList(service);
  }

  Future<void> addContextMenu(int formID, ContextMenuOpen msg) async {
    final pClient = ProtoClient(formID, _outCh);
    final service = CMService(this, pClient, _logger, msg.rows);
    _services[formID] = service;
    _navigator.openContexMenu(service);
  }

  bool remove(int formID) {
    return _services.remove(formID) != null;
  }

  Service getForHandle(int formID, String msgName) {
    final handler = _services[formID];
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

  Future<void> onFieldCheckResponse(int formID, FieldCheckResponse msg) async {
    getForHandle(formID, "FieldCheckResponse").onFieldCheckResponse(msg);
  }

  Future<void> onUserMessage(UserMessage msg) async {
    // TODO: implement onUserMessage
  }

  Future<void> onCloseForm(int formID) async {
    if (remove(formID)) {
      _navigator.popForm(formID);
    }
  }

  Future<void> onHideUI(HideUI msg) async {
    // TODO: show message if needed
    return _navigator.hideWindow();
  }

  Future<void> onCloseUI() async {
    return _navigator.closeWindow();
  }
}
