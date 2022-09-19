import 'package:runify/text.dart';
import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:runify/widgets/hdivider.dart';
import 'package:runify/widgets/command_card.dart';
import 'package:runify/widgets/search_field.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/screen/general/gen_controller.dart';

class GenView extends StatelessWidget {
  final GenController controller;

  const GenView(this.controller, {super.key});

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
              controller.onListItemEvent(context, event, filter[id]);
            },
            itemBuilder: (context, int id) {
              final item = filter[id];
              return CommandCard(
                name: item.name,
                category: item.category,
                icon: item.icon,
              );
            },
          ),
        ),
      ],
    );
  }
}
