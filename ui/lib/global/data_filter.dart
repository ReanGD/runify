import 'package:flutter/material.dart';

abstract class Matcher<Key> {
  bool match(RegExp rexp);
  bool equal(Iterable<Key> keys);
}

class DataFilter<Key, T extends Matcher<Key>> with ChangeNotifier {
  String _filter = "";
  final Comparator<T> _comparator;
  final List<T> _items = [];
  final List<int> _visibleItems = [];

  DataFilter(this._comparator) {
    _items.clear();
    _update();
  }

  DataFilter.future(Future<Iterable<T>> items, this._comparator) {
    items.then((Iterable<T> data) {
      _items.addAll(data);
      _update();
    });
  }

  DataFilter.value(Iterable<T> items, this._comparator) {
    _items.addAll(items);
    _update();
  }

  T operator [](int id) {
    return _items[id];
  }

  List<int> get visibleItems => _visibleItems;

  void _update() {
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

  upsert(Iterable<T> items) {
    _items.addAll(items);
    _items.sort(_comparator);
    _update();
  }

  remove(Iterable<Key> keys) {
    _items.removeWhere((T item) => item.equal(keys));
    _update();
  }

  applyFilter(String value) {
    final filter = value.toLowerCase();
    if (filter != _filter) {
      _filter = filter;
      _update();
    }
  }
}
