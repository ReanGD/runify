import 'package:flutter/material.dart';

import 'package:runify/style.dart';
import 'package:runify/system/settings.dart';
import 'package:runify/global/shortcuts.dart';
import 'package:runify/rpc/rpc_grpc_client.dart';
import 'package:runify/plugin/runify_native.dart';
import 'package:runify/navigator/nav_builder.dart';

void main() async {
  final instance = WidgetsFlutterBinding.ensureInitialized();
  final settings = Settings();
  final runifyPlugin = RunifyNative();
  final pluginFuture = runifyPlugin.initPlugin();
  final grpcClient = GrpcClient(settings);
  final shortcuts = ShortcutStorage();
  final builder = NavBuilder(settings, shortcuts, grpcClient, runifyPlugin);
  await pluginFuture;

  // ignore: invalid_use_of_protected_member
  instance.scheduleAttachRootWidget(
    instance.wrapWithDefaultView(
      ExcludeSemantics(
        child: MaterialApp(
          title: "Runify",
          shortcuts: builder.shortcuts,
          actions: builder.actions,
          debugShowCheckedModeBanner: false,
          theme: getLightTheme(),
          darkTheme: getDarkTheme(),
          themeMode: ThemeMode.system,
          home: builder,
        ),
      ),
    ),
  );

  instance.scheduleWarmUpFrame();
}
