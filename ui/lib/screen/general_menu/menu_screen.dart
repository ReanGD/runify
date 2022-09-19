import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:runify/screen/general_menu/menu_view.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';
import 'package:runify/screen/general_menu/menu_controller.dart';

class MenuScreen extends StatelessWidget {
  final MenuController controller;

  const MenuScreen(this.controller, {super.key});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final dialogTheme = theme.dialogTheme;
    final dividerTheme = theme.dividerTheme;

    return Padding(
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
              shortcuts: controller.listController.getShortcuts(),
              child: Actions(
                actions: controller.listController.getActions(),
                child: SizedBox(
                  width: dialogTheme.actionsWidth,
                  height: dialogTheme.actionsHeight,
                  child: MenuView(controller),
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
