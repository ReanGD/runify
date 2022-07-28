import 'package:flutter/material.dart';

class DataListViewController {
  int _selectedIndex = -1;
  final Map<int, _DataListItemState> _items = <int, _DataListItemState>{};

  bool select(int index) {
    if (_selectedIndex == index) {
      return false;
    }

    final prevItem = _items[_selectedIndex];
    if (prevItem != null) {
      prevItem.select(false);
    }

    final nextItem = _items[index];
    if (nextItem != null) {
      _selectedIndex = index;
      nextItem.select(true);
      return true;
    } else {
      _selectedIndex = -1;
    }

    return false;
  }

  bool isItemsExists() {
    return _items.isNotEmpty;
  }

  MapEntry<int, BuildContext>? getNearestItem(int index) {
    if (_items.isNotEmpty) {
      BuildContext? ctx = _items[index]?.context;
      if (ctx != null) {
        return MapEntry(index, ctx);
      }

      final nearestKey = _items.keys.reduce((curr, next) =>
          (curr - index).abs() < (next - index).abs() ? curr : next);
      ctx = _items[nearestKey]?.context;
      if (ctx != null) {
        return MapEntry(nearestKey, ctx);
      }
    }

    return null;
  }

  void _register(int index, _DataListItemState item) {
    _items[index] = item;
  }

  void _unregister(int index, _DataListItemState item) {
    if (_items[index] == item) {
      _items.remove(index);
    }
  }
}

class DataListItem extends StatefulWidget {
  final int index;
  final Widget child;
  final DataListViewController controller;

  const DataListItem(
      {required Key key,
      required this.index,
      required this.controller,
      required this.child})
      : super(key: key);

  @override
  State<DataListItem> createState() => _DataListItemState();
}

class _DataListItemState extends State<DataListItem> {
  bool _selected = false;

  void select(bool isSelected) {
    if (isSelected == _selected) {
      return;
    }
    _selected = isSelected;

    if (mounted) {
      setState(() {});
    }
  }

  @override
  void initState() {
    super.initState();
    widget.controller._register(widget.index, this);
  }

  @override
  void dispose() {
    widget.controller._unregister(widget.index, this);
    super.dispose();
  }

  @override
  void didUpdateWidget(DataListItem oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.index != widget.index || oldWidget.key != widget.key) {
      widget.controller._unregister(oldWidget.index, this);
      widget.controller._register(widget.index, this);
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_selected) {
      final ThemeData theme = Theme.of(context);
      final ListTileThemeData tileTheme = ListTileTheme.of(context);
      final selectedColor = tileTheme.selectedColor ??
          theme.listTileTheme.selectedColor ??
          theme.colorScheme.primary;

      return ColoredBox(
        color: selectedColor,
        child: widget.child,
      );
    }

    return widget.child;
  }
}
