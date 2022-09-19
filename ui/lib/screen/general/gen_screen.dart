import 'package:flutter/material.dart';
import 'package:runify/screen/general/gen_view.dart';
import 'package:runify/screen/general/gen_controller.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';

class GenScreen extends StatelessWidget {
  final GenController controller;

  const GenScreen(this.controller, {super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: DisableFocusTrapBehavior(
        child: Shortcuts(
          shortcuts: controller.listController.getShortcuts(),
          child: Actions(
            actions: controller.listController.getActions(),
            child: GenView(controller),
          ),
        ),
      ),
    );
  }
}
