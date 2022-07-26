import 'dart:math';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:provider/provider.dart';
import 'package:runify/model/cmd_storage.dart';
import 'package:runify/model/select_index.dart';
import 'package:runify/view/run/search_widget.dart';
import 'package:runify/view/run/filterable_list.dart';
import 'package:scrollable_positioned_list/scrollable_positioned_list.dart';

class ItemsList extends StatefulWidget {
  const ItemsList({Key? key}) : super(key: key);

  @override
  State<ItemsList> createState() => _ItemsListState();
}

class _ItemsListState extends State<ItemsList> {
  final _itemScrollController = ItemScrollController();
  final _itemPositionsListener = ItemPositionsListener.create();

  @override
  void initState() {
    super.initState();
  }

  void _applyFilter(String query) {
    final items = context.read<CommandStorage>();
    items.applyFilter(query);
  }

  void _applyOffset(int dt) {
    final items = context.read<CommandStorage>();
    final selectIndex = context.read<SelectIndex>();
    final index = max(min(selectIndex.index + dt, items.length - 1), 0);
    selectIndex.update(index);

    _itemScrollController.jumpTo(index: index, alignment: 0.0);
  }

  @override
  Widget build(BuildContext context) {
    return Focus(
      onKeyEvent: (FocusNode node, KeyEvent event) {
        if (event is KeyDownEvent) {
          if (event.logicalKey == LogicalKeyboardKey.arrowDown) {
            _applyOffset(1);
            return KeyEventResult.handled;
          } else if (event.logicalKey == LogicalKeyboardKey.arrowUp) {
            _applyOffset(-1);
            return KeyEventResult.handled;
          } else if (event.logicalKey == LogicalKeyboardKey.pageDown) {
            _applyOffset(-10);
            return KeyEventResult.handled;
          } else if (event.logicalKey == LogicalKeyboardKey.pageUp) {
            _applyOffset(10);
            return KeyEventResult.handled;
          }
        }

        return KeyEventResult.ignored;
      },
      child: Column(
        children: <Widget>[
          SearchWidget(
            text: "",
            hintText: 'Search text...',
            onChanged: _applyFilter,
          ),
          const Divider(),
          Expanded(
            child: FilterableList(
              itemScrollController: _itemScrollController,
              itemPositionsListener: _itemPositionsListener,
            ),
          ),
        ],
      ),
    );
  }
}
