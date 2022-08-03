import 'package:flutter/material.dart';
import 'package:runify/model/cmd.dart';

class CommandFilter with ChangeNotifier {
  String _filter = "";
  final List<Command> _items = [];
  final List<int> _visibleItems = [];

  CommandFilter(Future<List<Command>> items) {
    items.then((data) {
      _items.addAll(data);
      _update();
    });
  }

  Command operator [](int id) {
    return _items[id];
  }

  List<int> get visibleItems => _visibleItems;

  void _update() {
    _visibleItems.clear();
    for (var i = 0; i != _items.length; i++) {
      if (_items[i].name().toLowerCase().contains(_filter)) {
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
