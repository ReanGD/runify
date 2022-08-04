import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/command.dart';
import 'package:runify/model/command_menu.dart';
import 'package:runify/view/commands/action_list.dart';

class CommandMenuDialog extends StatelessWidget {
  final ActionListController controller;

  const CommandMenuDialog({super.key, required this.controller});

  @override
  Widget build(BuildContext context) {
    final commandMenu = context.watch<CommandMenu>();
    final visible = commandMenu.visible;
    if (!visible) {
      return const SizedBox.shrink();
    }

    final theme = Theme.of(context);
    final dialogTheme = theme.dialogTheme;

    return ChangeNotifierProvider<CommandActionFilter>(
      create: (_) => commandMenu.filter,
      child: Positioned(
        right: dialogTheme.horizontalOffset,
        bottom: dialogTheme.verticalOffset,
        child: SizedBox(
          width: dialogTheme.actionsWidth,
          height: dialogTheme.actionsHeight,
          child: DecoratedBox(
            decoration: BoxDecoration(
              color: theme.scaffoldBackgroundColor,
              border: Border.all(),
              borderRadius:
                  const BorderRadius.all(Radius.circular(defaultRadius)),
            ),
            child: ActionList(
              controller: controller,
            ),
          ),
        ),
      ),
    );
  }
}
