import 'package:flutter/material.dart';

class IndexedListViewController {
  int _selectPosition = -1;
  final Map<int, _IndexedListTileState> _tiles = <int, _IndexedListTileState>{};

  bool select(int position) {
    if (_selectPosition == position) {
      return false;
    }

    final prevTile = _tiles[_selectPosition];
    if (prevTile != null) {
      prevTile.select(false);
    }

    final nextTile = _tiles[position];
    if (nextTile != null) {
      _selectPosition = position;
      nextTile.select(true);
      return true;
    } else {
      _selectPosition = -1;
    }

    return false;
  }

  BuildContext? getTileContext(int position) {
    return _tiles[position]?.context;
  }

  bool isTilesExists() {
    return _tiles.isNotEmpty;
  }

  int? getNearestPosition(int position) {
    if (_tiles.isEmpty) {
      return null;
    }

    final sorted = _tiles.keys.toList()
      ..sort((int first, int second) => first.compareTo(second));

    final minPosition = sorted.first;
    final maxPosition = sorted.last;

    return (position - minPosition).abs() < (position - maxPosition).abs()
        ? minPosition
        : maxPosition;
  }

  void _register(int position, _IndexedListTileState tile) {
    _tiles[position] = tile;
  }

  void _unregister(int position, _IndexedListTileState tile) {
    if (_tiles[position] == tile) {
      _tiles.remove(position);
    }
  }
}

class IndexedListTile extends StatefulWidget {
  final int position;
  final Widget child;
  final IndexedListViewController controller;

  const IndexedListTile(
      {required Key key,
      required this.position,
      required this.controller,
      required this.child})
      : super(key: key);

  @override
  State<IndexedListTile> createState() => _IndexedListTileState();
}

class _IndexedListTileState extends State<IndexedListTile> {
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
  void didUpdateWidget(IndexedListTile oldWidget) {
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
