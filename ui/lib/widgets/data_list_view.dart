import 'package:flutter/material.dart';

class DataListViewController {
  int _selectPosition = -1;
  final Map<int, _DataListItemState> _items = <int, _DataListItemState>{};

  bool select(int position) {
    if (_selectPosition == position) {
      return false;
    }

    final prevTile = _items[_selectPosition];
    if (prevTile != null) {
      prevTile.select(false);
    }

    final nextTile = _items[position];
    if (nextTile != null) {
      _selectPosition = position;
      nextTile.select(true);
      return true;
    } else {
      _selectPosition = -1;
    }

    return false;
  }

  bool isTilesExists() {
    return _items.isNotEmpty;
  }

  MapEntry<int, BuildContext>? getNearestTileContext(int position) {
    if (_items.isNotEmpty) {
      BuildContext? ctx = _items[position]?.context;
      if (ctx != null) {
        return MapEntry(position, ctx);
      }

      final nearestKey = _items.keys.reduce((curr, next) =>
          (curr - position).abs() < (next - position).abs() ? curr : next);
      ctx = _items[nearestKey]?.context;
      if (ctx != null) {
        return MapEntry(nearestKey, ctx);
      }
    }

    return null;
  }

  void _register(int position, _DataListItemState tile) {
    _items[position] = tile;
  }

  void _unregister(int position, _DataListItemState tile) {
    if (_items[position] == tile) {
      _items.remove(position);
    }
  }
}

class DataListItem extends StatefulWidget {
  final int position;
  final Widget child;
  final DataListViewController controller;

  const DataListItem(
      {required Key key,
      required this.position,
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
    widget.controller._register(widget.position, this);
  }

  @override
  void dispose() {
    widget.controller._unregister(widget.position, this);
    super.dispose();
  }

  @override
  void didUpdateWidget(DataListItem oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.position != widget.position || oldWidget.key != widget.key) {
      widget.controller._unregister(oldWidget.position, this);
      widget.controller._register(widget.position, this);
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
