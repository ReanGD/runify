import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/command.dart';
import 'package:runify/view/commands/action_list.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';

class CommandMenuDialog extends StatelessWidget {
  final Command command;
  final controller = ActionListController();

  CommandMenuDialog({super.key, required this.command});

  static Future show(BuildContext context, Command command) {
    return showDialog(
      context: context,
      barrierColor: null,
      builder: (BuildContext context) {
        return CommandMenuDialog(command: command);
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final dialogTheme = theme.dialogTheme;
    final dividerTheme = theme.dividerTheme;

    return ChangeNotifierProvider<CommandActionFilter>(
      create: (_) => CommandActionFilter.value(command.actions),
      child: Padding(
        padding: EdgeInsets.only(
          right: dialogTheme.horizontalOffset,
          bottom: dialogTheme.verticalOffset,
        ),
        child: Align(
          alignment: AlignmentDirectional.bottomEnd,
          child: Material(
            color: theme.scaffoldBackgroundColor,
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(defaultRadius),
              side: BorderSide(color: dividerTheme.color ?? theme.dividerColor),
            ),
            child: DisableFocusTrapBehavior(
              child: Shortcuts(
                shortcuts: controller.getShortcuts(),
                child: Actions(
                  actions: controller.getActions(),
                  child: SizedBox(
                    width: dialogTheme.actionsWidth,
                    height: dialogTheme.actionsHeight,
                    child: ActionList(
                      controller: controller,
                    ),
                  ),
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
