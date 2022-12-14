import 'package:flutter/material.dart';

abstract class Matcher<Key> {
  bool match(RegExp rexp);
  bool equal(Iterable<Key> keys);
}

class DataFilter<Key, T extends Matcher<Key>> with ChangeNotifier {
  String _filter = "";
  final List<T> _items = [];
  final List<int> _visibleItems = [];

  DataFilter();

  add(Iterable<T> items) {
    _items.addAll(items);
  }

  remove(Iterable<Key> keys) {
    _items.removeWhere((T item) => item.equal(keys));
  }

  sort(Comparator<T> comparator) {
    _items.sort(comparator);
  }

  apply() {
    _visibleItems.clear();
    final rfilter = RegExp.escape(_filter.trim()).replaceAll(" ", ".*");
    final RegExp rexp = RegExp(rfilter, caseSensitive: false);
    for (var i = 0; i != _items.length; i++) {
      if (_items[i].match(rexp)) {
        _visibleItems.add(i);
      }
    }

    notifyListeners();
  }

  T operator [](int id) {
    return _items[id];
  }

  List<int> get visibleItems => _visibleItems;

  setFilter(String value) {
    final processedFilter = value.toLowerCase();
    if (processedFilter != _filter) {
      _filter = processedFilter;
      apply();
    }
  }
}
