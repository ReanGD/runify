import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/command.dart';
import 'package:runify/model/command_menu.dart';
import 'package:runify/model/command_loader.dart';
import 'package:runify/view/commands/action_list.dart';
import 'package:runify/view/commands/command_list.dart';
import 'package:runify/view/commands/command_menu_dialog.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';

class CommandScreen extends StatelessWidget {
  final actionsController = ActionListController();
  final commandsController = CommandListController();

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
          ChangeNotifierProvider<CommandMenu>(
            create: (_) => CommandMenu(),
          ),
        ],
        child: DisableFocusTrapBehavior(
          child: Shortcuts(
              shortcuts: commandsController.getShortcuts(),
              child: Stack(
                children: <Widget>[
                  Actions(
                    actions: commandsController.getActions(),
                    child: CommandList(controller: commandsController),
                  ),
                  Actions(
                    actions: actionsController.getActions(),
                    child: CommandMenuDialog(controller: actionsController),
                  ),
                ],
              )),
        ),
      ),
    );
  }
}
