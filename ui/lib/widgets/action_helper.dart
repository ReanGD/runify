import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'package:runify/style.dart';

class ActionHelper extends StatelessWidget {
  final String name;
  final SingleActivator shortcut;
  final bool? isMinor;

  // ignore: prefer_const_constructors_in_immutables
  ActionHelper(this.name, this.shortcut, {this.isMinor, super.key});

  Widget buildKey(ThemeData theme, String key, TextStyle? textStyle) {
    const double radius = 6.0;
    return Container(
      alignment: Alignment.center,
      padding: const EdgeInsets.symmetric(horizontal: 8),
      margin: const EdgeInsets.symmetric(vertical: 8, horizontal: 1),
      decoration: BoxDecoration(
        borderRadius: const BorderRadius.all(Radius.circular(radius)),
        color: theme.focusColor,
      ),
      child: Text(
        key,
        style: textStyle,
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    final ThemeData theme = Theme.of(context);

    var children = <Widget>[
      Text(
        name,
        style: isMinor ?? false
            ? theme.textTheme.labelRegular_60
            : theme.textTheme.labelRegular,
      ),
      const SizedBox(width: 8),
    ];

    final keys = <String>[
      if (shortcut.meta) "Meta",
      if (shortcut.control) "Ctrl",
      if (shortcut.alt) "Alt",
      if (shortcut.shift) "Shift",
      if (shortcut.trigger != LogicalKeyboardKey.enter)
        shortcut.trigger.keyLabel,
    ];

    final textStyle = theme.textTheme.labelRegular_80;
    for (final key in keys) {
      children.add(buildKey(theme, key, textStyle));
    }

    if (shortcut.trigger == LogicalKeyboardKey.enter) {
      children.add(buildKey(theme, "↵", theme.textTheme.bodyRegular_80));
    }

    return Row(
      mainAxisAlignment: MainAxisAlignment.start,
      children: children,
    );
  }
}
