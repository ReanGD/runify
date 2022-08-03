import 'dart:io';

abstract class Command {
  String name();
  String category();
  String? iconPath();

  void execute();
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
  String name() {
    return _name;
  }

  @override
  String category() {
    return "Application";
  }

  @override
  String? iconPath() {
    return _iconPath;
  }

  @override
  void execute() {
    Process.run("dex", [_applicationPath]);
  }
}
