import 'package:runify/text.dart';
import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:runify/widgets/hdivider.dart';
import 'package:runify/widgets/command_card.dart';
import 'package:runify/widgets/search_field.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/view/commands/command_screen.dart';

class ActionList extends StatelessWidget {
  final ActionListController controller;

  const ActionList({super.key, required this.controller});

  @override
  Widget build(BuildContext context) {
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

    return SizedBox(
      width: 350,
      height: 300,
      child: DecoratedBox(
        decoration: BoxDecoration(
          color: theme.scaffoldBackgroundColor,
          border: Border.all(),
          borderRadius: const BorderRadius.all(Radius.circular(defaultRadius)),
        ),
        child: Column(
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
                    return CommandCard(
                      name: "Action $id",
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
                onChanged: (String query) {},
              ),
            ),
          ],
        ),
      ),
    );
  }
}
