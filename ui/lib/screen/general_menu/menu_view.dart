import 'package:runify/text.dart';
import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:runify/widgets/hdivider.dart';
import 'package:runify/widgets/command_card.dart';
import 'package:runify/widgets/search_field.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/screen/general_menu/menu_controller.dart';

class MenuView extends StatelessWidget {
  final MenuController controller;

  const MenuView(this.controller, {super.key});

  @override
  Widget build(BuildContext context) {
    final filter = controller.filter;
    final theme = Theme.of(context);
    final cardTheme = theme.cardTheme;
    final dialogTheme = theme.dialogTheme;

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
              onDataItemEvent: (DataItemEvent event, int id) {
                controller.onListItemEvent(event, filter[id]);
              },
              itemBuilder: (context, int id) {
                final item = filter[id];
                return CommandCard(
                  name: item.name,
                );
              },
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
