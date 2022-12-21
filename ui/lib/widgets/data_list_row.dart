import 'package:runify/global/context_menu_row.dart';
import 'package:runify/global/root_list_row.dart';
import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:runify/widgets/text_size.dart';

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

    final left = Flexible(
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          _getIcon(data.icon, iconSize),
          Text(
            "  ${data.value}",
            style: theme.textTheme.majorText,
          ),
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
    final nameStyle = theme.textTheme.majorText?.copyWith(fontSize: 20);

    final items = data.value.split("\n");
    final left = Text(
      items[0],
      style: nameStyle,
    );
    final right = Text(
      items[1],
      style: nameStyle,
    );

    return Padding(
      padding: theme.cardTheme.commandPadding,
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceAround,
        children: <Widget>[left, right],
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

    final left = Flexible(
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        children: <Widget>[
          _getIcon(null, iconSize),
          Text(
            "  ${data.value}",
            style: theme.textTheme.majorText,
          ),
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
