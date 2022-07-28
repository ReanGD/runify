import 'dart:async';
import 'dart:collection';

import 'package:flutter/material.dart';
import 'package:flutter/scheduler.dart';
import 'package:flutter/rendering.dart';

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

class DataListScroll extends ScrollController {
  final double? suggestedRowHeight;
  final DataListViewController controller;

  DataListScroll(this.controller,
      {double initialScrollOffset = 0.0, this.suggestedRowHeight})
      : super(initialScrollOffset: initialScrollOffset, keepScrollOffset: true);

  Future scrollTo(int postition) async {
    return addToOrder(this, () => _scrollTo(postition));
  }

  /// return offset, which is a absolute offset to bring the target object into the viewport with "alignment".
  double _offsetToRevealInViewport(BuildContext object, double alignment) {
    final renderBox = object.findRenderObject()!;
    assert(Scrollable.of(object) != null);
    final RenderAbstractViewport viewport =
        RenderAbstractViewport.of(renderBox)!;
    final revealedOffset = viewport.getOffsetToReveal(renderBox, alignment);
    final absoluteOffset = revealedOffset.offset;

    return absoluteOffset;
  }

  Future<bool> _jumpToNearest(int pos, bool useSuggested) async {
    bool stop = false;
    final nearestItem = controller.getNearestItem(pos);
    if (nearestItem != null) {
      int attempts = 1;
      double alignment = pos > nearestItem.key ? 1.0 : 0.0;
      if (pos == nearestItem.key) {
        // not sure why it doesn't scroll to the given offset, try more within 5 times
        attempts = 5;
        alignment = 0.5;
        stop = true;
      }
      double targetOffset =
          _offsetToRevealInViewport(nearestItem.value, alignment);

      if (useSuggested && suggestedRowHeight != null && !stop) {
        targetOffset += ((pos - nearestItem.key) * suggestedRowHeight!);
      }

      // The content preferred position might be impossible to reach
      // for items close to the edges of the scroll content, e.g.
      // we cannot put the first item at the end of the viewport or
      // the last item at the beginning. Trying to do so might lead
      // to a bounce at either the top or bottom, unless the scroll
      // physics are set to clamp. To prevent this, we limit the
      // offset to not overshoot the extent in either direction.
      targetOffset = targetOffset.clamp(
          position.minScrollExtent, position.maxScrollExtent);

      while (attempts != 0 && hasClients && offset != targetOffset) {
        jumpTo(targetOffset);
        await _waitForWidgetStateBuild();
        attempts--;
      }
    }

    return stop;
  }

  Future _scrollTo(int position) async {
    // In listView init or reload case, widget state of list item may not be ready for query.
    // this prevent from over scrolling becoming empty screen or unnecessary scroll bounce.
    const maxBound = 30; // 0.5 second if 60fps
    for (var count = 0; count != maxBound; count++) {
      if (!controller.isItemsExists()) {
        await _waitForWidgetStateBuild();
      } else {
        break;
      }
    }

    bool usedSuggestedRowHeight = true;
    while (hasClients) {
      final oldOffset = offset;
      final stop = await _jumpToNearest(position, usedSuggestedRowHeight);
      usedSuggestedRowHeight = false; // just use once
      if (stop || offset == oldOffset) {
        break;
      }
    }

    //after auto scrolling, we should sync final scroll position without flag on
    if (hasClients) {
      notifyListeners();
    }
  }

  /// wait until the [SchedulerPhase] in [SchedulerPhase.persistentCallbacks].
  /// it means if we do animation scrolling to a position, the Future call back will in [SchedulerPhase.midFrameMicrotasks].
  /// if we want to search viewport element depending on Widget State, we must delay it to [SchedulerPhase.persistentCallbacks].
  /// which is the phase widget build/layout/draw
  Future _waitForWidgetStateBuild() => SchedulerBinding.instance.endOfFrame;
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

/// used to invoke async functions in order
Future<T> addToOrder<T>(key, FutureOr<T> Function() action) async {
  for (;;) {
    final c = _locks[key];
    if (c == null) break;
    try {
      await c.future;
    } catch (_) {} //ignore error (so it will continue)
  }

  final c = _locks[key] = Completer<T>();
  void then(T result) {
    final c2 = _locks.remove(key);
    c.complete(result);

    assert(identical(c, c2));
  }

  void catchError(ex, StackTrace st) {
    final c2 = _locks.remove(key);
    c.completeError(ex, st);

    assert(identical(c, c2));
  }

  try {
    final result = action();
    if (result is Future<T>) {
      result.then(then).catchError(catchError);
    } else {
      then(result);
    }
  } catch (ex, st) {
    catchError(ex, st);
  }

  return c.future;
}

final _locks = HashMap<dynamic, Completer>();
