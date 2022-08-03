import 'dart:io';
import 'dart:convert';

import 'package:runify/model/command.dart';

class CommandLoader {
  late Future<List<Command>> _applications;

  CommandLoader._() {
    _applications = _loadApplications();
  }

  static final CommandLoader instance = CommandLoader._();

  Future<List<Command>> get applications => _applications;

  Future<List<Command>> _loadApplications() async {
    final extensionPath =
        "${Directory.current.parent.path}/extensions/applications.py";
    final result = await Process.run("python", [extensionPath]);
    return List<Command>.from(jsonDecode(result.stdout)
        .map((data) => ApplicationCommand.fromJson(data))
        .toList());
  }
}
