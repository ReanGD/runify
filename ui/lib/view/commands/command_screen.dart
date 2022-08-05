import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/command.dart';
import 'package:runify/model/command_loader.dart';
import 'package:runify/view/commands/command_list.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';

class CommandScreen extends StatelessWidget {
  final controller = CommandListController();

  CommandScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: MultiProvider(
        providers: [
          ChangeNotifierProvider<CommandFilter>(
            create: (_) =>
                CommandFilter.future(CommandLoader.instance.applications),
          ),
        ],
        child: DisableFocusTrapBehavior(
          child: Shortcuts(
            shortcuts: controller.getShortcuts(),
            child: Actions(
              actions: controller.getActions(),
              child: CommandList(controller: controller),
            ),
          ),
        ),
      ),
    );
  }
}
