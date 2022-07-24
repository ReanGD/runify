import 'dart:convert';
import 'dart:io';

import 'package:runify/model/command.dart';

class CommandStorage {
  Future<List<Command>> get() async {
    final extensionPath =
        "${Directory.current.parent.path}/extensions/applications.py";
    final result = await Process.run("python", [extensionPath]);
    return List<Command>.from(jsonDecode(result.stdout)
        .map((data) => ApplicationCommand.fromJson(data))
        .toList());
  }
}
