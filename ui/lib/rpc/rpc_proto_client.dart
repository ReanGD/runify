import 'dart:async';

import 'package:runify/pb/runify.pbgrpc.dart';

class ProtoClient {
  final int _formID;
  final StreamController<UIMessage> _outCh;

  ProtoClient(this._formID, this._outCh);

  int get formID => _formID;

  void writeLog(LogLevel level, String message) {
    final msg = UIMessage(
      formID: _formID,
      writeLog: WriteLog(
        level: level,
        message: message,
      ),
    );

    _outCh.add(msg);
  }

  void filterChanged(String filter) {
    final msg = UIMessage(
      formID: _formID,
      filterChanged: FilterData(value: filter),
    );

    _outCh.add(msg);
  }

  void rootListRowActivated(int providerID, int rowID) {
    final msg = UIMessage(
      formID: _formID,
      rootListRowActivated:
          RootListRowGlobalID(providerID: providerID, rowID: rowID),
    );

    _outCh.add(msg);
  }

  void rootListMenuActivated(int providerID, int rowID) {
    final msg = UIMessage(
      formID: _formID,
      rootListMenuActivated:
          RootListRowGlobalID(providerID: providerID, rowID: rowID),
    );

    _outCh.add(msg);
  }

  void contextMenuRowActivated(int rowID) {
    final msg = UIMessage(
      formID: _formID,
      contextMenuRowActivated: ContextMenuRowID(rowID: rowID),
    );

    _outCh.add(msg);
  }

  void formClosed() {
    final msg = UIMessage(
      formID: _formID,
      formClosed: FormClosed(),
    );

    _outCh.add(msg);
  }
}
