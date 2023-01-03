import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

class MenuActivateIntent extends Intent {
  const MenuActivateIntent();
}

class SubmitIntent extends Intent {
  const SubmitIntent();
}

Map<ShortcutActivator, Intent> appShortcuts() {
  return <ShortcutActivator, Intent>{
    const SingleActivator(LogicalKeyboardKey.keyM, alt: true):
        const MenuActivateIntent(),
    const SingleActivator(LogicalKeyboardKey.enter, alt: true):
        const SubmitIntent(),
  };
}
