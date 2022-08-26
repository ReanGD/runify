import 'package:runify/text.dart';
import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/command.dart';
import 'package:runify/widgets/hdivider.dart';
import 'package:runify/model/data_provider.dart';
import 'package:runify/widgets/command_card.dart';
import 'package:runify/widgets/search_field.dart';
import 'package:runify/widgets/data_list_view.dart';
import 'package:runify/view/commands/command_menu_dialog.dart';

class CommandListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<CommandFilter>().visibleItems;
  }
}

class CommandList extends StatelessWidget {
  final CommandListController controller;

  const CommandList({super.key, required this.controller});

  void _onDataItemEvent(
      BuildContext context, DataItemEvent event, Command? command) {
    if (event == DataItemEvent.onMenu && command != null) {
      CommandMenuDialog.show(context, command);
      return;
    }
    if (event == DataItemEvent.onChoice && command != null) {
      DataProvider.instance.execute(command.id, 0);
      return;
    }
  }

  @override
  Widget build(BuildContext context) {
    final storage = context.read<CommandFilter>();

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
            onChanged: (String query) => storage.applyFilter(query),
          ),
        ),
        HDivider(padding: dividerPadding),
        Expanded(
          child: DataListView(
            controller: controller,
            padding: windowPadding,
            onDataItemEvent: (DataItemEvent event, int? id) {
              _onDataItemEvent(
                  context, event, (id != null) ? storage[id] : null);
            },
            itemBuilder: (context, int id) {
              final item = storage[id];
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
