import 'dart:convert';
import 'dart:io';

import 'package:runify/model/cmd.dart';

class CommandLoader {
  late Future<List<Command>> _commands;

  CommandLoader._() {
    _commands = _get();
  }

  static final CommandLoader instance = CommandLoader._();

  Future<List<Command>> get commands => _commands;

  Future<List<Command>> _get() async {
    final extensionPath =
        "${Directory.current.parent.path}/extensions/applications.py";
    final result = await Process.run("python", [extensionPath]);
    return List<Command>.from(jsonDecode(result.stdout)
        .map((data) => ApplicationCommand.fromJson(data))
        .toList());
  }
}
