import 'package:flutter/material.dart';

abstract class Matcher {
  bool match(String filter);
}

class DataFilter<T extends Matcher> with ChangeNotifier {
  String _filter = "";
  final List<T> _items = [];
  final List<int> _visibleItems = [];

  DataFilter(Future<Iterable<T>> items) {
    items.then((Iterable<T> data) {
      _items.addAll(data);
      _update();
    });
  }

  T operator [](int id) {
    return _items[id];
  }

  List<int> get visibleItems => _visibleItems;

  void _update() {
    _visibleItems.clear();
    for (var i = 0; i != _items.length; i++) {
      if (_items[i].match(_filter)) {
        _visibleItems.add(i);
      }
    }

    notifyListeners();
  }

  void applyFilter(String value) {
    final filter = value.toLowerCase();
    if (filter != _filter) {
      _filter = filter;
      _update();
    }
  }
}
