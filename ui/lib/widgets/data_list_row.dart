import 'package:flutter/material.dart';

import 'package:runify/style.dart';
import 'package:runify/widgets/text_size.dart';
import 'package:runify/global/root_list_row.dart';
import 'package:runify/global/context_menu_row.dart';

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

class RootListRowWidget extends StatelessWidget {
  final RootListRow data;

  const RootListRowWidget(this.data, {super.key});

  Widget buildDefault(BuildContext context) {
    final theme = Theme.of(context);
    final iconSize = TextSizeCalculator.instance
        .getCachedHeight(context, "majorText", theme.textTheme.majorText);

    final icon = _getIcon(data.icon, iconSize);
    final name = Flexible(
      child: Text(
        "  ${data.displayName}",
        style: theme.textTheme.majorText,
        maxLines: 1,
        overflow: TextOverflow.ellipsis,
      ),
    );

    final left = Flexible(
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          icon,
          name,
        ],
      ),
    );

    final right = Text(
      data.typeName,
      style: theme.textTheme.minorText,
    );

    return Padding(
      padding: theme.cardTheme.commandPadding,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: <Widget>[
          left,
          right,
        ],
      ),
    );
  }

  Widget buildCalc(BuildContext context) {
    final theme = Theme.of(context);
    final nameStyle = theme.textTheme.majorText;

    final items = data.displayName.split("\n");

    final left = Flexible(
      fit: FlexFit.tight,
      child: Center(
        child: Text(
          items[0],
          style: nameStyle,
        ),
      ),
    );
    final center = Text(
      " = ",
      style: nameStyle,
    );
    final right = Flexible(
      fit: FlexFit.tight,
      child: Center(
        child: Text(
          items[1],
          style: nameStyle,
        ),
      ),
    );

    return Padding(
      padding: theme.cardTheme.commandPadding,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: <Widget>[left, center, right],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    if (data.rowType != RootListRowType.calculator) {
      return buildDefault(context);
    }

    return buildCalc(context);
  }
}

class ContextMenuRowWidget extends StatelessWidget {
  final ContextMenuRow data;

  const ContextMenuRowWidget(this.data, {super.key});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final iconSize = TextSizeCalculator.instance
        .getCachedHeight(context, "majorText", theme.textTheme.majorText);

    final icon = _getIcon(null, iconSize);
    final name = Flexible(
      child: Text(
        "  ${data.displayName}",
        style: theme.textTheme.majorText,
        maxLines: 1,
        overflow: TextOverflow.ellipsis,
      ),
    );

    final left = Flexible(
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          icon,
          name,
        ],
      ),
    );

    return Padding(
      padding: theme.cardTheme.commandPadding,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: <Widget>[
          left,
        ],
      ),
    );
  }
}
