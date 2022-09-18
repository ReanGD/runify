import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/form/meta/meta_model.dart';
import 'package:runify/form/meta/meta_widget.dart';
import 'package:runify/form/meta/meta_controller.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';

class MetaView extends StatelessWidget {
  final MetaModel model;
  final MetaController controller;
  final listController = ListController();

  MetaView(this.model, this.controller, {super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ChangeNotifierProvider.value(
        value: model.filter,
        child: DisableFocusTrapBehavior(
          child: Shortcuts(
            shortcuts: listController.getShortcuts(),
            child: Actions(
              actions: listController.getActions(),
              child: MetaWidget(controller, listController),
            ),
          ),
        ),
      ),
    );
  }
}
