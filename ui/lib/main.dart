import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:runify/actions/actions.dart';
import 'package:bitsdojo_window/bitsdojo_window.dart';
import 'package:runify/view/commands/command_screen.dart';
import 'package:runify/view/settings/setting_screen.dart';

void main() {
  runApp(
    ExcludeSemantics(
      child: MaterialApp(
        title: "Runify",
        shortcuts: <LogicalKeySet, Intent>{
          closeKeySet: const CloseIntent(),
        },
        actions: <Type, Action<Intent>>{
          CloseIntent: CloseAction(),
        },
        debugShowCheckedModeBanner: false,
        theme: getLightTheme(),
        darkTheme: getDarkTheme(),
        themeMode: ThemeMode.system,
        initialRoute: "/commands",
        routes: <String, WidgetBuilder>{
          "/commands": (BuildContext context) => CommandScreen(),
          "/settings": (BuildContext context) => const SettingsScreen(),
        },
      ),
    ),
  );

  doWhenWindowReady(() {
    appWindow.minSize = const Size(1280, 720);
    appWindow.size = const Size(1280, 720);
    appWindow.alignment = Alignment.center;
    appWindow.show();
  });
}
