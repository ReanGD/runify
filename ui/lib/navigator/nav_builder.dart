import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'package:runify/system/settings.dart';
import 'package:runify/rpc/rpc_grpc_client.dart';
import 'package:runify/navigator/navigator.dart';
import 'package:runify/plugin/runify_native.dart';

class OnBackIntent extends Intent {
  const OnBackIntent();
}

class OnBackAction extends Action<OnBackIntent> {
  void Function()? callback;

  OnBackAction();

  @override
  Object? invoke(covariant OnBackIntent intent) {
    callback?.call();
    return null;
  }
}

class NavBuilder extends StatelessWidget {
  final Settings _settings;
  final GrpcClient _grpcClient;
  final RunifyNative _runifyPlugin;
  final _onBackAction = OnBackAction();

  NavBuilder(this._settings, this._grpcClient, this._runifyPlugin, {super.key});

  Map<Type, Action<Intent>> get actions {
    return <Type, Action<Intent>>{
      ...WidgetsApp.defaultActions,
      OnBackIntent: _onBackAction,
    };
  }

  Map<ShortcutActivator, Intent> get shortcuts {
    return <ShortcutActivator, Intent>{
      ...WidgetsApp.defaultShortcuts,
      const SingleActivator(LogicalKeyboardKey.escape): const OnBackIntent(),
    };
  }

  @override
  Widget build(BuildContext context) {
    final navigator = RunifyNavigator(
        _settings, _runifyPlugin, Navigator.of(context), _grpcClient.logger);
    _onBackAction.callback = () => navigator.back();
    _grpcClient.start(navigator);
    final theme = Theme.of(context);
    return Container(
      color: theme.backgroundColor,
    );
  }
}
