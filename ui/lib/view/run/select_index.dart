import 'package:flutter/material.dart';

class SelectIndex with ChangeNotifier {
  int _index = 0;

  int get index => _index;

  void update(int index) {
    _index = index;
    notifyListeners();
  }
}
