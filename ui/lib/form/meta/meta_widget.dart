import 'package:runify/text.dart';
import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/command.dart';
import 'package:runify/widgets/hdivider.dart';
import 'package:runify/widgets/command_card.dart';
import 'package:runify/form/meta/meta_controller.dart';
import 'package:runify/widgets/search_field.dart';
import 'package:runify/widgets/data_list_view.dart';

class ListController extends DataListController {
  @override
  List<int> getVisibleItems(BuildContext context) {
    return context.watch<CommandFilter>().visibleItems;
  }
}

class MetaWidget extends StatelessWidget {
  final MetaController controller;
  final ListController listController;

  const MetaWidget(this.controller, this.listController, {super.key});

  @override
  Widget build(BuildContext context) {
    final model = context.read<CommandFilter>();

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
            onChanged: (String query) => model.applyFilter(query),
          ),
        ),
        HDivider(padding: dividerPadding),
        Expanded(
          child: DataListView(
            controller: listController,
            padding: windowPadding,
            onDataItemEvent: (DataItemEvent event, int id) {
              controller.onFormListItemEvent(context, event, model[id]);
            },
            itemBuilder: (context, int id) {
              final item = model[id];
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
