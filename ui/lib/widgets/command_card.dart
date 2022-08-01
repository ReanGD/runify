import 'package:runify/style.dart';
import 'package:flutter/material.dart';

class CommandCard extends StatelessWidget {
  final String name;
  final String? category;

  const CommandCard({super.key, required this.name, this.category});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    final widgets = <Widget>[
      Text(
        name,
        style: theme.textTheme.majorText,
      ),
    ];

    if (category != null) {
      widgets.add(
        Text(
          category!,
          style: theme.textTheme.minorText,
        ),
      );
    }

    return Padding(
      padding: theme.cardTheme.commandPadding,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: widgets,
      ),
    );
  }
}
