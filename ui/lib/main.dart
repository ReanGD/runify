import 'dart:async';

import 'package:flutter/material.dart';

import 'package:runify/style.dart';
import 'package:runify/system/settings.dart';
import 'package:runify/rpc/rpc_grpc_client.dart';
import 'package:runify/plugin/runify_native.dart';
import 'package:runify/navigator/nav_builder.dart';

void main() async {
  final instance = WidgetsFlutterBinding.ensureInitialized();
  final settings = Settings();
  final runifyPlugin = RunifyNative();
  final pluginFuture = runifyPlugin.initPlugin();
  final grpcClient = GrpcClient(settings);
  final builder = NavBuilder(settings, grpcClient, runifyPlugin);
  await pluginFuture;

  // instance.scheduleAttachRootWidget()
  Timer.run(() {
    instance.attachRootWidget(
      ExcludeSemantics(
        child: MaterialApp(
          title: "Runify",
          shortcuts: builder.getShortcuts(),
          actions: builder.getActions(),
          debugShowCheckedModeBanner: false,
          theme: getLightTheme(),
          darkTheme: getDarkTheme(),
          themeMode: ThemeMode.system,
          home: builder,
        ),
      ),
    );
  });

  instance.scheduleWarmUpFrame();
}
