import 'dart:io';

import 'package:runify/model/command.dart';

class ExecuteApplicationAction extends CommandAction {
  final String _applicationPath;

  ExecuteApplicationAction(this._applicationPath);

  @override
  String get name => "Execute";

  @override
  bool match(String filter) {
    return name.toLowerCase().contains(filter);
  }

  @override
  void execute() {
    Process.run("dex", [_applicationPath]);
  }
}

class CopyToClipboardAction extends CommandAction {
  final String _name;
  final String _value;

  CopyToClipboardAction(this._name, this._value);

  @override
  String get name => _name;

  @override
  bool match(String filter) {
    return name.toLowerCase().contains(filter);
  }

  @override
  void execute() {
    // ignore: avoid_print
    print("copy $_value to clipboard");
  }
}

class ApplicationCommand extends Command {
  final String _name;
  final String _applicationPath;
  final String? _iconPath;

  ApplicationCommand.fromJson(Map<String, dynamic> json)
      : _name = json["name"] as String,
        _applicationPath = json["app"] as String,
        _iconPath = json["icon"] == null ? null : json["icon"] as String;

  @override
  String get name => _name;

  @override
  String get category => "Application";

  @override
  String? get iconPath => _iconPath;

  @override
  CommandActions get actions {
    return [
      ExecuteApplicationAction(_applicationPath),
      CopyToClipboardAction("Copy name to clipboard", _name),
      CopyToClipboardAction("Copy path to clipboard", _applicationPath),
    ];
  }

  @override
  bool match(String filter) {
    return _name.toLowerCase().contains(filter);
  }
}
