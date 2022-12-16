import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:runify/widgets/text_size.dart';

class CommandCard extends StatelessWidget {
  final String name;
  final String? category;
  final String? icon;

  const CommandCard({super.key, required this.name, this.category, this.icon});

  Widget _getIcon(String? path, double size) {
    if (path == null || path.isEmpty) {
      return Icon(Icons.settings, size: size);
    }

    return Image.asset(
      path,
      width: size,
      height: size,
      cacheWidth: size.toInt(),
      cacheHeight: size.toInt(),
      filterQuality: FilterQuality.high,
      isAntiAlias: true,
    );
  }

  Widget buildCalc(BuildContext context) {
    final theme = Theme.of(context);
    final nameStyle = theme.textTheme.majorText?.copyWith(fontSize: 20);

    final items = name.substring(1).split(" = ");
    final left = items[0];
    final right = items[1];

    final row = Flexible(
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: <Widget>[
          Text(
            left,
            style: nameStyle,
          ),
          Text(
            right,
            style: nameStyle,
          ),
        ],
      ),
    );

    return Padding(
      padding: theme.cardTheme.commandPadding,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: <Widget>[row],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    if (name.startsWith("|")) {
      return buildCalc(context);
    }

    final theme = Theme.of(context);
    final nameStyle = theme.textTheme.majorText;
    final iconSize = TextSizeCalculator.instance
        .getCachedHeight(context, "majorText", nameStyle);

    final nameWidget = Flexible(
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          _getIcon(icon, iconSize),
          Text(
            "  $name",
            style: nameStyle,
          ),
        ],
      ),
    );

    final widgets = <Widget>[
      nameWidget,
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
