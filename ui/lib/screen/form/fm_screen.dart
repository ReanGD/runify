import 'package:flutter/material.dart';
import 'package:runify/global/shortcuts.dart';

import 'package:runify/screen/form/fm_view.dart';
import 'package:runify/screen/form/fm_controller.dart';

class MenuActivateAction extends Action<MenuActivateIntent> {
  final FMController controller;

  MenuActivateAction(this.controller);

  @override
  void invoke(MenuActivateIntent intent) {
    controller.onMenuActivate();
  }
}

class SubmitAction extends Action<SubmitIntent> {
  final FMController controller;

  SubmitAction(this.controller);

  @override
  void invoke(SubmitIntent intent) {
    controller.onSubmit();
  }
}

class FMScreen extends StatelessWidget {
  final FMController controller;

  const FMScreen(this.controller, {super.key});

  Map<Type, Action<Intent>> get actions {
    return <Type, Action<Intent>>{
      MenuActivateIntent: MenuActivateAction(controller),
      SubmitIntent: SubmitAction(controller),
    };
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Actions(
        actions: actions,
        child: FocusScope(
          autofocus: true,
          child: FMView(controller),
        ),
      ),
    );
  }
}
