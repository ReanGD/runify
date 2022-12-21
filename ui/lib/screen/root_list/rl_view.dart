import 'package:flutter/material.dart';

import 'package:runify/text.dart';
import 'package:runify/style.dart';
import 'package:runify/widgets/hdivider.dart';
import 'package:runify/widgets/search_field.dart';
import 'package:runify/widgets/data_list_row.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/screen/root_list/rl_controller.dart';

class RLView extends StatelessWidget {
  final RLController controller;

  const RLView(this.controller, {super.key});

  @override
  Widget build(BuildContext context) {
    final filter = controller.filter;
    final theme = Theme.of(context);
    final cardTheme = theme.cardTheme;
    final dialogTheme = theme.dialogTheme;

    final dividerPadding = EdgeInsets.only(
        top: dialogTheme.verticalOffset, bottom: dialogTheme.verticalOffset);
    final windowPadding =
        EdgeInsets.symmetric(horizontal: dialogTheme.horizontalOffset);
    final searchFieldPadding = EdgeInsets.only(
      top: cardTheme.commandPadding.top + dialogTheme.verticalOffset,
      left: cardTheme.commandPadding.left,
      right: cardTheme.commandPadding.right,
    );

    return Column(
      children: <Widget>[
        Padding(
          padding: windowPadding,
          child: SearchField(
            padding: searchFieldPadding,
            hintText: UIText.searchCommandHint,
            onChanged: (String query) => controller.onApplyFilter(query),
          ),
        ),
        HDivider(padding: dividerPadding),
        Expanded(
          child: DataListView(
            controller: controller.listController,
            padding: windowPadding,
            onDataItemEvent: (DataItemEvent event, int id) {
              controller.onListItemEvent(event, filter[id]);
            },
            itemBuilder: (context, int id) => RootListRowWidget(filter[id]),
          ),
        ),
      ],
    );
  }
}
