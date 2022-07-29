import 'dart:async';
import 'dart:collection';

import 'package:flutter/material.dart';
import 'package:flutter/scheduler.dart';
import 'package:flutter/rendering.dart';

class DataListController {
  _DataListScroll? _dataScroll;

  void _attach(_DataListScroll dataScroll) {
    _dataScroll = dataScroll;
  }

  Future scrollTo(int index) async {
    return _dataScroll?.scrollTo(index);
  }
}

class _DataListScroll extends ScrollController {
  int _selectedIndex = 0;
  final Map<int, _DataListItemState> _items = <int, _DataListItemState>{};

  _DataListScroll({double initialScrollOffset = 0.0})
      : super(initialScrollOffset: initialScrollOffset, keepScrollOffset: true);

  Future scrollTo(int index) async {
    return addToOrder(this, () => _scrollTo(index));
  }

  void register(int index, _DataListItemState item) {
    _items[index] = item;
  }

  void unregister(int index, _DataListItemState item) {
    if (_items[index] == item) {
      _items.remove(index);
    }
  }

  bool isSelected(int index) {
    return index == _selectedIndex;
  }

  MapEntry<int, BuildContext>? _getNearestItem(int index) {
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

  /// return offset, which is a absolute offset to bring the target object into the viewport with "alignment".
  RevealedOffset _offsetToRevealInViewport(
      BuildContext object, double alignment) {
    final renderBox = object.findRenderObject()!;
    assert(Scrollable.of(object) != null);
    final RenderAbstractViewport viewport =
        RenderAbstractViewport.of(renderBox)!;
    return viewport.getOffsetToReveal(renderBox, alignment);
  }

  Future<bool> _jumpToNearest(int index, bool useSuggested) async {
    bool foundTarget = false;
    final nearestItem = _getNearestItem(index);
    if (nearestItem != null) {
      int attempts = 1;
      double alignment = index > nearestItem.key ? 1.0 : 0.0;
      if (index == nearestItem.key) {
        // not sure why it doesn't scroll to the given offset, try more within 5 times
        attempts = 5;
        alignment = 0.5;
        foundTarget = true;
      }
      RevealedOffset revealedOffset =
          _offsetToRevealInViewport(nearestItem.value, alignment);

      double scrollOffset = revealedOffset.offset;
      if (useSuggested && !foundTarget) {
        scrollOffset +=
            ((index - nearestItem.key) * revealedOffset.rect.height);
      }

      // The content preferred position might be impossible to reach
      // for items close to the edges of the scroll content, e.g.
      // we cannot put the first item at the end of the viewport or
      // the last item at the beginning. Trying to do so might lead
      // to a bounce at either the top or bottom, unless the scroll
      // physics are set to clamp. To prevent this, we limit the
      // offset to not overshoot the extent in either direction.
      scrollOffset = scrollOffset.clamp(
          position.minScrollExtent, position.maxScrollExtent);

      if (foundTarget && _selectedIndex != index) {
        final prevIndex = _selectedIndex;
        final nextIndex = index;
        _selectedIndex = index;
        _items[prevIndex]?.update();
        _items[nextIndex]?.update();
      }

      while (attempts != 0 && hasClients && offset != scrollOffset) {
        jumpTo(scrollOffset);
        await _waitForWidgetStateBuild();
        attempts--;
      }
    }

    return foundTarget;
  }

  Future _scrollTo(int index) async {
    // In listView init or reload case, widget state of list item may not be ready for query.
    // this prevent from over scrolling becoming empty screen or unnecessary scroll bounce.
    const maxBound = 30; // 0.5 second if 60fps
    for (var count = 0; count != maxBound; count++) {
      if (_items.isEmpty) {
        await _waitForWidgetStateBuild();
      } else {
        break;
      }
    }

    bool usedSuggestedRowHeight = true;
    while (hasClients) {
      final oldOffset = offset;
      final foundTarget = await _jumpToNearest(index, usedSuggestedRowHeight);
      usedSuggestedRowHeight = false; // just use once
      if (foundTarget || offset == oldOffset) {
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

class _DataListItem extends StatefulWidget {
  final int index;
  final Widget child;
  final _DataListScroll dataScroll;

  const _DataListItem(
      {required Key key,
      required this.index,
      required this.dataScroll,
      required this.child})
      : super(key: key);

  @override
  State<_DataListItem> createState() => _DataListItemState();
}

class _DataListItemState extends State<_DataListItem> {
  void update() {
    if (mounted) {
      setState(() {});
    }
  }

  @override
  void initState() {
    super.initState();
    widget.dataScroll.register(widget.index, this);
  }

  @override
  void dispose() {
    widget.dataScroll.unregister(widget.index, this);
    super.dispose();
  }

  @override
  void didUpdateWidget(_DataListItem oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.index != widget.index || oldWidget.key != widget.key) {
      widget.dataScroll.unregister(oldWidget.index, this);
      widget.dataScroll.register(widget.index, this);
    }
  }

  @override
  Widget build(BuildContext context) {
    if (widget.dataScroll.isSelected(widget.index)) {
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

Widget dataListView(
    {required DataListController controller,
    required int itemCount,
    required IndexedWidgetBuilder itemBuilder}) {
  final dataScroll = _DataListScroll();
  controller._attach(dataScroll);
  return ListView.builder(
    scrollDirection: Axis.vertical,
    reverse: false,
    controller: dataScroll,
    itemCount: itemCount,
    itemBuilder: (context, index) {
      return _DataListItem(
        key: ValueKey(index),
        dataScroll: dataScroll,
        index: index,
        child: itemBuilder(context, index),
      );
    },
  );
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
