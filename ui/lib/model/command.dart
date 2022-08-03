import 'dart:io';

import 'package:runify/model/data_filter.dart';

abstract class Command implements Matcher {
  String name();
  String category();
  String? iconPath();

  void execute();
}

typedef CommandFilter = DataFilter<Command>;

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
  bool match(String filter) {
    return _name.toLowerCase().contains(filter);
  }

  @override
  void execute() {
    Process.run("dex", [_applicationPath]);
  }
}
