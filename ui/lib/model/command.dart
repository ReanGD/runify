import 'dart:io';

abstract class Command {
  String name();
  String category();

  void execute();
}

class ApplicationCommand extends Command {
  final String _name;
  final String _applicationPath;

  ApplicationCommand.fromJson(Map<String, dynamic> json)
      : _name = json["name"] as String,
        _applicationPath = json["app"] as String;

  @override
  String name() {
    return _name;
  }

  @override
  String category() {
    return "Application";
  }

  @override
  void execute() {
    Process.run("dex", [_applicationPath]);
  }
}
