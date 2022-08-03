import 'package:runify/style.dart';
import 'package:runify/text.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/widgets/hdivider.dart';
import 'package:runify/model/cmd_storage.dart';
import 'package:runify/widgets/command_card.dart';
import 'package:runify/widgets/search_field.dart';
import 'package:runify/widgets/data_list_view.dart';

class CommandListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<CommandStorage>().getVisibleItems();
  }
}

class CommandList extends StatelessWidget {
  final DataListController controller = CommandListController();

  CommandList({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final storage = context.read<CommandStorage>();

    final theme = Theme.of(context);
    final cardTheme = theme.cardTheme;
    final dialogTheme = theme.dialogTheme;

    final dividerPadding = EdgeInsets.only(
        top: dialogTheme.horizontalOffset,
        bottom: dialogTheme.horizontalOffset);
    final windowPadding =
        EdgeInsets.symmetric(horizontal: dialogTheme.horizontalOffset);
    final searchFieldPadding = EdgeInsets.only(
      top: cardTheme.commandPadding.top + dialogTheme.verticalOffset,
      left: cardTheme.commandPadding.left,
      right: cardTheme.commandPadding.right,
    );

    return Shortcuts(
      shortcuts: controller.getShortcuts(),
      child: Actions(
        actions: controller.getActions(),
        child: Column(
          children: <Widget>[
            Padding(
              padding: windowPadding,
              child: SearchField(
                padding: searchFieldPadding,
                hintText: UIText.searchFieldHint,
                onChanged: (String query) => storage.applyFilter(query),
              ),
            ),
            HDivider(padding: dividerPadding),
            Expanded(
              child: Focus(
                canRequestFocus: true,
                child: DataListView(
                  controller: controller,
                  padding: windowPadding,
                  onDataItemEvent: (DataItemEvent event, int? id) {
                    // print("event: $event, id: $id");
                  },
                  itemBuilder: (context, int id) {
                    final item = storage[id];
                    return CommandCard(
                      name: item.name(),
                      category: item.category(),
                      iconPath: item.iconPath(),
                    );
                  },
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
