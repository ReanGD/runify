import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/command.dart';
import 'package:runify/model/data_filter.dart';
import 'package:runify/form/meta/meta_widget.dart';
import 'package:runify/form/meta/meta_controller.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';

class MetaView extends StatelessWidget {
  final MetaController controller;
  final DataFilter<Command> model;
  final listController = ListController();

  MetaView(this.controller, this.model, {super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ChangeNotifierProvider.value(
        value: model,
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
