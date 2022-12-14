import 'package:flutter/material.dart';

import 'package:runify/screen/root_list/rl_view.dart';
import 'package:runify/screen/root_list/rl_controller.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';

class RLScreen extends StatelessWidget {
  final RLController controller;

  const RLScreen(this.controller, {super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: DisableFocusTrapBehavior(
        child: Shortcuts(
          shortcuts: controller.listController.getShortcuts(),
          child: Actions(
            actions: controller.listController.getActions(),
            child: RLView(controller),
          ),
        ),
      ),
    );
  }
}
