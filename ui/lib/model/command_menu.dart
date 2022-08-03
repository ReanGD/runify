import 'package:flutter/material.dart';
import 'package:runify/model/command.dart';

class CommandMenu with ChangeNotifier {
  Command? _command;

  CommandMenu();

  bool get visible => _command != null;

  CommandActionFilter get filter {
    if (_command == null) {
      return CommandActionFilter();
    }

    return CommandActionFilter.value(_command!.actions);
  }

  void show(Command? command) {
    if (command != null) {
      _command = command;
      notifyListeners();
    }
  }

  void hide() {
    if (_command != null) {
      _command = null;
      notifyListeners();
    }
  }
}
