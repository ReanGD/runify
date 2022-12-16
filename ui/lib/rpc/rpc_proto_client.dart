import 'dart:async';

import 'package:runify/pb/runify.pbgrpc.dart';

class ProtoClient {
  final int _formID;
  final StreamController<UIMessage> _outCh;

  ProtoClient(this._formID, this._outCh);

  writeLog(LogLevel level, String message) {
    final msg = UIMessage(
      formID: _formID,
      writeLog: WriteLog(
        level: level,
        message: message,
      ),
    );

    _outCh.add(msg);
  }

  filterChanged(String filter) {
    final msg = UIMessage(
      formID: _formID,
      filterChanged: FilterData(value: filter),
    );

    _outCh.add(msg);
  }

  rootListRowActivated(int providerID, int rowID) {
    final msg = UIMessage(
      formID: _formID,
      rootListRowActivated:
          RootListRowGlobalID(providerID: providerID, rowID: rowID),
    );

    _outCh.add(msg);
  }

  rootListMenuActivated(int providerID, int rowID) {
    final msg = UIMessage(
      formID: _formID,
      rootListMenuActivated:
          RootListRowGlobalID(providerID: providerID, rowID: rowID),
    );

    _outCh.add(msg);
  }

  contextMenuRowActivated(int rowID) {
    final msg = UIMessage(
      formID: _formID,
      contextMenuRowActivated: ContextMenuRowID(rowID: rowID),
    );

    _outCh.add(msg);
  }

  formClosed() {
    final msg = UIMessage(
      formID: _formID,
      formClosed: FormClosed(),
    );

    _outCh.add(msg);
  }
}
