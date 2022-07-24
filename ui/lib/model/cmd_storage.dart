import 'package:flutter/material.dart';
import 'package:runify/model/cmd.dart';
import 'package:runify/model/cmd_loader.dart';

// TODO: add sync
class CommandStorage with ChangeNotifier {
  List<Command> _items = [];
  final List<int> _filteredItems = [];
  String _filter = "";

  CommandStorage() {
    CommandLoader.instance.commands.then((data) {
      _items = data;
      _update();
    });
  }

  int get length => _filteredItems.length;
  Command operator [](int index) {
    return _items[_filteredItems[index]];
  }

  void _update() {
    _filteredItems.clear();
    for (var i = 0; i != _items.length; i++) {
      if (_items[i].name().toLowerCase().contains(_filter)) {
        _filteredItems.add(i);
      }
    }
    notifyListeners();
  }

  void applyFilter(String query) {
    final filter = query.toLowerCase();
    if (filter != _filter) {
      _filter = filter;
      _update();
    }
  }
}
