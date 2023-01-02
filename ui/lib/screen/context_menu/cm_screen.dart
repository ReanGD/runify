import 'package:flutter/material.dart';

import 'package:runify/style.dart';
import 'package:runify/screen/context_menu/cm_view.dart';
import 'package:runify/screen/context_menu/cm_controller.dart';

class CMScreen extends StatelessWidget {
  final CMController controller;

  const CMScreen(this.controller, {super.key});

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
          child: Shortcuts(
            shortcuts: controller.listController.shortcuts,
            child: Actions(
              actions: controller.listController.actions,
              child: FocusScope(
                autofocus: true,
                child: SizedBox(
                  width: dialogTheme.actionsWidth,
                  height: dialogTheme.actionsHeight,
                  child: CMView(controller),
                ),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
