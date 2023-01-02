import 'dart:async';
import 'dart:collection';

import 'package:runify/style.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter/scheduler.dart';
import 'package:flutter/rendering.dart';

enum DataItemEvent {
  onTap,
  onMenu,
  onChoice,
  onFocus,
}

typedef OnDataItemEvent = void Function(DataItemEvent event, int id);
typedef _OnDataItemIndexEvent = void Function(DataItemEvent event, int index);

class OnEventIntent extends Intent {
  final DataItemEvent event;

  const OnEventIntent(this.event);
}

class OnEventAction extends Action<OnEventIntent> {
  final DataListController controller;

  OnEventAction(this.controller);

  @override
  Object? invoke(covariant OnEventIntent intent) {
    controller.onItemEvent(intent.event);
    return null;
  }
}

class MoveFocusIntent extends Intent {
  final int offset;

  const MoveFocusIntent(this.offset);
}

class MoveFocusAction extends Action<MoveFocusIntent> {
  final DataListController controller;

  MoveFocusAction(this.controller);

  @override
  Object? invoke(covariant MoveFocusIntent intent) {
    controller.moveFocus(intent.offset);
    return null;
  }
}

abstract class DataListController {
  _DataListScroll? _dataScroll;
  static const int _pageOffset = 10;

  void _attach(_DataListScroll dataScroll) {
    _dataScroll = dataScroll;
  }

  List<int> getVisibleItems(BuildContext context);

  Map<ShortcutActivator, Intent> get shortcuts {
    return <ShortcutActivator, Intent>{
      const SingleActivator(LogicalKeyboardKey.enter):
          const OnEventIntent(DataItemEvent.onChoice),
      const SingleActivator(LogicalKeyboardKey.keyM, alt: true):
          const OnEventIntent(DataItemEvent.onMenu),
      const SingleActivator(LogicalKeyboardKey.arrowUp):
          const MoveFocusIntent(-1),
      const SingleActivator(LogicalKeyboardKey.arrowDown):
          const MoveFocusIntent(1),
      const SingleActivator(LogicalKeyboardKey.pageUp):
          const MoveFocusIntent(-_pageOffset),
      const SingleActivator(LogicalKeyboardKey.pageDown):
          const MoveFocusIntent(_pageOffset),
    };
  }

  Map<Type, Action<Intent>> get actions {
    return <Type, Action<Intent>>{
      OnEventIntent: OnEventAction(this),
      MoveFocusIntent: MoveFocusAction(this),
    };
  }

  Future onItemEvent(DataItemEvent event) async {
    return _dataScroll?.onItemIndexEvent(event);
  }

  Future moveFocus(int moveOffset) async {
    return _dataScroll?.moveFocus(moveOffset);
  }
}

class _DataListScroll extends ScrollController {
  int _indexCount = 0;
  int _focusedIndex = 0;
  _OnDataItemIndexEvent? _onItemIndexEvent;
  final Map<int, _DataListItemState> _items = <int, _DataListItemState>{};

  _DataListScroll({double initialScrollOffset = 0.0})
      : super(initialScrollOffset: initialScrollOffset, keepScrollOffset: true);

  Future onItemIndexEvent(DataItemEvent event, {int? index}) async {
    return _onItemIndexEvent?.call(event, index ?? _focusedIndex);
  }

  Future moveFocus(int moveOffset) async {
    return _runByOrder(this, () => _moveFocus(moveOffset));
  }

  void update(int indexCount, _OnDataItemIndexEvent callbackOnItemIndexEvent) {
    _indexCount = indexCount;
    _onItemIndexEvent = callbackOnItemIndexEvent;
    _focusedIndex =
        indexCount > 0 ? _focusedIndex.clamp(0, _indexCount - 1) : -1;
    onItemIndexEvent(DataItemEvent.onFocus, index: _focusedIndex);
  }

  void register(int index, _DataListItemState item) {
    _items[index] = item;
  }

  void unregister(int index, _DataListItemState item) {
    if (_items[index] == item) {
      _items.remove(index);
    }
  }

  bool isFocused(int index) {
    return index == _focusedIndex;
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

      if (foundTarget && _focusedIndex != index) {
        final prevIndex = _focusedIndex;
        final nextIndex = index;
        _focusedIndex = index;
        _items[prevIndex]?.update();
        _items[nextIndex]?.update();
        onItemIndexEvent(DataItemEvent.onFocus);
      }

      while (attempts != 0 && hasClients && offset != scrollOffset) {
        jumpTo(scrollOffset);
        await _waitForWidgetStateBuild();
        attempts--;
      }
    }

