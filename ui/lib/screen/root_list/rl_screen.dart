import 'package:flutter/material.dart';

import 'package:runify/screen/root_list/rl_view.dart';
import 'package:runify/screen/root_list/rl_controller.dart';

class RLScreen extends StatelessWidget {
  final RLController controller;

  const RLScreen(this.controller, {super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Shortcuts(
        shortcuts: controller.listController.shortcuts,
        child: Actions(
          actions: controller.listController.actions,
          child: FocusScope(
            autofocus: true,
            child: RLView(controller),
          ),
        ),
      ),
    );
  }
}
