import 'package:flutter/material.dart';

import 'package:runify/style.dart';

class BottomHelperBar extends StatelessWidget {
  final List<Widget> children;

  const BottomHelperBar(this.children, {super.key});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final colorScheme = theme.colorScheme;
    final dialogTheme = theme.dialogTheme;
    final navigationBarTheme = theme.navigationBarTheme;

    const height = 47.0;
    final padding =
        EdgeInsets.symmetric(horizontal: dialogTheme.horizontalOffset);

    final rowChildren = <Widget>[];
    for (var i = 0; i < children.length; i++) {
      rowChildren.add(children[i]);
      if (i != children.length - 1) {
        rowChildren.add(const SizedBox(width: 8));
      }
    }

    return Material(
      color: navigationBarTheme.backgroundColor ?? colorScheme.surface,
      elevation: navigationBarTheme.elevation ?? 3.0,
      shadowColor: navigationBarTheme.shadowColor ?? Colors.transparent,
      surfaceTintColor:
          navigationBarTheme.surfaceTintColor ?? colorScheme.surfaceTint,
      child: SizedBox(
        height: height,
        child: Padding(
          padding: padding,
          child: Row(
            mainAxisAlignment: MainAxisAlignment.end,
            children: rowChildren,
          ),
        ),
      ),
    );
  }
}
