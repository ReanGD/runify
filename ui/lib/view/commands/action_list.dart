import 'package:runify/text.dart';
import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/command.dart';
import 'package:runify/widgets/hdivider.dart';
import 'package:runify/widgets/command_card.dart';
import 'package:runify/widgets/search_field.dart';
import 'package:runify/widgets/data_list_view.dart';

class ActionListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<CommandActionFilter>().visibleItems;
  }
}

class ActionList extends StatelessWidget {
  final ActionListController controller;

  const ActionList({super.key, required this.controller});

  @override
  Widget build(BuildContext context) {
    final storage = context.read<CommandActionFilter>();

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
              controller: controller,
              onDataItemEvent: (DataItemEvent event, int? id) {
                // print("event: $event, id: $id");
              },
              itemBuilder: (context, int id) {
                final item = storage[id];
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
            onChanged: (String query) => storage.applyFilter(query),
          ),
        ),
      ],
    );
  }
}
