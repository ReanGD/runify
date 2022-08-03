import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:runify/widgets/text_size.dart';

class CommandCard extends StatelessWidget {
  final String name;
  final String? category;
  final String? iconPath;

  const CommandCard(
      {super.key, required this.name, this.category, this.iconPath});

  Widget _getIcon(String? path, double size) {
    if (path == null) {
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

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final nameStyle = theme.textTheme.majorText;
    final iconSize = TextSizeCalculator.instance
        .getCachedHeight(context, "majorText", nameStyle);

    final nameWidget = Flexible(
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          _getIcon(iconPath, iconSize),
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