    return foundTarget;
  }

  Future _moveFocus(int moveOffset) async {
    int index = (_focusedIndex + moveOffset).clamp(0, _indexCount - 1);
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
  final int id;
  final int index;
  final Widget child;
  final _DataListScroll dataScroll;

  const _DataListItem(
      {Key? key,
      required this.id,
      required this.index,
      required this.dataScroll,
      required this.child})
      : super(key: key);

  @override
  State<_DataListItem> createState() => _DataListItemState();
}

class _DataListItemState extends State<_DataListItem>
    with AutomaticKeepAliveClientMixin<_DataListItem> {
  bool _hovering = false;
  InkHighlight? _highlight;

  void update() {
    if (mounted) {
      setState(() {
        if (_hovering) {
          _hovering = false;
          _highlight?.dispose();
          _highlight = null;
        }
      });
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
    if (oldWidget.id != widget.id || oldWidget.index != widget.index) {
      widget.dataScroll.unregister(oldWidget.index, this);
      widget.dataScroll.register(widget.index, this);
    }
  }

  @override
  void deactivate() {
    _highlight?.dispose();
    _highlight = null;
    super.deactivate();
  }

  @override
  bool get wantKeepAlive => _highlight != null;

  void _handleMouseEnter() {
    _hovering = true;
    void handleInkRemoval() {
      if (_highlight != null) {
        _highlight = null;
        updateKeepAlive();
      }
    }

    if (_highlight != null && _highlight!.active) {
      return;
    }

    if (_highlight == null) {
      final RenderBox referenceBox = context.findRenderObject()! as RenderBox;
      final ThemeData theme = Theme.of(context);
      _highlight = InkHighlight(
        controller: Material.of(context)!,
        referenceBox: referenceBox,
        color: theme.hoverColor,
        onRemoved: handleInkRemoval,
        textDirection: Directionality.of(context),
        fadeDuration: const Duration(milliseconds: 50),
      );
      updateKeepAlive();
    } else {
      _highlight?.activate();
    }

    assert(_highlight != null && _highlight!.active);
  }

  void _handleMouseExit() {
    _hovering = false;
    _highlight?.deactivate();
  }

  @override
  Widget build(BuildContext context) {
    super.build(context);
    final ThemeData theme = Theme.of(context);

    return MouseRegion(
      onEnter: (event) => _handleMouseEnter(),
      onExit: (event) => _handleMouseExit(),
      child: GestureDetector(
        onTap: () {
          widget.dataScroll
              .onItemIndexEvent(DataItemEvent.onTap, index: widget.index);
        },
        behavior: HitTestBehavior.opaque,
        excludeFromSemantics: true,
        child: DecoratedBox(
          decoration: BoxDecoration(
            borderRadius:
                const BorderRadius.all(Radius.circular(defaultRadius)),
            color: widget.dataScroll.isFocused(widget.index)
                ? theme.focusColor
                : null,
          ),
          child: widget.child,
        ),
      ),
    );
  }
}

class DataListView extends StatelessWidget {
  final bool shrinkWrap;
  final EdgeInsetsGeometry? padding;
  final List<int> _visibleItems = [];
  final _DataListScroll _dataScroll = _DataListScroll();
  final DataListController controller;
  final OnDataItemEvent? onDataItemEvent;
  final IndexedWidgetBuilder itemBuilder;

  DataListView(
      {super.key,
      this.shrinkWrap = false,
      this.padding,
      required this.controller,
      this.onDataItemEvent,
      required this.itemBuilder});

  void _onItemIndexEvent(DataItemEvent event, int index) {
    if (index >= 0 && index < _visibleItems.length) {
      final id = _visibleItems[index];
      onDataItemEvent?.call(event, id);
    }
  }

  @override
  Widget build(BuildContext context) {
    final visibleItems = controller.getVisibleItems(context);
    _visibleItems.clear();
    _visibleItems.addAll(visibleItems);
    controller._attach(_dataScroll);
    _dataScroll.update(visibleItems.length, _onItemIndexEvent);

    return ListView.builder(
      scrollDirection: Axis.vertical,
      reverse: false,
      controller: _dataScroll,
      shrinkWrap: shrinkWrap,
      padding: padding,
      itemCount: visibleItems.length,
      itemBuilder: (context, index) {
        final id = _visibleItems[index];
        return _DataListItem(
          id: id,
          index: index,
          dataScroll: _dataScroll,
          child: itemBuilder(context, id),
        );
      },
    );
  }
}

/// used to invoke async functions in order
Future<T> _runByOrder<T>(key, FutureOr<T> Function() action) async {
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
