import 'dart:async';

import 'package:flutter/material.dart';

import 'package:runify/style.dart';
import 'package:runify/screen/router.dart';

void main() async {
  final instance = WidgetsFlutterBinding.ensureInitialized();
  final router = ScreenRouter();
  await router.init();

  // instance.scheduleAttachRootWidget()
  Timer.run(() {
    instance.attachRootWidget(
      ExcludeSemantics(
        child: MaterialApp(
          title: "Runify",
          shortcuts: router.getShortcuts(),
          actions: router.getActions(),
          debugShowCheckedModeBanner: false,
          theme: getLightTheme(),
          darkTheme: getDarkTheme(),
          themeMode: ThemeMode.system,
          home: router,
        ),
      ),
    );
  });

  instance.scheduleWarmUpFrame();
}
