import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:runify/screen/router.dart';
import 'package:runify/actions/actions.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final router = ScreenRouter();

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
        home: router,
      ),
    ),
  );
}
