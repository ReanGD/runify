import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:runify/actions/actions.dart';
import 'package:runify/view/run/run_screen.dart';
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
        initialRoute: "/run",
        routes: <String, WidgetBuilder>{
          "/run": (BuildContext context) => const RunScreen(),
          "/settings": (BuildContext context) => const SettingsScreen(),
        },
      ),
    ),
  );
}
