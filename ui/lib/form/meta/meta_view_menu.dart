import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/form/meta/meta_model.dart';
import 'package:runify/form/meta/meta_controller.dart';
import 'package:runify/form/meta/meta_widget_menu.dart';
import 'package:runify/widgets/disable_focus_trap_behavior.dart';

class MetaViewMenu extends StatelessWidget {
  final MetaMenuModel model;
  final MetaController controller;
  final listController = ListControllerMenu();

  MetaViewMenu(this.model, this.controller, {super.key});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final dialogTheme = theme.dialogTheme;
    final dividerTheme = theme.dividerTheme;

    return ChangeNotifierProvider.value(
      value: model.filter,
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
                shortcuts: listController.getShortcuts(),
                child: Actions(
                  actions: listController.getActions(),
                  child: SizedBox(
                    width: dialogTheme.actionsWidth,
                    height: dialogTheme.actionsHeight,
                    child: MetaWidgetMenu(
                      controller,
                      listController,
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
