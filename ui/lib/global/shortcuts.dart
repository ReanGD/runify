import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

enum DataListEvent {
  onTap,
  onMenu,
  onChoice,
  onFocus,
}

enum DataListMoveDir {
  up,
  down,
  pageUp,
  pageDown,
}

class MenuActivateIntent extends Intent {
  const MenuActivateIntent();
}

class SubmitIntent extends Intent {
  const SubmitIntent();
}

class OnDataListEventIntent extends Intent {
  final DataListEvent event;

  const OnDataListEventIntent(this.event);
}

class DataListMoveIntent extends Intent {
  final DataListMoveDir direction;

  const DataListMoveIntent(this.direction);
}

class ShortcutStorage with ChangeNotifier {
  final Map<Intent, SingleActivator> _intentIndex = {};
  final Map<ShortcutActivator, Intent> _appShortcuts = {};
  final Map<ShortcutActivator, Intent> _listShortcuts = {};

  ShortcutStorage() {
    final appShortcuts = {
      const SingleActivator(LogicalKeyboardKey.keyM, alt: true):
          const MenuActivateIntent(),
      const SingleActivator(LogicalKeyboardKey.enter, alt: true):
          const SubmitIntent(),
    };
    _appShortcuts.addAll(appShortcuts);

    final listShortcuts = {
      const SingleActivator(LogicalKeyboardKey.enter):
          const OnDataListEventIntent(DataListEvent.onChoice),
      const SingleActivator(LogicalKeyboardKey.arrowUp):
          const DataListMoveIntent(DataListMoveDir.up),
      const SingleActivator(LogicalKeyboardKey.arrowDown):
          const DataListMoveIntent(DataListMoveDir.down),
      const SingleActivator(LogicalKeyboardKey.pageUp):
          const DataListMoveIntent(DataListMoveDir.pageUp),
      const SingleActivator(LogicalKeyboardKey.pageDown):
          const DataListMoveIntent(DataListMoveDir.pageDown),
    };
    _listShortcuts.addAll(listShortcuts);

    for (final entry in appShortcuts.entries) {
      _intentIndex[entry.value] = entry.key;
    }
    for (final entry in listShortcuts.entries) {
      _intentIndex[entry.value] = entry.key;
    }
  }

  void setAppShortcut(SingleActivator activator, Intent intent) {
    _intentIndex[intent] = activator;
    _appShortcuts[activator] = intent;
    notifyListeners();
  }

  void setListShortcut(SingleActivator activator, Intent intent) {
    _intentIndex[intent] = activator;
    _appShortcuts[activator] = intent;
    notifyListeners();
  }

  Map<ShortcutActivator, Intent> get appShortcuts => _appShortcuts;
  Map<ShortcutActivator, Intent> get listShortcuts => _listShortcuts;
  SingleActivator? getShortcut(Intent intent) => _intentIndex[intent];
}
