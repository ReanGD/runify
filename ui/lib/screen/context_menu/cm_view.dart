import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:runify/text.dart';
import 'package:runify/style.dart';
import 'package:runify/global/shortcuts.dart';
import 'package:runify/widgets/hdivider.dart';
import 'package:runify/widgets/search_field.dart';
import 'package:runify/widgets/data_list_row.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/global/context_menu_row.dart';
import 'package:runify/screen/context_menu/cm_controller.dart';

class CMView extends StatelessWidget {
  final CMController controller;

  const CMView(this.controller, {super.key});

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);
    final cardTheme = theme.cardTheme;
    final dialogTheme = theme.dialogTheme;
    final filter = context.read<ContextMenuRowFilter>();

    final searchFieldPadding = EdgeInsets.only(
      bottom: cardTheme.commandPadding.bottom,
      left: cardTheme.commandPadding.left,
      right: cardTheme.commandPadding.right,
    );
    final windowPadding = EdgeInsets.symmetric(
        vertical: dialogTheme.verticalOffset,
        horizontal: dialogTheme.horizontalOffset);

    return Column(
      children: <Widget>[
        Expanded(
          child: Padding(
            padding: windowPadding,
            child: DataListView(
              controller: controller.listController,
              onDataItemEvent: (DataListEvent event, int id) {
                controller.onListItemEvent(event, filter[id]);
              },
              itemBuilder: (context, int id) =>
                  ContextMenuRowWidget(filter[id]),
            ),
          ),
        ),
        const HDivider(),
        Padding(
          padding: windowPadding,
          child: SearchField(
            padding: searchFieldPadding,
            hintText: UIText.searchActionHint,
            onChanged: (String query) => controller.onApplyFilter(query),
          ),
        ),
      ],
    );
  }
}
