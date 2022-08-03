import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/command.dart';
import 'package:runify/model/command_loader.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/view/commands/action_list.dart';
import 'package:runify/view/commands/command_list.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';

class CommandListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<CommandFilter>().visibleItems;
  }
}

class ActionListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return [0, 1, 2, 3, 4, 5, 6];
  }
}

class CommandScreen extends StatelessWidget {
  final DataListController commandsController = CommandListController();
  final ActionListController actionsController = ActionListController();

  CommandScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final dialogTheme = theme.dialogTheme;

    return Scaffold(
      body: MultiProvider(
        providers: [
          ChangeNotifierProvider<CommandFilter>(
              create: (_) =>
                  CommandFilter(CommandLoader.instance.applications)),
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
                  Positioned(
                    right: dialogTheme.horizontalOffset,
                    bottom: dialogTheme.verticalOffset,
                    child: Actions(
                      actions: actionsController.getActions(),
                      child: ActionList(controller: actionsController),
                    ),
                  ),
                ],
              )),
        ),
      ),
    );
  }
}
